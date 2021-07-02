package duck_feeds_service

import (
	"context"

	"github.com/esvm/duck_feed_api/src/duck_feed"
	"github.com/esvm/duck_feed_api/src/duck_feeds_service/store"
	"github.com/go-kit/kit/log"
)

type DuckFeedService interface {
	InsertDuckFeedReport(context.Context, *duck_feed.DuckFeed) (*duck_feed.DuckFeed, error)
	GetDuckFeedsReport(context.Context) ([]*duck_feed.DuckFeed, error)
	GetDuckFeedReportById(context.Context, int) (*duck_feed.DuckFeed, error)
	InsertDuckFeedSchedule(context.Context, *duck_feed.DuckFeedSchedule) (*duck_feed.DuckFeedSchedule, error)
	GetDuckFeedSchedules(context.Context) ([]*duck_feed.DuckFeedSchedule, error)
}

type basicService struct {
	store store.Store
}

func NewDuckFeedService(logger log.Logger) DuckFeedService {
	var service DuckFeedService
	service = basicService{store.New(logger)}
	service = loggingMiddleware{logger, service}
	return service
}

func (s basicService) InsertDuckFeedReport(ctx context.Context, duckFeed *duck_feed.DuckFeed) (*duck_feed.DuckFeed, error) {
	return s.store.InsertDuckFeedReport(duckFeed)
}

func (s basicService) GetDuckFeedsReport(ctx context.Context) ([]*duck_feed.DuckFeed, error) {
	return s.store.GetDuckFeedsReport()
}

func (s basicService) GetDuckFeedReportById(ctx context.Context, id int) (*duck_feed.DuckFeed, error) {
	return s.store.GetDuckFeedReportById(id)
}

func (s basicService) InsertDuckFeedSchedule(ctx context.Context, schedule *duck_feed.DuckFeedSchedule) (*duck_feed.DuckFeedSchedule, error) {
	return s.store.InsertDuckFeedSchedule(schedule)
}

func (s basicService) GetDuckFeedSchedules(ctx context.Context) ([]*duck_feed.DuckFeedSchedule, error) {
	return s.store.GetDuckFeedSchedules()
}
