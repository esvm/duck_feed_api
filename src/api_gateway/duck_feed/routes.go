package duck_feed

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esvm/duck_feed_api/src/api_gateway/context"
	"github.com/esvm/duck_feed_api/src/duck_feed"
	"github.com/esvm/duck_feed_api/src/duck_feeds_service"
	"github.com/labstack/echo"
)

const (
	contentType = "application/json"

	EntryPoint = "/reports"

	GetDuckFeedsReportRoute     = "/"
	InsertDuckFeedReportRoute   = "/"
	InsertDuckFeedScheduleRoute = "/schedule/"
)

type DuckFeedAPI struct {
	duckFeedService duck_feeds_service.DuckFeedService
}

func MakeDuckFeedRoutes(
	g *echo.Group,
	duckFeedService duck_feeds_service.DuckFeedService,
) {
	api := &DuckFeedAPI{duckFeedService}

	g.GET(
		GetDuckFeedsReportRoute,
		api.GetDuckFeedsReportHandler,
		instrumentingMiddleware("GetDuckFeedsReport"),
	)
	g.POST(
		InsertDuckFeedReportRoute,
		api.InsertDuckFeedReportHandler,
		instrumentingMiddleware("InsertDuckFeedReport"),
	)
	g.POST(
		InsertDuckFeedScheduleRoute,
		api.InsertDuckFeedScheduleHandler,
		instrumentingMiddleware("InsertDuckFeedSchedule"),
	)
}

func (api *DuckFeedAPI) GetDuckFeedsReportHandler(ctx echo.Context) error {
	c := context.GetContext(ctx)
	reports, err := api.duckFeedService.GetDuckFeedsReport(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Duck feed service failed: %s", err.Error()),
		}
	}

	body, err := json.Marshal(reports)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to get duck feed reports"}
	}

	return ctx.Blob(http.StatusOK, contentType, body)
}

func (api *DuckFeedAPI) InsertDuckFeedReportHandler(ctx echo.Context) error {
	report, err := UnmarshalDuckFeedReport(ctx)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Failed to parse body %s", err.Error())}
	}

	err = ValidateDuckFeed(report)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Duck feed is not valid %s", err.Error()),
		}
	}

	c := context.GetContext(ctx)

	created, err := api.duckFeedService.InsertDuckFeedReport(c, ParseRequestDuckFeedToDomainDuckFeed(report))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Duck feed service failed %s", err.Error()),
		}
	}

	body, _ := json.Marshal(created)

	return ctx.Blob(http.StatusCreated, contentType, body)
}

func (api *DuckFeedAPI) InsertDuckFeedScheduleHandler(ctx echo.Context) error {
	report, err := UnmarshalDuckFeedReport(ctx)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Failed to parse body %s", err.Error())}
	}

	err = ValidateDuckFeed(report)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Duck feed is not valid %s", err.Error()),
		}
	}

	c := context.GetContext(ctx)

	created, err := api.duckFeedService.InsertDuckFeedReport(c, ParseRequestDuckFeedToDomainDuckFeed(report))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Duck feed service failed to create report %s", err.Error()),
		}
	}

	schedule := &duck_feed.DuckFeedSchedule{
		DuckFeedId: created.ID,
	}

	createdSchedule, err := api.duckFeedService.InsertDuckFeedSchedule(c, schedule)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Duck feed service failed to create schedule %s", err.Error()),
		}
	}

	body, _ := json.Marshal(createdSchedule)
	return ctx.Blob(http.StatusCreated, contentType, body)
}
