package client

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
)

type Client interface {
	Get(ctx context.Context, url string) ([]byte, error)
}

type Factory interface {
	Create(p *entity.Proxy) Client
}
