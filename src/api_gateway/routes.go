package api_gateway

import (
	"github.com/esvm/duck_feed_api/src/api_gateway/duck_feed"
	"github.com/esvm/duck_feed_api/src/duck_feeds_service"
	"github.com/labstack/echo"
)

type Clients struct {
	DuckFeedService duck_feeds_service.DuckFeedService
}

func MakeRoutes(app *echo.Echo, clients Clients) {
	duckFeedReportRoutes := app.Group(duck_feed.EntryPoint)
	duck_feed.MakeDuckFeedRoutes(duckFeedReportRoutes, clients.DuckFeedService)
}
