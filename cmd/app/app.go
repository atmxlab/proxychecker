package app

import (
	"context"

	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/queue"
)

type App struct {
	cfg   Config
	queue *queue.Queue
	ports Ports
}

func NewApp(cfg Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Init() {
	a.initPorts()
	a.initQueue()
}

func (a *App) Start(ctx context.Context) error {
	if err := a.queue.Run(ctx); err != nil {
		return errors.Wrap(err, "starting queue")
	}

	return nil
}
