package time

import (
	"context"
	"time"
)

type NowProvider struct {
}

func NewNowProvider() NowProvider {
	return NowProvider{}
}

func (n NowProvider) CurrentTime(_ context.Context) time.Time {
	return time.Now()
}

func (n NowProvider) Since(ctx context.Context, t time.Time) time.Duration {
	return n.CurrentTime(ctx).Sub(t)
}
