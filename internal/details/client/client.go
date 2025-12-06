package client

import (
	"context"
)

type Client interface {
	Get(ctx context.Context, url string) ([]byte, error)
}

type Factory interface {
	Create(opts ...Option) Client
}
