package queue

import (
	"context"
	"log/slog"

	"github.com/atmxlab/proxychecker/pkg/errors"
)

type tasksWorker struct {
	tasks   chan Task
	results chan result
	handler Handler
}

func newTasksWorker(tasks chan Task, results chan result, handler Handler) *tasksWorker {
	return &tasksWorker{tasks: tasks, results: results, handler: handler}
}

func (w *tasksWorker) run(ctx context.Context) {
	for task := range w.tasks {
		err := w.handler.Handle(ctx, task)
		w.results <- newResult(err, task)
	}
}

type resultsWorker struct {
	results chan result
	repo    Repository
}

func newResultsWorker(results chan result, repo Repository) *resultsWorker {
	return &resultsWorker{results: results, repo: repo}
}

func (w *resultsWorker) run(ctx context.Context) {
	for res := range w.results {
		switch {
		case res.err == nil:
			res.task.status = StatusSuccess
		case errors.Is(res.err, ErrNonRetriable):
			res.task.status = StatusFailure
		default:
			res.task.status = StatusPending
		}

		if err := w.repo.UpdateTask(ctx, res.task); err != nil {
			slog.Error("w.repo.UpdateTask: err: [%s]", err)
		}
	}
}
