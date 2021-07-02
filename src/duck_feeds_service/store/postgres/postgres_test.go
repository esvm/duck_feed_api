package postgres

import (
	"testing"

	"github.com/esvm/duck_feed_api/src/duck_feed"
	"github.com/esvm/duck_feed_api/src/duck_feed/fixtures"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

var logger = log.NewNopLogger()

func assertEqual(t *testing.T, expected, received *duck_feed.DuckFeed) {
	assert.Equal(t, expected.Country, received.Country)
	assert.Equal(t, expected.ParkName, received.ParkName)
	assert.Equal(t, expected.Food, received.Food)
	assert.Equal(t, expected.FoodQuantity, received.FoodQuantity)
	assert.Equal(t, expected.FoodType, received.FoodType)
	assert.Equal(t, expected.DucksCount, received.DucksCount)
}

func truncateTables(db DuckFeedDatabase) {
	conn, _ := db.GetConnection()
	conn.Exec("DELETE FROM duck_feeds")
	conn.Exec("DELETE FROM duck_feed_schedules")
}

func TestInsertDuckFeedSucess(t *testing.T) {
	db := NewDatabase(logger)
	defer truncateTables(db)

	duckFeed := fixtures.DuckFeed()

	created, err := db.InsertDuckFeedReport(duckFeed)
	if assert.NoError(t, err) && assert.NotNil(t, created) {
		assert.NotZero(t, created.ID)
		assertEqual(t, duckFeed, created)
		assert.Equal(t, duckFeed.Time, created.Time)
		assert.NotZero(t, created.CreatedAt)
	}
}

func TestInsertDuckFeedFail(t *testing.T) {
	db := NewDatabase(logger)
	defer truncateTables(db)

	duckFeed := fixtures.DuckFeed()
	duckFeed.DucksCount = 0

	created, err := db.InsertDuckFeedReport(duckFeed)
	assert.Error(t, err)
	assert.Nil(t, created)
}

func TestGetDuckFeeds(t *testing.T) {
	db := NewDatabase(logger)
	defer truncateTables(db)

	duckFeed := fixtures.DuckFeed()

	created, err := db.InsertDuckFeedReport(duckFeed)
	if !assert.NoError(t, err) {
		return
	}

	duckFeeds, err := db.GetDuckFeedsReport()
	if assert.NoError(t, err) && assert.NotNil(t, duckFeeds) {
		assert.Len(t, duckFeeds, 1)
		assertEqual(t, created, duckFeeds[0])
	}
}

func TestGetDuckFeedsWithNoRows(t *testing.T) {
	db := NewDatabase(logger)
	defer truncateTables(db)

	duckFeeds, err := db.GetDuckFeedsReport()
	if assert.NoError(t, err) && assert.NotNil(t, duckFeeds) {
		assert.Len(t, duckFeeds, 0)
	}
}

func TestGetDuckFeedById(t *testing.T) {
	db := NewDatabase(logger)
	defer truncateTables(db)

	created, err := db.InsertDuckFeedReport(fixtures.DuckFeed())
	if !assert.NoError(t, err) {
		return
	}

	duckFeed, err := db.GetDuckFeedReportById(created.ID)
	if assert.NoError(t, err) && assert.NotNil(t, duckFeed) {
		assertEqual(t, created, duckFeed)
	}
}

func TestGetDuckFeedByIdNoRows(t *testing.T) {
	db := NewDatabase(logger)
	defer truncateTables(db)

	duckFeed, err := db.GetDuckFeedReportById(1)
	assert.NoError(t, err)
	assert.Nil(t, duckFeed)
}

func TestInsertDuckFeedScheduleSucess(t *testing.T) {
	db := NewDatabase(logger)
	defer truncateTables(db)

	duckFeed := fixtures.DuckFeed()
	created, err := db.InsertDuckFeedReport(duckFeed)
	if !assert.NoError(t, err) {
		return
	}

	duckFeedSchedule := &duck_feed.DuckFeedSchedule{
		DuckFeedId: created.ID,
	}

	createdSchedule, err := db.InsertDuckFeedSchedule(duckFeedSchedule)
	if assert.NoError(t, err) && assert.NotNil(t, created) {
		assert.NotZero(t, createdSchedule.ID)
		assert.Equal(t, createdSchedule.DuckFeedId, created.ID)
	}
}

func TestInsertDuckFeedScheduleFail(t *testing.T) {
	db := NewDatabase(logger)
	defer truncateTables(db)

	duckFeedSchedule := &duck_feed.DuckFeedSchedule{
		DuckFeedId: 1,
	}

	createdSchedule, err := db.InsertDuckFeedSchedule(duckFeedSchedule)
	assert.Error(t, err)
	assert.Nil(t, createdSchedule)
}

func TestGetDuckFeedSchedules(t *testing.T) {
	db := NewDatabase(logger)
	defer truncateTables(db)

	duckFeed := fixtures.DuckFeed()

	created, err := db.InsertDuckFeedReport(duckFeed)
	if !assert.NoError(t, err) {
		return
	}

	duckFeedSchedule := &duck_feed.DuckFeedSchedule{
		DuckFeedId: created.ID,
	}

	createdSchedule, err := db.InsertDuckFeedSchedule(duckFeedSchedule)
	if !assert.NoError(t, err) {
		return
	}

	duckFeedSchedules, err := db.GetDuckFeedSchedules()
	if assert.NoError(t, err) && assert.NotNil(t, duckFeedSchedules) {
		assert.Len(t, duckFeedSchedules, 1)
		assert.Equal(t, createdSchedule.ID, duckFeedSchedules[0].ID)
		assert.Equal(t, createdSchedule.DuckFeedId, duckFeedSchedules[0].DuckFeedId)
	}
}

func TestGetDuckFeedSchedulesWithNoRows(t *testing.T) {
	db := NewDatabase(logger)
	defer truncateTables(db)

	duckFeedSchedules, err := db.GetDuckFeedSchedules()
	if assert.NoError(t, err) && assert.NotNil(t, duckFeedSchedules) {
		assert.Len(t, duckFeedSchedules, 0)
	}
}
