package inmemory

import (
	"context"
	"sync"

	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/queue"
)

type InMemory struct {
	mu    sync.Mutex
	tasks map[queue.ID]queue.Task
}

func New() *InMemory {
	return &InMemory{
		tasks: make(map[queue.ID]queue.Task),
	}
}

func (i *InMemory) UpdateTask(ctx context.Context, task queue.Task) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if t, ok := i.tasks[task.ID()]; ok {
		i.tasks[task.ID()] = t
	}

	return errors.Newf("task not found: id: [%s]", task.ID())
}

func (i *InMemory) AcquireTasks(ctx context.Context, kind queue.Kind) ([]queue.Task, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	acquired := make([]queue.Task, 0)
	for _, t := range i.tasks {
		if t.Kind() == kind && t.Status() == queue.StatusPending {
			acquired = append(acquired, t)
		}
	}

	return acquired, nil
}

func (i *InMemory) PushTasks(ctx context.Context, task ...queue.Task) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	for _, t := range task {
		i.tasks[t.ID()] = t
	}

	return nil
}
