package queue

import (
	"context"
	"sync"
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

	q.handlers = append(q.handlers, newHandler(k, h, opt))
}

func (q *Queue) Run(ctx context.Context) error {
	for _, h := range q.handlers {
		tasksCh := make(chan Task, q.bufferSize)
		resultsCh := make(chan result, q.bufferSize)

		for range h.options.workerCount {
			// Обрабатывает все задачи и кладет результат в канал с результатами.
			tw := newTasksWorker(tasksCh, resultsCh, h.handler)
			q.wg.Add(1)
			go func() {
				defer q.wg.Done()
				defer close(resultsCh)

				tw.run(ctx)
			}()

			// Обрабатывает результат выполнения.
			rw := newResultsWorker(resultsCh, q.repo)
			q.wg.Add(1)
			go func() {
				defer q.wg.Done()

				rw.run(ctx)
			}()
		}

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
