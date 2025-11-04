package pg

import (
	"context"

	"github.com/atmxlab/proxychecker/pkg/queue"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r Repository) UpdateTask(ctx context.Context, task queue.Task) error {
	// TODO implement me
	panic("implement me")
}

func (r Repository) AcquireTasks(ctx context.Context, kind queue.Kind) ([]queue.Task, error) {
	// TODO implement me
	panic("implement me")
}
