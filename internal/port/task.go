package port

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
)

type InsertTask interface {
	Execute(ctx context.Context, task ...*entity.Task) error
}

type UpdateTask interface {
	Execute(ctx context.Context, task ...*entity.Task) error
}

type GetTask interface {
	Execute(ctx context.Context, id task.ID) (*entity.Task, error)
}
