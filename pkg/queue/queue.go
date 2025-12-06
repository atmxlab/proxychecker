package queue

import (
	"context"
	"sync"

	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Queue struct {
	handlers   []handler
	repo       Repository
	bufferSize int
	wg         sync.WaitGroup
}

func NewQueue(
	repo Repository,
	bufferSize int,
) *Queue {
	return &Queue{
		repo:       repo,
		bufferSize: bufferSize,
	}
}

func (q *Queue) Add(k Kind, h Handler, opts ...Option) {
	opt := newOptions()
	for _, o := range opts {
		o(&opt)
	}

	for _, hh := range q.handlers {
		if hh.kind == k {
			panic(errors.Newf("queue handler already exists: kind: [%s]", k))
		}
	}

	q.handlers = append(q.handlers, newHandler(k, h, opt))
}

func (q *Queue) Run(ctx context.Context) error {
	for _, h := range q.handlers {
		logrus.Infof("run queue: kind: [%d], workers count: [%d]", h.kind, h.options.workerCount)

		tasksCh := make(chan Task, q.bufferSize)
		resultsCh := make(chan result, q.bufferSize)

		logrus.Infof("run tasks workers: kind: [%d], count: [%d]", h.kind, h.options.workerCount)
		q.wg.Add(1)
		go func() {
			defer q.wg.Done()
			defer close(resultsCh)
			q.runTasksWorkers(ctx, tasksCh, resultsCh, h)
		}()

		logrus.Infof("run results workers: kind: [%d], count: [%d]", h.kind, h.options.workerCount)
		q.wg.Add(1)
		go func() {
			defer q.wg.Done()
			q.runResultsWorkers(ctx, resultsCh, h)
		}()

		logrus.Infof("run fetcher: kind: [%d]", h.kind)
		// Достает таски и кладет в канал.
		f := newFetcher(tasksCh, q.repo)
		q.wg.Add(1)
		go func() {
			defer q.wg.Done()
			defer close(tasksCh)

			f.run(ctx, h.kind)
		}()
	}

	q.wg.Wait()

	return nil
}

func (q *Queue) runTasksWorkers(ctx context.Context, tasksCh chan Task, resultsCh chan result, h handler) {
	wg := sync.WaitGroup{}
	for range h.options.workerCount {
		wg.Add(1)
		// Обрабатывает все задачи и кладет результат в канал с результатами.
		tw := newTasksWorker(tasksCh, resultsCh, h.handler)
		go func() {
			defer wg.Done()
			tw.run(ctx)
		}()
	}

	wg.Wait()
}

func (q *Queue) runResultsWorkers(ctx context.Context, resultsCh chan result, h handler) {
	wg := sync.WaitGroup{}
	for range h.options.workerCount {
		wg.Add(1)
		// Обрабатывает результат выполнения.
		rw := newResultsWorker(resultsCh, q.repo)
		go func() {
			defer wg.Done()
			rw.run(ctx)
		}()
	}

	wg.Wait()
}

func (q *Queue) PushTasks(ctx context.Context, task ...Task) error {
	if err := q.repo.PushTasks(ctx, task...); err != nil {
		return errors.Wrap(err, "q.repo.PushTasks")
	}

	return nil
}

func (q *Queue) GetNonTerminatedTasks(ctx context.Context) ([]Task, error) {
	tasks, err := q.repo.GetNonTerminatedTasks(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "q.repo.GetNonTerminatedTasks")
	}

	return tasks, nil
}
