package app

import "github.com/atmxlab/proxychecker/pkg/queue"

type App struct {
	cfg   Config
	queue *queue.Queue
}

func NewApp(cfg Config) *App {
	return &App{
		cfg: cfg,
	}
}
