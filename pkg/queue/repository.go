package queue

import "context"

type Repository interface {
	UpdateTask(ctx context.Context, task Task) error
	AcquireTasks(ctx context.Context, kind Kind) ([]Task, error)
	PushTasks(ctx context.Context, task ...Task) error
}
