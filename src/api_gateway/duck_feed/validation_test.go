package duck_feed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDuckFeedSucess(t *testing.T) {
	duckFeed := duckFeed()
	err := ValidateDuckFeed(duckFeed)
	assert.NoError(t, err)
}

func TestValidateDuckFeedFailCountry(t *testing.T) {
	duckFeed := duckFeed()
	duckFeed.Country = ""

	err := ValidateDuckFeed(duckFeed)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Country")
}

func TestValidateDuckFeedFailDucksCount(t *testing.T) {
	duckFeed := duckFeed()
	duckFeed.DucksCount = 0

	err := ValidateDuckFeed(duckFeed)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DucksCount")
}

func TestValidateDuckFeedFailFood(t *testing.T) {
	duckFeed := duckFeed()
	duckFeed.Food = ""

	err := ValidateDuckFeed(duckFeed)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Food")
}

func TestValidateDuckFeedFailFoodType(t *testing.T) {
	duckFeed := duckFeed()
	duckFeed.FoodType = ""

	err := ValidateDuckFeed(duckFeed)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "FoodType")
}

func TestValidateDuckFeedFailFoodQuantity(t *testing.T) {
	duckFeed := duckFeed()
	duckFeed.FoodQuantity = 0

	err := ValidateDuckFeed(duckFeed)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "FoodQuantity")
}

func TestValidateDuckFeedFailParkName(t *testing.T) {
	duckFeed := duckFeed()
	duckFeed.ParkName = ""

	err := ValidateDuckFeed(duckFeed)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ParkName")
}
