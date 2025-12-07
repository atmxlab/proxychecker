package query

import (
	"context"
	"sort"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/port"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/validator"
)

type CheckResultInput struct {
	GroupID task.GroupID
}

func (i CheckResultInput) Validate() error {
	v := validator.New()

	if i.GroupID == "" {
		v.Failed("empty group id")
	}

	return v.Err()
}

type CheckResultOutput struct {
	TasksStatistic   CheckResultOutputTasksStatistic
	ProxiesStatistic CheckResultOutputProxiesStatistic
	Proxies          []CheckResultOutputProxyInfo
	IsChecked        bool
}

type CheckResultOutputTasksStatistic struct {
	Count        int
	SuccessCount int
	FailureCount int
	PendingCount int
}

type CheckResultOutputProxiesStatistic struct {
	Count        int
	CheckedCount int
	PendingCount int
}

type CheckResultOutputProxyInfo struct {
	Proxy          *entity.Proxy
	Tasks          []*entity.Task
	TasksStatistic CheckResultOutputTasksStatistic
	IsChecked      bool
}

type CheckResultQuery struct {
	getGroupAgg port.GetGroupAgg
}

func NewCheckResultQuery(getGroupAgg port.GetGroupAgg) *CheckResultQuery {
	return &CheckResultQuery{getGroupAgg: getGroupAgg}
}

func (q *CheckResultQuery) Execute(ctx context.Context, input CheckResultInput) (CheckResultOutput, error) {
	if err := input.Validate(); err != nil {
		return CheckResultOutput{}, errors.Wrap(err, "input.Validate")
	}

	agg, err := q.getGroupAgg.Execute(ctx, input.GroupID)
	if err != nil {
		return CheckResultOutput{}, errors.Wrap(err, "q.getGroupAgg")
	}

	proxiesInfo := make([]CheckResultOutputProxyInfo, 0)
	for _, px := range agg.Proxies() {
		proxiesInfo = append(proxiesInfo, CheckResultOutputProxyInfo{
			Proxy:     px.Proxy(),
			Tasks:     px.Tasks(),
			IsChecked: px.IsChecked(),
			TasksStatistic: CheckResultOutputTasksStatistic{
				Count:        px.TasksCount(),
				SuccessCount: px.SuccessTasksCount(),
				FailureCount: px.FailureTasksCount(),
				PendingCount: px.PendingTasksCount(),
			},
		})
	}

	sort.Slice(proxiesInfo, func(i, j int) bool {
		return proxiesInfo[i].TasksStatistic.SuccessCount > proxiesInfo[j].TasksStatistic.SuccessCount
	})

	return CheckResultOutput{
		TasksStatistic: CheckResultOutputTasksStatistic{
			Count:        agg.TasksCount(),
			SuccessCount: agg.SuccessTasksCount(),
			FailureCount: agg.FailureTasksCount(),
			PendingCount: agg.PendingTasksCount(),
		},
		ProxiesStatistic: CheckResultOutputProxiesStatistic{
			Count:        agg.ProxiesCount(),
			CheckedCount: agg.ProxiesCheckedCount(),
			PendingCount: agg.ProxiesPendingCount(),
		},
		Proxies:   proxiesInfo,
		IsChecked: agg.IsChecked(),
	}, nil
}
