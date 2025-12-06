package app

import (
	"context"

	"github.com/atmxlab/proxychecker/pkg/errors"
)

type App struct {
	c *Container
}

func NewApp(c *Container) *App {
	return &App{
		c: c,
	}
}

func (a *App) Container() *Container {
	return a.c
}

func (a *App) Init() {
	a.initQueue()
}

func (a *App) Start(ctx context.Context) error {
	if err := a.Container().Entities().Queue().Run(ctx); err != nil {
		return errors.Wrap(err, "starting queue")
	}

	return nil
}
