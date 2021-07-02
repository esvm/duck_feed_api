package duck_feed

import (
	"github.com/esvm/duck_feed_api/src/validation"
)

func ValidateDuckFeed(duckFeed *DuckFeed) error {
	if err := validation.Validate(duckFeed); err != nil {
		return err
	}

	return nil
}
