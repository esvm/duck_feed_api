package duck_feed

import "time"

type DuckFeed struct {
	ID           int    `json:"uid" pg:"id" validate:"-"`
	ParkName     string `json:"park_name" validate:"max=100"`
	Country      string `json:"country" validate:"max=100"`
	Food         string `json:"food" validate:"max=100"`
	FoodType     string `json:"food_type" validate:"max=100"`
	FoodQuantity int    `json:"food_quantity" validate:"required"`
	DucksCount   int    `json:"ducks_count" validate:"required"`

	Time      time.Time `json:"time" pg:"time" validate:"required"`
	CreatedAt time.Time `json:"datetime" pg:"created_at" validate:"-"`
}

type DuckFeedSchedule struct {
	ID         int `pg:"id" validate:"-"`
	DuckFeedId int `pg:"duck_feed_id" validate:"-"`
}
