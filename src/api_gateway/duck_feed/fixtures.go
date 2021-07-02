package duck_feed

import "time"

func duckFeed() *DuckFeed {
	return &DuckFeed{
		ParkName:     "Beacon Hill Park",
		Country:      "Canada",
		Food:         "Banana",
		FoodType:     "Vegetable",
		FoodQuantity: 2000,
		DucksCount:   20,
		Time:         time.Now(),
	}
}
