package inmemory

import (
	"context"
)

type TxRunner struct {
}

func (t TxRunner) Run(ctx context.Context, f func(ctx context.Context) error) error {
	return f(ctx)
}
