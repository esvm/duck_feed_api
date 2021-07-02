package middleware

import (
	"time"

	"github.com/esvm/duck_feed_api/src/logger_builder"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rollbar/rollbar-go"
)

type LoggerConfig struct {
	Skipper middleware.Skipper
	Logger  log.Logger
}

var DefaultLoggerConfig = LoggerConfig{
	Skipper: middleware.DefaultSkipper,
	Logger:  logger_builder.NewLogger("api-gateway-rest"),
}

func Logger() echo.MiddlewareFunc {
	return LoggerWithConfig(DefaultLoggerConfig)
}

func LoggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {

	if config.Skipper == nil {
		config.Skipper = DefaultLoggerConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			start := time.Now()
			req := c.Request()
			res := c.Response()
			err := next(c)
			stop := time.Now()

			if err != nil {
				c.Error(err)
				rollbar.Error(err)
			}

			level.Info(config.Logger).Log(
				"request_id", c.Get("request_id"),
				"status", res.Status,
				"uri", req.RequestURI,
				"method", req.Method,
				"path", req.URL.Path,
				"latency", stop.Sub(start).String(),
				"err", err,
			)

			return err
		}
	}
}
