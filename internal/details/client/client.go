package client

import (
	"context"
	"net/http"
)

type Client interface {
	Get(ctx context.Context, url string) ([]byte, error)
	Do(ctx context.Context, url string) (*http.Response, error)
}

type Factory interface {
	Create(opts ...Option) Client
}
