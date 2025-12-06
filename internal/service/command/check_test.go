package command_test

import (
	"context"
	"testing"
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/service/command"
	"github.com/atmxlab/proxychecker/internal/test"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCheck(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
		now = time.Now()
	)
	app := test.NewApp(t)

	cmd := command.CheckInput{
		OperationTime: now,
		Proxies: []string{
			"https://proxy-1.io",
			"https://proxy-2.io",
			"https://proxy-3.io",
			"https://proxy-4.io",
			"https://proxy-5.io",
		},
		Checkers: []checker.Kind{
			checker.KindGEO,
			checker.KindLatency,
		},
	}

	app.Mocks().GeoChecker().
		EXPECT().
		Run(gomock.Any(), gomock.Any()).
		Return(task.Result{}, nil).Times(len(cmd.Proxies))

	app.Mocks().LatencyChecker().
		EXPECT().
		Run(gomock.Any(), gomock.Any()).
		Return(task.Result{}, nil).Times(len(cmd.Proxies))

	res, err := app.Commands().Check().Execute(ctx, cmd)
	require.NoError(t, err)

	app.WaitTasksTerminated()

	tasks, err := app.Ports().GetTasksByGroupID().Execute(ctx, res.TaskGroupID)
	require.NoError(t, err)
	require.Len(t, tasks, len(cmd.Checkers)*len(cmd.Proxies))

	proxies, err := app.Ports().GetProxiesByTaskGroupID().Execute(ctx, res.TaskGroupID)
	require.NoError(t, err)
	require.Len(t, proxies, len(cmd.Proxies))

	for _, tk := range tasks {
		require.Equal(t, res.TaskGroupID, tk.GroupID())
		require.Equal(t, task.StatusSuccess, tk.Status())
		tk.TargetURL()
	}

	for _, px := range proxies {
		require.EqualValues(t, "https", px.Protocol())
	}

	require.ElementsMatch(t, cmd.Proxies, lo.Map(proxies, func(item *entity.Proxy, _ int) string {
		return item.URL()
	}))
}
