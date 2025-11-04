package inmemory

import (
	"context"
	"sync"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
)

type InsertProxy struct {
	mu      sync.Mutex
	proxies map[proxy.ID]*entity.Proxy
}

func (r *InsertProxy) Execute(ctx context.Context, proxy ...*entity.Proxy) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, p := range proxy {
		r.proxies[p.ID()] = p
	}

	return nil
}

type InsertTask struct {
	mu    sync.Mutex
	tasks map[task.ID]*entity.Task
}

func (r *InsertTask) Execute(ctx context.Context, task ...*entity.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, t := range task {
		r.tasks[t.ID()] = t
	}

	return nil
}
