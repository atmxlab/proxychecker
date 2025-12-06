package time

import (
	"context"
	"time"
)

type Provider interface {
	CurrentTime(ctx context.Context) time.Time
	Since(ctx context.Context, t time.Time) time.Duration
}
