package queue

import "context"

type fetcher struct {
}

func newFetcher() *fetcher {
	return &fetcher{}
}

func (f *fetcher) fetch(ctx context.Context, kind Kind) ([]Task, error) {
	return nil, nil
}
