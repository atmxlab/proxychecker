package inmemory

import (
	"context"
	"sync"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/port"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/samber/lo"
)

type ProxySharedState struct {
	mu      sync.Mutex
	proxies map[proxy.ID]*entity.Proxy
}

func NewProxySharedState() *ProxySharedState {
	return &ProxySharedState{
		proxies: make(map[proxy.ID]*entity.Proxy),
	}
}

type InsertProxy struct {
	state *ProxySharedState
}

func NewInsertProxy(state *ProxySharedState) *InsertProxy {
	return &InsertProxy{state: state}
}

func (r *InsertProxy) Execute(ctx context.Context, proxy ...*entity.Proxy) error {
	r.state.mu.Lock()
	defer r.state.mu.Unlock()

	for _, p := range proxy {
		r.state.proxies[p.ID()] = p
	}

	return nil
}

type GetProxies struct {
	state *ProxySharedState
}

func NewGetProxies(state *ProxySharedState) *GetProxies {
	return &GetProxies{state: state}
}

func (r *GetProxies) Execute(ctx context.Context) ([]*entity.Proxy, error) {
	r.state.mu.Lock()
	defer r.state.mu.Unlock()

	return lo.Values(r.state.proxies), nil
}

type GetProxy struct {
	state *ProxySharedState
}

func NewGetProxy(state *ProxySharedState) *GetProxy {
	return &GetProxy{state: state}
}

func (r *GetProxy) Execute(ctx context.Context, id proxy.ID) (*entity.Proxy, error) {
	r.state.mu.Lock()
	defer r.state.mu.Unlock()
	if p, ok := r.state.proxies[id]; ok {
		return p, nil
	}

	return nil, errors.NotFoundf("proxy not found: id: [%s]", id)
}

type GetProxiesByTaskGroupID struct {
	state             *ProxySharedState
	getTasksByGroupID port.GetTasksByGroupID
}

func NewGetProxiesByTaskGroupID(
	state *ProxySharedState,
	getTasksByGroupID port.GetTasksByGroupID,
) *GetProxiesByTaskGroupID {
	return &GetProxiesByTaskGroupID{state: state, getTasksByGroupID: getTasksByGroupID}
}

func (r *GetProxiesByTaskGroupID) Execute(ctx context.Context, groupID task.GroupID) ([]*entity.Proxy, error) {
	r.state.mu.Lock()
	defer r.state.mu.Unlock()

	tasks, err := r.getTasksByGroupID.Execute(ctx, groupID)
	if err != nil {
		return nil, errors.Wrapf(err, "getTasksByGroupID.Execute: [%s]", groupID)
	}

	proxies := make(map[proxy.ID]*entity.Proxy)
	for _, tk := range tasks {
		proxies[tk.ProxyID()] = r.state.proxies[tk.ProxyID()]
	}

	return lo.Values(proxies), nil
}

type TaskSharedState struct {
	mu    sync.Mutex
	tasks map[task.ID]*entity.Task
}

func NewTaskSharedState() *TaskSharedState {
	return &TaskSharedState{
		tasks: make(map[task.ID]*entity.Task),
	}
}

type InsertTask struct {
	state *TaskSharedState
}

func NewInsertTask(state *TaskSharedState) *InsertTask {
	return &InsertTask{state: state}
}

func (r *InsertTask) Execute(ctx context.Context, task ...*entity.Task) error {
	r.state.mu.Lock()
	defer r.state.mu.Unlock()

	for _, t := range task {
		r.state.tasks[t.ID()] = t
	}

	return nil
}

type UpdateTask struct {
	state *TaskSharedState
}

func NewUpdateTask(state *TaskSharedState) *UpdateTask {
	return &UpdateTask{state: state}
}

func (r *UpdateTask) Execute(ctx context.Context, task ...*entity.Task) error {
	r.state.mu.Lock()
	defer r.state.mu.Unlock()

	for _, t := range task {
		r.state.tasks[t.ID()] = t
	}

	return nil
}

type GetTask struct {
	state *TaskSharedState
}

func NewGetTask(state *TaskSharedState) *GetTask {
	return &GetTask{state: state}
}

func (r *GetTask) Execute(ctx context.Context, id task.ID) (*entity.Task, error) {
	r.state.mu.Lock()
	defer r.state.mu.Unlock()
	if t, ok := r.state.tasks[id]; ok {
		return t, nil
	}

	return nil, errors.NotFoundf("task not found: id: [%s]", id)
}

type GetTasks struct {
	state *TaskSharedState
}

func NewGetTasks(state *TaskSharedState) *GetTasks {
	return &GetTasks{state: state}
}

func (r *GetTasks) Execute(ctx context.Context) ([]*entity.Task, error) {
	r.state.mu.Lock()
	defer r.state.mu.Unlock()

	return lo.Values(r.state.tasks), nil
}

type GetTasksByGroupID struct {
	state *TaskSharedState
}

func NewGetTasksByGroupID(state *TaskSharedState) *GetTasksByGroupID {
	return &GetTasksByGroupID{state: state}
}

func (r *GetTasksByGroupID) Execute(ctx context.Context, groupID task.GroupID) ([]*entity.Task, error) {
	r.state.mu.Lock()
	defer r.state.mu.Unlock()

	tasks := make([]*entity.Task, 0)
	for _, t := range r.state.tasks {
		if t.GroupID() == groupID {
			tasks = append(tasks, t)
		}
	}

	return tasks, nil
}
