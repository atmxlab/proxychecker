package app

import (
	"github.com/atmxlab/proxychecker/internal/service/task"
	"github.com/atmxlab/proxychecker/internal/service/task/handler"
	"github.com/atmxlab/proxychecker/pkg/queue"
)

func (a *App) initQueue() {
	a.queue = queue.NewQueue(1000) // TODO: cfg

	addHandler := func(kind task.Kind, handler queue.Handler) {
		a.queue.Add(queue.Kind(kind), handler, queue.WithWorkerCount(10)) // TODO: cfg
	}

	addHandler(task.KindCheckLatency, handler.NewCheckLatencyHandler())
	addHandler(task.KindCheckGEO, handler.NewCheckGEOHandler())
}
