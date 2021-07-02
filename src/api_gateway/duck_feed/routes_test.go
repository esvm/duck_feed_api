package duck_feed

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esvm/duck_feed_api/src/duck_feeds_service"
	"github.com/esvm/duck_feed_api/src/duck_feeds_service/store/postgres"
	"github.com/go-kit/kit/log"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var logger = log.NewNopLogger()

func buildAPI() *DuckFeedAPI {
	svc := duck_feeds_service.NewDuckFeedService(logger)
	return &DuckFeedAPI{svc}
}

func truncateTable() {
	db := postgres.NewDatabase(logger)
	conn, _ := db.GetConnection()
	conn.Exec("DELETE FROM duck_feeds")
	conn.Exec("DELETE FROM duck_feed_schedules")
}

func TestInsertDuckFeedReportHandlerSuccess(t *testing.T) {
	defer truncateTable()
	api := buildAPI()

	report := duckFeed()
	json, _ := json.Marshal(report)
	body := bytes.NewReader(json)

	req, err := http.NewRequest(echo.POST, "/", body)
	if !assert.NoError(t, err) {
		return
	}

	req.Header.Set(echo.HeaderContentType, "application/hal+json")

	e := echo.New()
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err = api.InsertDuckFeedReportHandler(ctx)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestInsertDuckFeedReportHandlerFailParseBody(t *testing.T) {
	defer truncateTable()
	api := buildAPI()

	json, _ := json.Marshal("nil")
	body := bytes.NewReader(json)

	req, err := http.NewRequest(echo.POST, "/", body)
	if !assert.NoError(t, err) {
		return
	}

	req.Header.Set(echo.HeaderContentType, "application/hal+json")

	e := echo.New()
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err = api.InsertDuckFeedReportHandler(ctx)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	}
}

func TestInsertDuckFeedReportHandlerFailValidation(t *testing.T) {
	defer truncateTable()
	api := buildAPI()

	report := duckFeed()
	report.DucksCount = 0
	json, _ := json.Marshal(report)
	body := bytes.NewReader(json)

	req, err := http.NewRequest(echo.POST, "/", body)
	if !assert.NoError(t, err) {
		return
	}

	req.Header.Set(echo.HeaderContentType, "application/hal+json")

	e := echo.New()
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err = api.InsertDuckFeedReportHandler(ctx)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	}
}

func TestGetDuckFeedsReportHandler(t *testing.T) {
	api := buildAPI()

	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := api.GetDuckFeedsReportHandler(ctx)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestInsertDuckFeedscheduleHandlerSuccess(t *testing.T) {
	defer truncateTable()
	api := buildAPI()

	report := duckFeed()
	json, _ := json.Marshal(report)
	body := bytes.NewReader(json)

	req, err := http.NewRequest(echo.POST, "/schedule/", body)
	if !assert.NoError(t, err) {
		return
	}

	req.Header.Set(echo.HeaderContentType, "application/hal+json")

	e := echo.New()
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err = api.InsertDuckFeedScheduleHandler(ctx)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestInsertDuckFeedScheduleHandlerFailParseBody(t *testing.T) {
	defer truncateTable()
	api := buildAPI()

	json, _ := json.Marshal("nil")
	body := bytes.NewReader(json)

	req, err := http.NewRequest(echo.POST, "/", body)
	if !assert.NoError(t, err) {
		return
	}

	req.Header.Set(echo.HeaderContentType, "application/hal+json")

	e := echo.New()
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err = api.InsertDuckFeedScheduleHandler(ctx)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	}
}

func TestInsertDuckFeedScheduleHandlerFailValidation(t *testing.T) {
	defer truncateTable()
	api := buildAPI()

	report := duckFeed()
	report.DucksCount = 0
	json, _ := json.Marshal(report)
	body := bytes.NewReader(json)

	req, err := http.NewRequest(echo.POST, "/", body)
	if !assert.NoError(t, err) {
		return
	}

	req.Header.Set(echo.HeaderContentType, "application/hal+json")

	e := echo.New()
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err = api.InsertDuckFeedScheduleHandler(ctx)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusBadRequest, err.(*echo.HTTPError).Code)
	}
}
