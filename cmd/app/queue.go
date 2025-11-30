package app

import (
	"github.com/atmxlab/proxychecker/internal/details/checker/geo"
	"github.com/atmxlab/proxychecker/internal/details/checker/latency"
	"github.com/atmxlab/proxychecker/internal/details/factory"
	stask "github.com/atmxlab/proxychecker/internal/service/task"
	"github.com/atmxlab/proxychecker/internal/service/task/handler"
	"github.com/atmxlab/proxychecker/pkg/queue"
	"github.com/atmxlab/proxychecker/pkg/queue/inmemory"
)

func (a *App) initQueue() {
	a.queue = queue.NewQueue(inmemory.New(), a.cfg.Queue.QueueBufferSize)

	addHandler := func(kind stask.Kind, checker handler.Checker) {
		a.queue.Add(
			queue.Kind(kind),
			handler.NewBaseCheckHandler(
				checker,
				a.ports.getTaskAgg,
				a.ports.saveTaskAgg,
			),
			queue.WithWorkerCount(a.cfg.Queue.QueueWorkerCount),
		)
	}

	addHandler(stask.KindCheckLatency, latency.New())
	addHandler(stask.KindCheckGEO, geo.New(
		factory.NewClientFactory(),
		factory.NewIPApiFactory(),
	))
}
