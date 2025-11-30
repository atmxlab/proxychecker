package inmemory

import (
	"context"
)

type RunTx struct {
}

func NewRunTx() *RunTx {
	return &RunTx{}
}

func (t RunTx) Execute(ctx context.Context, f func(ctx context.Context) error) error {
	return f(ctx)
}
