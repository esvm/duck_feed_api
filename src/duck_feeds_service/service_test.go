package duck_feeds_service

import (
	"context"
	"testing"

	"github.com/esvm/duck_feed_api/src/duck_feed"
	"github.com/esvm/duck_feed_api/src/duck_feed/fixtures"
	"github.com/esvm/duck_feed_api/src/duck_feeds_service/store/postgres"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

func assertEqual(t *testing.T, expected, received *duck_feed.DuckFeed) {
	assert.Equal(t, expected.Country, received.Country)
	assert.Equal(t, expected.ParkName, received.ParkName)
	assert.Equal(t, expected.Food, received.Food)
	assert.Equal(t, expected.FoodQuantity, received.FoodQuantity)
	assert.Equal(t, expected.FoodType, received.FoodType)
	assert.Equal(t, expected.DucksCount, received.DucksCount)
}

func truncateTable() {
	db := postgres.NewDatabase(logger)
	conn, _ := db.GetConnection()
	conn.Exec("DELETE FROM duck_feeds")
	conn.Exec("DELETE FROM duck_feed_schedules")
}

var logger = log.NewNopLogger()

func TestInsertDuckFeedReport(t *testing.T) {
	defer truncateTable()
	svc := NewDuckFeedService(logger)
	duckFeed := fixtures.DuckFeed()

	created, err := svc.InsertDuckFeedReport(context.Background(), duckFeed)
	if assert.NoError(t, err) {
		assert.NotNil(t, created)
		assert.NotZero(t, created.ID)
		assert.NotZero(t, created.CreatedAt)
		assertEqual(t, duckFeed, created)
	}
}

func TestGetDuckFeedsReport(t *testing.T) {
	defer truncateTable()
	svc := NewDuckFeedService(logger)
	duckFeed := fixtures.DuckFeed()

	created, err := svc.InsertDuckFeedReport(context.Background(), duckFeed)
	if !assert.NoError(t, err) {
		return
	}

	reports, err := svc.GetDuckFeedsReport(context.Background())
	if assert.NoError(t, err) {
		assert.Len(t, reports, 1)
		assertEqual(t, created, reports[0])
	}
}

func TestGetDuckFeedReportById(t *testing.T) {
	defer truncateTable()
	svc := NewDuckFeedService(logger)
	duckFeed := fixtures.DuckFeed()

	created, err := svc.InsertDuckFeedReport(context.Background(), duckFeed)
	if !assert.NoError(t, err) {
		return
	}

	report, err := svc.GetDuckFeedReportById(context.Background(), created.ID)
	if assert.NoError(t, err) {
		assertEqual(t, created, report)
	}
}

func TestInsertDuckFeedSchedule(t *testing.T) {
	defer truncateTable()
	svc := NewDuckFeedService(logger)
	duckFeed := fixtures.DuckFeed()

	created, err := svc.InsertDuckFeedReport(context.Background(), duckFeed)
	if !assert.NoError(t, err) {
		return
	}

	duckFeedSchedule := &duck_feed.DuckFeedSchedule{
		DuckFeedId: created.ID,
	}
	createdSchedule, err := svc.InsertDuckFeedSchedule(context.Background(), duckFeedSchedule)
	if assert.NoError(t, err) {
		assert.NotNil(t, created)
		assert.NotZero(t, created.ID)
		assert.Equal(t, createdSchedule.DuckFeedId, created.ID)
	}
}

func TestGetDuckFeedsSchedule(t *testing.T) {
	defer truncateTable()
	svc := NewDuckFeedService(logger)
	duckFeed := fixtures.DuckFeed()

	created, err := svc.InsertDuckFeedReport(context.Background(), duckFeed)
	if !assert.NoError(t, err) {
		return
	}

	duckFeedSchedule := &duck_feed.DuckFeedSchedule{
		DuckFeedId: created.ID,
	}
	createdSchedule, err := svc.InsertDuckFeedSchedule(context.Background(), duckFeedSchedule)
	if !assert.NoError(t, err) {
		return
	}

	schedules, err := svc.GetDuckFeedSchedules(context.Background())
	if assert.NoError(t, err) {
		assert.Len(t, schedules, 1)
		assert.Equal(t, createdSchedule.ID, schedules[0].ID)
		assert.Equal(t, createdSchedule.DuckFeedId, schedules[0].DuckFeedId)
	}
}
