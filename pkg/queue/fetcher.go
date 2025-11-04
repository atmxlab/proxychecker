package queue

import (
	"context"
	"log/slog"
	"time"

	"github.com/atmxlab/proxychecker/pkg/errors"
)

type fetcher struct {
	tasksCh chan Task
	repo    Repository
}

func newFetcher(tasksCh chan Task, repo Repository) *fetcher {
	return &fetcher{
		tasksCh: tasksCh,
		repo:    repo,
	}
}

func (f *fetcher) run(ctx context.Context, kind Kind) {
	for {
		select {
		case <-ctx.Done():
			slog.Info("context done: kind: [%d]", kind)
			return
		default:
			tasks, err := f.fetch(ctx, kind)
			if err != nil {
				slog.Error("fetcher.fetch: err: [%s]", err)
				time.Sleep(5 * time.Second)
			} else {
				for _, t := range tasks {
					f.tasksCh <- t
				}
			}
		}
	}
}

func (f *fetcher) fetch(ctx context.Context, kind Kind) ([]Task, error) {
	tasks, err := f.repo.AcquireTasks(ctx, kind)
	if err != nil {
		return nil, errors.Wrap(err, "f.repo.AcquireTasks")
	}

	return tasks, nil
}
