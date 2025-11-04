package queue

import (
	"context"
)

type worker struct {
	tasks   chan Task
	results chan result
	handler Handler
}

func newWorker(tasks chan Task, results chan result, handler Handler) *worker {
	return &worker{tasks: tasks, results: results, handler: handler}
}

func (w *worker) run(ctx context.Context) {
	for task := range w.tasks {
		err := w.handler.Handle(ctx, task)
		w.results <- newResult(err, task)
	}
}
