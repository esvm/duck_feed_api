package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/esvm/duck_feed_api/src/api_gateway"
	custom_middleware "github.com/esvm/duck_feed_api/src/api_gateway/middleware"
	"github.com/esvm/duck_feed_api/src/duck_feeds_service"
	"github.com/esvm/duck_feed_api/src/logger_builder"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rollbar/rollbar-go"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	return port
}

func createClients(logger log.Logger) api_gateway.Clients {
	clients := api_gateway.Clients{}
	clients.DuckFeedService = duck_feeds_service.NewDuckFeedService(logger)

	return clients
}

func startServer(app *echo.Echo, port string) {
	app.Logger.Info("starting the server")

	if err := app.Start(":" + port); err != nil {
		app.Logger.Errorf("shutting down the server: %s", err)
	}
}

func setupRollbar() {
	rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
	rollbar.SetEnvironment(os.Getenv("ENVIRONMENT"))
	rollbar.SetServerRoot("github.com/esvm/duck_feed_api")
}

func createCORSConfig(logger log.Logger) echo.MiddlewareFunc {
	maxAge, err := strconv.Atoi(os.Getenv("API_GATEWAY_CORS_MAX_AGE"))
	if err != nil {
		level.Error(logger).Log("err", err, "message", "API_GATEWAY_CORS_MAX_AGE should be an integer")
	}

	origins := strings.Split(os.Getenv("API_GATEWAY_ALLOWED_ORIGINS"), ",")

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: origins,
		AllowMethods: []string{echo.OPTIONS, echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
		MaxAge:       maxAge,
	})
}

func setupAppAndMiddlewares(logger log.Logger) *echo.Echo {
	app := echo.New()

	app.Use(custom_middleware.LoggerWithConfig(
		custom_middleware.LoggerConfig{
			Logger: logger,
		},
	))
	app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	app.Use(createCORSConfig(logger))
	app.Use(custom_middleware.ContextInjector())
	app.Use(custom_middleware.RequestID())

	return app
}

func getNewTime(oldTime time.Time) time.Time {
	now := time.Now()
	year, month, day := now.Date()
	hour := oldTime.Hour()
	minute := oldTime.Minute()

	return time.Date(year, month, day, hour, minute, 0, 0, now.Location())
}

func executeSchedules(clients api_gateway.Clients, logger log.Logger) {
	for {
		schedules, err := clients.DuckFeedService.GetDuckFeedSchedules(context.Background())
		if err != nil {
			level.Error(logger).Log("err", err, "message", "Fail to load schedules")
		}

		for _, schedule := range schedules {
			duckFeed, err := clients.DuckFeedService.GetDuckFeedReportById(context.Background(), schedule.DuckFeedId)
			if err != nil {
				level.Error(logger).Log("err", err, "message", fmt.Sprintf("Fail to load schedule with DuckFeedId %d", schedule.DuckFeedId))
			}

			if duckFeed != nil {
				duckFeed.ID = 0
				duckFeed.Time = getNewTime(duckFeed.Time)
				duckFeed.CreatedAt = time.Now()

				_, err = clients.DuckFeedService.InsertDuckFeedReport(context.Background(), duckFeed)
				if err != nil {
					level.Error(logger).Log("err", err, "message", fmt.Sprintf("Fail to insert new report from schedule with DuckFeedId %d", schedule.DuckFeedId))
				}
			}
		}

		time.Sleep(24 * time.Hour)
	}
}

func wakeupServer() {
	for {
		time.Sleep(60 * time.Second)
		http.Get("https://duck-feed-api.herokuapp.com/reports/")
	}
}

func main() {
	logger := logger_builder.NewLogger("api-gateway")
	setupRollbar()

	app := setupAppAndMiddlewares(logger)
	clients := createClients(logger)
	port := getPort()

	api_gateway.MakeRoutes(app, clients)

	go startServer(app, port)
	go func() {
		rollbar.Wait()
	}()
	go executeSchedules(clients, logger)
	go wakeupServer()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds, any request that lasts more than that
	// will be dropped, looking as 504 (GATEWAY TIMEOUT) to the client
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}
}
