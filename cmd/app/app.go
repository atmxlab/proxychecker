package app

import (
	"context"

	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/waiter"
	"github.com/sirupsen/logrus"
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

func (a *App) WaitTasksTerminated(ctx context.Context, opts ...waiter.Option) error {
	return waiter.Wait(func() (bool, error) {
		if ctx.Err() != nil {
			return false, errors.Wrap(ctx.Err(), "ctx error")
		}

		tasks, err := a.Container().Entities().Queue().GetNonTerminatedTasks(ctx)
		if err != nil {
			return false, errors.Wrap(err, "container.Entities.Queue.GetNonTerminatedTasks")
		}
		if len(tasks) == 0 {
			return true, nil
		}

		logrus.Infof("waiting tasks terminated with: count: [%d]", len(tasks))
		return false, nil
	}, opts...)
}
