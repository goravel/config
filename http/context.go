package http

import (
	"fmt"

	"github.com/gookit/color"

	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/http"
)

type Driver string

type Context struct {
	http.Context
	config config.Config
}

func Background() http.Context {
	c := NewContext(ConfigFacade)
	return c.Context
}

func NewContext(config config.Config) *Context {
	defaultDriver := config.GetString("http.driver")
	if defaultDriver == "" {
		color.Redln("[http] please set default driver")

		return nil
	}

	driver, err := NewDriver(config, defaultDriver)
	if err != nil {
		color.Redf("[http] %s\n", err)

		return nil
	}

	return &Context{
		Context: driver,
		config:  config,
	}
}

func NewDriver(config config.Config, driver string) (http.Context, error) {
	context, ok := config.Get("http.drivers." + driver + ".http").(http.Context)
	if ok {
		return context, nil
	}

	contextCallback, ok := config.Get("http.drivers." + driver + ".http").(func() (http.Context, error))
	if ok {
		return contextCallback()
	}

	return nil, fmt.Errorf("init http driver fail: http must be implement http.Context or func() (http.Context, error)")
}
