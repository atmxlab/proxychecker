package time

import (
	"context"
	"time"
)

type Provider interface {
	CurrentTime(ctx context.Context) time.Time
}

type NowProvider struct {
}

func NewNowProvider() NowProvider {
	return NowProvider{}
}

func (n NowProvider) CurrentTime(_ context.Context) time.Time {
	return time.Now()
}
