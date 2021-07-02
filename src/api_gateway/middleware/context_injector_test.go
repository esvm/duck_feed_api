package middleware

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestContextInjector(t *testing.T) {
	e := echo.New()
	handler := func(c echo.Context) error {
		return c.String(200, "test")
	}

	req := httptest.NewRequest(echo.GET, "/", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	h := ContextInjector()(handler)
	err := h(c)
	assert.NoError(t, err, "Handler should not return errors.")

	ctx := c.Get("context")
	assert.NotNil(t, ctx, "Context must be set.")
	assert.IsType(
		t,
		context.Background(),
		ctx,
		"Context must be of correct type.",
	)
}
