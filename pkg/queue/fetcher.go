package queue

import (
	"context"
	"time"

	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/sirupsen/logrus"
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
			logrus.Infof("context done: kind: [%d]", kind)
			return
		default:
			tasks, err := f.fetch(ctx, kind)
			switch {
			case err != nil:
				logrus.Errorf("fetcher.fetch: kind: [%d], err: [%s]", kind, err)
				time.Sleep(5 * time.Second) // TODO: cfg
			case len(tasks) == 0:
				logrus.Infof("fetcher.fetch: kind: [%d], no tasks", kind)
				time.Sleep(1 * time.Second) // TODO: cfg
			default:
				logrus.Infof("fetcher.fetch: kind: [%d], len: [%d]", kind, len(tasks))
				time.Sleep(2 * time.Second) // TODO: cfg
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
