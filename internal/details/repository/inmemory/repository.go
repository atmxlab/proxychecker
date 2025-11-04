package inmemory

import (
	"context"
	"sync"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
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
