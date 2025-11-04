package port

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
)

type InsertProxy interface {
	Execute(ctx context.Context, proxy ...*entity.Proxy) error
}
