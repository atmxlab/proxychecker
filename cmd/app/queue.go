package app

import (
	stask "github.com/atmxlab/proxychecker/internal/service/task"
	"github.com/atmxlab/proxychecker/internal/service/task/handler"
	"github.com/atmxlab/proxychecker/pkg/queue"
)

func (a *App) initQueue() {
	addHandler := func(kind stask.Kind, checker handler.Checker) {
		a.Container().Entities().Queue().Add(
			queue.Kind(kind),
			handler.NewBaseCheckHandler(
				checker,
				a.Container().Ports().GetTaskAgg(),
				a.Container().Ports().SaveTaskAgg(),
			),
			queue.WithWorkerCount(a.Container().Config().Queue.QueueWorkerCount),
		)
	}

	addHandler(stask.KindCheckLatency, a.Container().Checkers().Latency())
	addHandler(stask.KindCheckGEO, a.Container().Checkers().GEO())
	addHandler(stask.KindCheckExternalIP, a.Container().Checkers().ExternalIP())
	addHandler(stask.KindCheckURL, a.Container().Checkers().URL())
	addHandler(stask.KindCheckURL, a.Container().Checkers().HTTPS())
	addHandler(stask.KindCheckURL, a.Container().Checkers().MITM())
	addHandler(stask.KindCheckURL, a.Container().Checkers().Type())
}
