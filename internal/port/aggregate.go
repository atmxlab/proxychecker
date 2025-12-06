package port

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
)

type GetTaskAgg interface {
	Execute(ctx context.Context, id task.ID) (*aggregate.Task, error)
}

type GetTaskAggsByGroupID interface {
	Execute(ctx context.Context, groupID task.GroupID) ([]*aggregate.Task, error)
}

type SaveTaskAgg interface {
	Execute(ctx context.Context, agg *aggregate.Task) error
}
