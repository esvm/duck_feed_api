package fixtures

import (
	"time"

	"github.com/esvm/duck_feed_api/src/duck_feed"
)

func DuckFeed() *duck_feed.DuckFeed {
	return &duck_feed.DuckFeed{
		ParkName:     "Beacon Hill Park",
		Country:      "Canada",
		Food:         "Banana",
		FoodType:     "Vegetable",
		FoodQuantity: 2000,
		DucksCount:   20,
		Time:         time.Now(),
	}
}
