package inmemory

import (
	"context"
	"sync"

	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/queue"
	"github.com/samber/lo"
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

	if err := i.updateTask(ctx, task); err != nil {
		return errors.Wrap(err, "i.updateTask")
	}

	return nil
}

func (i *InMemory) updateTask(ctx context.Context, task queue.Task) error {
	if _, ok := i.tasks[task.ID()]; ok {
		i.tasks[task.ID()] = task
		return nil
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

			if err := i.updateTask(ctx, t.SetStatus(queue.StatusRunning)); err != nil {
				return nil, errors.Wrap(err, "i.updateTask")
			}
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

func (i *InMemory) GetNonTerminatedTasks(ctx context.Context) ([]queue.Task, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	tasks := make([]queue.Task, 0)
	for _, t := range i.tasks {
		if lo.Contains(queue.NonTerminateStatus(), t.Status()) {
			tasks = append(tasks, t)
		}
	}

	return tasks, nil
}
