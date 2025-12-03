package test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/atmxlab/proxychecker/cmd/app"
	"github.com/stretchr/testify/require"
)

type App struct {
	t   *testing.T
	ctx context.Context
	app *app.App
}

func NewApp(t *testing.T) *App {
	ctx, cancel := context.WithCancel(context.Background())

	cfg := app.Config{
		Queue: app.Queue{
			QueueWorkerCount: 2,
			QueueBufferSize:  2,
		},
	}

	a := app.NewApp(cfg)

	a.Init()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		require.NoError(t, a.Start(ctx))
	}()

	t.Cleanup(func() {
		cancel()
		wg.Wait()
	})

	return &App{t: t, app: a, ctx: ctx}
}

func (a *App) Commands() app.Commands {
	return a.app.Commands()
}

func (a *App) Ports() app.Ports {
	return a.app.Ports()
}

func (a *App) WaitTasksTerminated() {
	require.Eventuallyf(
		a.t,
		func() bool {
			tasks, err := a.app.Queue().GetNonTerminatedTasks(a.ctx)
			require.NoError(a.t, err)
			if len(tasks) == 0 {
				return true
			}

			a.t.Logf("wating tasks terminated with [%d] tasks: %v", len(tasks), tasks)
			return false
		},
		3*time.Second,
		300*time.Millisecond,
		"wating task terminated timed out",
	)
}
