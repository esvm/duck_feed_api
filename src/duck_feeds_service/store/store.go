package store

import (
	"github.com/esvm/duck_feed_api/src/duck_feed"
	"github.com/esvm/duck_feed_api/src/duck_feeds_service/store/postgres"
	"github.com/go-kit/kit/log"
)

type Store interface {
	InsertDuckFeedReport(*duck_feed.DuckFeed) (*duck_feed.DuckFeed, error)
	GetDuckFeedsReport() ([]*duck_feed.DuckFeed, error)
	GetDuckFeedReportById(int) (*duck_feed.DuckFeed, error)
	InsertDuckFeedSchedule(*duck_feed.DuckFeedSchedule) (*duck_feed.DuckFeedSchedule, error)
	GetDuckFeedSchedules() ([]*duck_feed.DuckFeedSchedule, error)
}

type basicStore struct {
	logger log.Logger
}

func New(logger log.Logger) Store {
	return basicStore{logger}
}

func (s basicStore) InsertDuckFeedReport(report *duck_feed.DuckFeed) (*duck_feed.DuckFeed, error) {
	database := postgres.NewDatabase(s.logger)
	return database.InsertDuckFeedReport(report)
}

func (s basicStore) GetDuckFeedsReport() ([]*duck_feed.DuckFeed, error) {
	database := postgres.NewDatabase(s.logger)
	return database.GetDuckFeedsReport()
}

func (s basicStore) GetDuckFeedReportById(id int) (*duck_feed.DuckFeed, error) {
	database := postgres.NewDatabase(s.logger)
	return database.GetDuckFeedReportById(id)
}

func (s basicStore) InsertDuckFeedSchedule(schedule *duck_feed.DuckFeedSchedule) (*duck_feed.DuckFeedSchedule, error) {
	database := postgres.NewDatabase(s.logger)
	return database.InsertDuckFeedSchedule(schedule)
}

func (s basicStore) GetDuckFeedSchedules() ([]*duck_feed.DuckFeedSchedule, error) {
	database := postgres.NewDatabase(s.logger)
	return database.GetDuckFeedSchedules()
}
