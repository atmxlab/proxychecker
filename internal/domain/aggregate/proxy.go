package aggregate

import (
	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
)

type Proxy struct {
	proxy *entity.Proxy
	tasks []*entity.Task
}

func NewProxy(proxy *entity.Proxy, tasks []*entity.Task) *Proxy {
	return &Proxy{proxy: proxy, tasks: tasks}
}

func (agg *Proxy) Proxy() *entity.Proxy {
	return agg.proxy
}

func (agg *Proxy) Tasks() []*entity.Task {
	return agg.tasks
}

func (agg *Proxy) IsPending() bool {
	return !agg.IsChecked()
}

func (agg *Proxy) IsChecked() bool {
	return agg.PendingTasksCount() == 0
}

func (agg *Proxy) TasksCount() int {
	return len(agg.tasks)
}

func (agg *Proxy) PendingTasksCount() int {
	return agg.tasksCountByStatus(task.StatusPending)
}

func (agg *Proxy) SuccessTasksCount() int {
	return agg.tasksCountByStatus(task.StatusSuccess)
}

func (agg *Proxy) FailureTasksCount() int {
	return agg.tasksCountByStatus(task.StatusFailure)
}

func (agg *Proxy) tasksCountByStatus(status task.Status) int {
	count := 0
	for _, tk := range agg.tasks {
		if tk.Status() == status {
			count++
		}
	}

	return count
}
