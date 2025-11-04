package app

import (
	"github.com/atmxlab/proxychecker/internal/service/task"
	"github.com/atmxlab/proxychecker/internal/service/task/handler"
	"github.com/atmxlab/proxychecker/pkg/queue"
	"github.com/atmxlab/proxychecker/pkg/queue/inmemory"
)

func (a *App) initQueue() {
	a.queue = queue.NewQueue(inmemory.New(), a.cfg.Queue.QueueBufferSize)

	addHandler := func(kind task.Kind, handler queue.Handler) {
		a.queue.Add(queue.Kind(kind), handler, queue.WithWorkerCount(a.cfg.Queue.QueueWorkerCount))
	}

	addHandler(task.KindCheckLatency, handler.NewCheckLatencyHandler())
	addHandler(task.KindCheckGEO, handler.NewCheckGEOHandler())
}
