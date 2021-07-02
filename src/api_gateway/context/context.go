package context

import (
	"context"

	"github.com/labstack/echo"
)

func GetContext(ctx echo.Context) context.Context {
	c := ctx.Get("context")
	if c == nil {
		return context.Background()
	}
	return c.(context.Context)
}
