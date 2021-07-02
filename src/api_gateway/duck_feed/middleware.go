package duck_feed

import (
	"time"

	"github.com/labstack/echo"
)

func instrumentingMiddleware(endpoint string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			begin := time.Now()

			err := next(c)

			DuckFeedReportAPIRequestsTotal.With("endpoint", endpoint).Add(1)
			DuckFeedReportAPIRequestsDuration.With("endpoint", endpoint).Observe(time.Since(begin).Seconds())

			return err
		}
	}
}
