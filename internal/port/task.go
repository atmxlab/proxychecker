package port

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
)

type InsertTask interface {
	Execute(ctx context.Context, task ...*entity.Task) error
}
