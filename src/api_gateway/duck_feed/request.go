package duck_feed

import (
	"encoding/json"
	"time"

	"github.com/esvm/duck_feed_api/src/duck_feed"
	"github.com/labstack/echo"
)

type DuckFeed struct {
	ParkName     string    `json:"park_name" validate:"required,max=100"`
	Country      string    `json:"country" validate:"required,max=100"`
	Food         string    `json:"food" validate:"required,max=100"`
	FoodType     string    `json:"food_type" validate:"required,max=100"`
	FoodQuantity int       `json:"food_quantity" validate:"required"`
	DucksCount   int       `json:"ducks_count" validate:"required"`
	Time         time.Time `json:"time" validate:"required"`
}

func UnmarshalDuckFeedReport(ctx echo.Context) (*DuckFeed, error) {
	req := ctx.Request()

	duckFeed := &DuckFeed{}
	err := json.NewDecoder(req.Body).Decode(&duckFeed)

	return duckFeed, err
}

func ParseRequestDuckFeedToDomainDuckFeed(request *DuckFeed) *duck_feed.DuckFeed {
	return &duck_feed.DuckFeed{
		ParkName:     request.ParkName,
		Country:      request.Country,
		Food:         request.Food,
		FoodType:     request.FoodType,
		FoodQuantity: request.FoodQuantity,
		DucksCount:   request.DucksCount,
		Time:         request.Time,
	}
}
