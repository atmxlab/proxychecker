package acceptance_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/service/command"
	"github.com/atmxlab/proxychecker/internal/test"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCheckGEO(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
		now = time.Now()
	)
	app := test.NewApp(t)

	expectedResult := task.Result{
		GEOResult: &task.GEOResult{
			ContinentCode: "Eurasia",
			Continent:     "Eurasia",
			CountryCode:   "USA",
			Country:       "USA",
			Region:        "New-York",
			City:          "New-York",
			Timezone:      "USA",
		},
	}

	app.Mocks().GeoChecker().
		EXPECT().
		Run(gomock.Any(), gomock.Any()).
		Return(expectedResult, nil)

	res, err := app.Commands().Check().Execute(
		ctx,
		command.CheckInput{
			OperationTime: now,
			Proxies:       []string{"https://proxy.io"},
			Checkers: []checker.Kind{
				checker.KindGEO,
			},
		},
	)
	require.NoError(t, err)
	app.WaitTasksTerminated()

	proxies, err := app.Ports().GetProxies().Execute(ctx)
	require.NoError(t, err)
	require.Len(t, proxies, 1)

	px := proxies[0]
	require.EqualValues(t, "https://proxy.io", px.URL())
	require.EqualValues(t, "https", px.Protocol())

	tasks, err := app.Ports().GetTasks().Execute(ctx)
	require.NoError(t, err)
	require.Len(t, tasks, 1)

	tk := tasks[0]
	require.Equal(t, res.TaskGroupID, tk.GroupID())
	require.Equal(t, task.StatusSuccess, tk.Status())
	require.Equal(t, checker.KindGEO, tk.CheckerKind())
	require.Equal(t, px.ID(), tk.ProxyID())
	require.Equal(t, expectedResult, tk.State().Result())
}

func TestCheckLatency(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
		now = time.Now()
	)
	app := test.NewApp(t)

	expectedResult := task.Result{
		LatencyResult: &task.LatencyResult{
			LatencyToProxy:  1 * time.Second,
			LatencyToTarget: 2 * time.Second,
		},
	}

	app.Mocks().LatencyChecker().
		EXPECT().
		Run(gomock.Any(), gomock.Any()).
		Return(expectedResult, nil)

	res, err := app.Commands().Check().Execute(
		ctx,
		command.CheckInput{
			OperationTime: now,
			Proxies:       []string{"https://proxy.io"},
			Checkers: []checker.Kind{
				checker.KindLatency,
			},
		},
	)
	require.NoError(t, err)
	app.WaitTasksTerminated()

	proxies, err := app.Ports().GetProxies().Execute(ctx)
	require.NoError(t, err)
	require.Len(t, proxies, 1)

	px := proxies[0]
	require.EqualValues(t, "https://proxy.io", px.URL())
	require.EqualValues(t, "https", px.Protocol())

	tasks, err := app.Ports().GetTasks().Execute(ctx)
	require.NoError(t, err)
	require.Len(t, tasks, 1)

	tk := tasks[0]
	require.Equal(t, res.TaskGroupID, tk.GroupID())
	require.Equal(t, task.StatusSuccess, tk.Status())
	require.Equal(t, checker.KindLatency, tk.CheckerKind())
	require.Equal(t, px.ID(), tk.ProxyID())
	require.Equal(t, expectedResult, tk.State().Result())
}

func TestCheckWithCheckerError(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
		now = time.Now()
	)
	app := test.NewApp(t)

	app.Mocks().LatencyChecker().
		EXPECT().
		Run(gomock.Any(), gomock.Any()).
		Return(task.Result{}, errors.ErrUnsupported)

	res, err := app.Commands().Check().Execute(
		ctx,
		command.CheckInput{
			OperationTime: now,
			Proxies:       []string{"https://proxy.io"},
			Checkers: []checker.Kind{
				checker.KindLatency,
			},
		},
	)
	require.NoError(t, err)
	app.WaitTasksTerminated()

	proxies, err := app.Ports().GetProxies().Execute(ctx)
	require.NoError(t, err)
	require.Len(t, proxies, 1)

	px := proxies[0]
	require.EqualValues(t, "https://proxy.io", px.URL())
	require.EqualValues(t, "https", px.Protocol())

	tasks, err := app.Ports().GetTasks().Execute(ctx)
	require.NoError(t, err)
	require.Len(t, tasks, 1)

	tk := tasks[0]
	require.Equal(t, res.TaskGroupID, tk.GroupID())
	require.Equal(t, task.StatusFailure, tk.Status())
	require.Equal(t, checker.KindLatency, tk.CheckerKind())
	require.Equal(t, px.ID(), tk.ProxyID())
}
