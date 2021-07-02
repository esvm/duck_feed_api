package middleware

import (
	"context"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type ContextFunc func() context.Context

type ContextInjectorConfig struct {
	Skipper     middleware.Skipper
	ContextFunc ContextFunc
}

var DefaultContextInjectorConfig = ContextInjectorConfig{
	Skipper: middleware.DefaultSkipper,
	ContextFunc: func() context.Context {
		return context.Background()
	},
}

func ContextInjector() echo.MiddlewareFunc {
	return ContextInjectorWithConfig(DefaultContextInjectorConfig)
}

func ContextInjectorWithConfig(config ContextInjectorConfig) echo.MiddlewareFunc {

	if config.Skipper == nil {
		config.Skipper = DefaultContextInjectorConfig.Skipper
	}

	if config.Skipper == nil {
		config.ContextFunc = DefaultContextInjectorConfig.ContextFunc
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			c.Set("context", config.ContextFunc())
			err := next(c)

			if err != nil {
				c.Error(err)
			}

			return err
		}
	}
}
