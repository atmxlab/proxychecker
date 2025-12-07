package aggregate

import "github.com/samber/lo"

type Group struct {
	proxies []*Proxy
}

func NewGroup(proxies []*Proxy) *Group {
	return &Group{proxies: proxies}
}

func (agg *Group) Proxies() []*Proxy {
	return agg.proxies
}

func (agg *Group) ProxiesCount() int {
	return len(agg.proxies)
}

func (agg *Group) ProxiesCheckedCount() int {
	return lo.SumBy(agg.proxies, func(item *Proxy) int {
		if item.IsChecked() {
			return 1
		}
		return 0
	})
}

func (agg *Group) ProxiesPendingCount() int {
	return lo.SumBy(agg.proxies, func(item *Proxy) int {
		if item.IsPending() {
			return 1
		}
		return 0
	})
}

func (agg *Group) TasksCount() int {
	return lo.SumBy(agg.proxies, func(item *Proxy) int {
		return item.TasksCount()
	})
}

func (agg *Group) SuccessTasksCount() int {
	return lo.SumBy(agg.proxies, func(item *Proxy) int {
		return item.SuccessTasksCount()
	})
}

func (agg *Group) FailureTasksCount() int {
	return lo.SumBy(agg.proxies, func(item *Proxy) int {
		return item.FailureTasksCount()
	})
}

func (agg *Group) PendingTasksCount() int {
	return lo.SumBy(agg.proxies, func(item *Proxy) int {
		return item.PendingTasksCount()
	})
}

func (agg *Group) IsChecked() bool {
	for _, proxy := range agg.proxies {
		if !proxy.IsChecked() {
			return false
		}
	}

	return true
}
