package app

import (
	"context"

	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/queue"
	"github.com/atmxlab/proxychecker/pkg/time"
)

type App struct {
	cfg          Config
	timeProvider time.Provider
	queue        *queue.Queue
	ports        Ports
	commands     Commands
}

func NewApp(cfg Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Queue() *queue.Queue {
	return a.queue
}

func (a *App) Cfg() Config {
	return a.cfg
}

func (a *App) TimeProvider() time.Provider {
	return a.timeProvider
}

func (a *App) Ports() Ports {
	return a.ports
}

func (a *App) Commands() Commands {
	return a.commands
}

func (a *App) Init() {
	a.timeProvider = time.NewNowProvider()
	a.initQueue()
	a.initPorts()
	a.initCommands()
}

func (a *App) Start(ctx context.Context) error {
	if err := a.queue.Run(ctx); err != nil {
		return errors.Wrap(err, "starting queue")
	}

	return nil
}
