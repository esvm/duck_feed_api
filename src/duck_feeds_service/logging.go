package duck_feeds_service

import (
	"context"
	"time"

	"github.com/esvm/duck_feed_api/src/duck_feed"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type loggingMiddleware struct {
	logger log.Logger
	next   DuckFeedService
}

func (mw loggingMiddleware) InsertDuckFeedReport(ctx context.Context, duckFeed *duck_feed.DuckFeed) (*duck_feed.DuckFeed, error) {
	begin := time.Now()

	created, err := mw.next.InsertDuckFeedReport(ctx, duckFeed)
	arguments := []interface{}{"method", "InsertDuckFeedReport", "err", err, "took", time.Since(begin)}
	level.Debug(mw.logger).Log(arguments...)

	return created, err
}

func (mw loggingMiddleware) GetDuckFeedsReport(ctx context.Context) ([]*duck_feed.DuckFeed, error) {
	begin := time.Now()

	res, err := mw.next.GetDuckFeedsReport(ctx)
	arguments := []interface{}{"method", "GetDuckFeedsReport", "err", err, "took", time.Since(begin)}
	level.Debug(mw.logger).Log(arguments...)

	return res, err
}

func (mw loggingMiddleware) GetDuckFeedReportById(ctx context.Context, id int) (*duck_feed.DuckFeed, error) {
	begin := time.Now()

	res, err := mw.next.GetDuckFeedReportById(ctx, id)
	arguments := []interface{}{
		"method", "GetDuckFeedReportById",
		"id", id,
		"err", err,
		"took", time.Since(begin),
	}
	level.Debug(mw.logger).Log(arguments...)

	return res, err
}

func (mw loggingMiddleware) InsertDuckFeedSchedule(ctx context.Context, duckFeedSchedule *duck_feed.DuckFeedSchedule) (*duck_feed.DuckFeedSchedule, error) {
	begin := time.Now()

	created, err := mw.next.InsertDuckFeedSchedule(ctx, duckFeedSchedule)
	arguments := []interface{}{"method", "InsertDuckFeedSchedule", "err", err, "took", time.Since(begin)}
	level.Debug(mw.logger).Log(arguments...)

	return created, err
}

func (mw loggingMiddleware) GetDuckFeedSchedules(ctx context.Context) ([]*duck_feed.DuckFeedSchedule, error) {
	begin := time.Now()

	res, err := mw.next.GetDuckFeedSchedules(ctx)
	arguments := []interface{}{"method", "GetDuckFeedSchedules", "err", err, "took", time.Since(begin)}
	level.Debug(mw.logger).Log(arguments...)

	return res, err
}
