package test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/atmxlab/proxychecker/cmd/app"
	handlermocks "github.com/atmxlab/proxychecker/internal/service/task/handler/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type App struct {
	t     *testing.T
	ctx   context.Context
	mocks Mocks
	app   *app.App
}

func NewApp(t *testing.T) *App {
	ctx, cancel := context.WithCancel(context.Background())

	cfg := app.Config{
		Queue: app.Queue{
			QueueWorkerCount: 2,
			QueueBufferSize:  2,
		},
	}
	ctrl := gomock.NewController(t)
	mocks := Mocks{
		geoChecker:     handlermocks.NewMockChecker(ctrl),
		latencyChecker: handlermocks.NewMockChecker(ctrl),
	}

	cb := app.SetupContainerBuilder(cfg)
	cb.WithCheckers(func(cb *app.CheckersBuilder) {
		cb.
			GEO(mocks.geoChecker).
			Latency(mocks.latencyChecker)
	})
	a := app.NewApp(cb.Build())

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

	return &App{
		t:     t,
		ctx:   ctx,
		app:   a,
		mocks: mocks,
	}
}

func (a *App) Commands() app.Commands {
	return a.app.Container().Commands()
}

func (a *App) Ports() app.Ports {
	return a.app.Container().Ports()
}

func (a *App) Mocks() Mocks {
	return a.mocks
}

type Mocks struct {
	geoChecker     *handlermocks.MockChecker
	latencyChecker *handlermocks.MockChecker
}

func (m Mocks) GeoChecker() *handlermocks.MockChecker {
	return m.geoChecker
}

func (m Mocks) LatencyChecker() *handlermocks.MockChecker {
	return m.latencyChecker
}

func (a *App) WaitTasksTerminated() {
	require.Eventuallyf(
		a.t,
		func() bool {
			tasks, err := a.app.Container().Entities().Queue().GetNonTerminatedTasks(a.ctx)
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
