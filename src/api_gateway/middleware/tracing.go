package middleware

import (
	"context"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	opentracing "github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
)

type TracerConfig struct {
	Skipper middleware.Skipper
	Tracer  opentracing.Tracer
}

var DefaultTracerConfig = TracerConfig{
	Skipper: middleware.DefaultSkipper,
	Tracer:  nil,
}

func Tracer() echo.MiddlewareFunc {
	return TracerWithConfig(DefaultTracerConfig)
}

func TracerWithConfig(config TracerConfig) echo.MiddlewareFunc {

	if config.Skipper == nil {
		config.Skipper = DefaultTracerConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			ctx := c.Get("context").(context.Context)

			req := c.Request()
			res := c.Response()
			span, ctx := opentracing.StartSpanFromContext(ctx, strings.ToUpper(req.Method)+" "+req.URL.Path)
			span.SetTag("component", "HTTP")
			span.SetTag("http.method", req.Method)
			span.SetTag("http.status_code", res.Status)
			span.SetTag("http.url", req.Host+req.URL.String())
			c.Set("context", ctx)

			err := next(c)

			if err != nil {
				c.Error(err)
				span.LogFields(
					otlog.String("event", "error"),
					otlog.String("message", err.Error()),
				)
			}

			span.Finish()

			return err
		}
	}
}
