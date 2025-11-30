package port

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/service/task"
)

type ScheduleTask interface {
	Execute(ctx context.Context, task ...task.Task) error
}
