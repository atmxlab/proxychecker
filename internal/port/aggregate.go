package port

import (
	"context"
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
)

type AcquireTasks interface {
	Execute(ctx context.Context, deadline time.Time) ([]*aggregate.Task, error)
}
