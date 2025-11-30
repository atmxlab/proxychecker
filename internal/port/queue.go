package port

import (
	"context"

	"github.com/atmxlab/proxychecker/pkg/queue"
)

type ScheduleTask interface {
	Execute(ctx context.Context, task ...queue.Task) error
}
