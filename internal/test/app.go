package test

import (
	"testing"

	"github.com/atmxlab/proxychecker/cmd/app"
	"github.com/stretchr/testify/require"
)

type App struct {
	app *app.App
}

func NewApp(t *testing.T) *App {
	cfg := app.Config{
		Queue: app.Queue{
			QueueWorkerCount: 2,
			QueueBufferSize:  2,
		},
	}

	a := app.NewApp(cfg)

	a.Init()

	go func() {
		require.NoError(t, a.Start(t.Context()))
	}()

	return &App{a}
}

func (a *App) Commands() app.Commands {
	return a.app.Commands()
}
