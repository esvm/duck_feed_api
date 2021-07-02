package middleware

import (
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := uuid.NewV4().String()
			c.Set("request_id", requestID)

			res := next(c)

			if err, ok := res.(*echo.HTTPError); ok {
				err.Message = map[string]interface{}{
					"message":    err.Message,
					"request_id": requestID,
				}
			}

			return res
		}
	}
}
