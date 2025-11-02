package cache

import "context"

type Input interface {
	Hash() string
}

type Service[I Input, R any] interface {
	Execute(ctx context.Context, input Input) (R, error)
}

// Infinite cache without TTL.
type Infinite[I Input, R any] struct {
	cache   map[string]R
	service Service[I, R]
}

func NewInfinite[I Input, R any](service Service[I, R]) *Infinite[I, R] {
	return &Infinite[I, R]{
		service: service,
		cache:   make(map[string]R),
	}
}

func (i *Infinite[I, R]) Execute(ctx context.Context, input Input) (R, error) {
	r, ok := i.cache[input.Hash()]
	if ok {
		return r, nil
	}

	return i.service.Execute(ctx, input)
}
