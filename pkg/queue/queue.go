package queue

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

type Queue struct {
	handlers   []handler
	fetcher    *fetcher
	bufferSize int
	wg         sync.WaitGroup
}

func NewQueue(bufferSize int) *Queue {
	return &Queue{
		bufferSize: bufferSize,
		fetcher:    newFetcher(),
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
		tasksCh := make(chan Task)
		resultsCh := make(chan result)

		for range h.options.workerCount {
			w := newWorker(tasksCh, resultsCh, h.handler)
			q.wg.Add(1)
			go func() {
				defer q.wg.Done()
				defer close(resultsCh)

				w.run(ctx)
			}()
		}

		q.wg.Add(1)
		// Запускается обработка для одного вида задачи.
		go func() {
			defer q.wg.Done()
			defer close(tasksCh)

			for {
				select {
				case <-ctx.Done():
					slog.Info("context done: kind: [%d]", h.kind)
					return
				default:
					tasks, err := q.fetcher.fetch(ctx, h.kind)
					if err != nil {
						slog.Error("fetcher.fetch: err: [%s]", err.Error())
						time.Sleep(5 * time.Second)
					} else {
						for _, t := range tasks {
							tasksCh <- t
						}
					}
				}
			}
		}()
	}

	q.wg.Wait()

	return nil
}
