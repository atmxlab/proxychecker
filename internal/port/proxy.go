package port

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
)

type InsertProxy interface {
	Execute(ctx context.Context, proxy ...*entity.Proxy) error
}

type GetProxy interface {
	Execute(ctx context.Context, id proxy.ID) (*entity.Proxy, error)
}

type GetProxies interface {
	Execute(ctx context.Context) ([]*entity.Proxy, error)
}
