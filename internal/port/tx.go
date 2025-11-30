package port

import "context"

type RunTx interface {
	Execute(context.Context, func(ctx context.Context) error) error
}
