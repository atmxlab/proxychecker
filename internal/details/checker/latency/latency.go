package latency

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
)

type Checker struct {
}

func New() *Checker {
	return &Checker{}
}

func (c Checker) Run(ctx context.Context, tAgg *aggregate.Task) (task.Result, error) {
	// TODO implement me
	panic("implement me")
}
