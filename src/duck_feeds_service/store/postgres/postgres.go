package postgres

import (
	"os"

	"github.com/esvm/duck_feed_api/src/duck_feed"
	"github.com/go-kit/kit/log"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
)

var connection *pg.DB

type BasicDatabase interface {
	GetConnection() (*pg.DB, error)
}

type DuckFeedDatabase interface {
	BasicDatabase

	InsertDuckFeedReport(*duck_feed.DuckFeed) (*duck_feed.DuckFeed, error)
	GetDuckFeedsReport() ([]*duck_feed.DuckFeed, error)
	GetDuckFeedReportById(int) (*duck_feed.DuckFeed, error)
	InsertDuckFeedSchedule(*duck_feed.DuckFeedSchedule) (*duck_feed.DuckFeedSchedule, error)
	GetDuckFeedSchedules() ([]*duck_feed.DuckFeedSchedule, error)
}

type duckFeedDatabase struct {
	BasicDatabase

	logger log.Logger
}

func NewDatabase(logger log.Logger) DuckFeedDatabase {
	var database DuckFeedDatabase = duckFeedDatabase{logger: logger}

	return database
}

func (d duckFeedDatabase) GetConnection() (*pg.DB, error) {
	if connection == nil {
		url := os.Getenv("DATABASE_URL")

		options, err := pg.ParseURL(url)
		if err != nil {
			return nil, err
		}

		connection = pg.Connect(options)
	}

	return connection, nil
}

func (d duckFeedDatabase) InsertDuckFeedReport(duckFeed *duck_feed.DuckFeed) (*duck_feed.DuckFeed, error) {
	db, err := d.GetConnection()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect with database")
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "Begin transaction failed")
	}

	err = tx.Insert(duckFeed)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "Insert Duck Feed query failed")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "Insert Duck Feed commit query failed")
	}

	return duckFeed, nil
}

func (d duckFeedDatabase) GetDuckFeedsReport() ([]*duck_feed.DuckFeed, error) {
	db, err := d.GetConnection()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect with database")
	}

	duckFeeds := []*duck_feed.DuckFeed{}
	if err := db.Model(&duckFeeds).Select(); err != nil {
		return nil, errors.Wrap(err, "Failed to select Duck Feed Reports")
	}

	return duckFeeds, nil
}

func (d duckFeedDatabase) GetDuckFeedReportById(id int) (*duck_feed.DuckFeed, error) {
	db, err := d.GetConnection()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect with database")
	}

	duckFeed := &duck_feed.DuckFeed{}
	if err := db.Model(duckFeed).Where("id = ?", id).Select(); err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Failed to select Duck Feed Report By ID")
	}

	return duckFeed, nil
}

func (d duckFeedDatabase) InsertDuckFeedSchedule(duckFeedSchedule *duck_feed.DuckFeedSchedule) (*duck_feed.DuckFeedSchedule, error) {
	db, err := d.GetConnection()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect with database")
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "Begin transaction failed")
	}

	err = tx.Insert(duckFeedSchedule)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "Insert Duck Feed Schedule query failed")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "Insert Duck Feed Schedule commit query failed")
	}

	return duckFeedSchedule, nil
}

func (d duckFeedDatabase) GetDuckFeedSchedules() ([]*duck_feed.DuckFeedSchedule, error) {
	db, err := d.GetConnection()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect with database")
	}

	duckFeedSchedules := []*duck_feed.DuckFeedSchedule{}
	if err := db.Model(&duckFeedSchedules).Select(); err != nil {
		return nil, errors.Wrap(err, "Failed to select Duck Feed Schedules")
	}

	return duckFeedSchedules, nil
}
