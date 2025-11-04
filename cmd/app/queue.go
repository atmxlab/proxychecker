package app

import (
	"github.com/atmxlab/proxychecker/internal/details/checker/geo"
	"github.com/atmxlab/proxychecker/internal/details/checker/latency"
	"github.com/atmxlab/proxychecker/internal/details/factory"
	"github.com/atmxlab/proxychecker/internal/service/task"
	"github.com/atmxlab/proxychecker/internal/service/task/handler"
	"github.com/atmxlab/proxychecker/pkg/queue"
	"github.com/atmxlab/proxychecker/pkg/queue/inmemory"
)

func (a *App) initQueue() {
	a.queue = queue.NewQueue(inmemory.New(), a.cfg.Queue.QueueBufferSize)

	addHandler := func(kind task.Kind, checker handler.Checker) {
		a.queue.Add(
			queue.Kind(kind),
			handler.NewBaseCheckHandler(
				checker,
				nil,
				nil,
				nil,
			),
			queue.WithWorkerCount(a.cfg.Queue.QueueWorkerCount),
		)
	}

	addHandler(task.KindCheckLatency, latency.New())
	addHandler(task.KindCheckGEO, geo.New(
		factory.NewClientFactory(),
		factory.NewIPApiFactory(),
	))

	// ctx := context.Background()
	//
	// px := entity.
	// 	NewProxyBuilder().
	// 	Build()
	//
	// tk := entity.
	// 	NewTaskBuilder().
	// 	Build()
	//
	// if err := a.ports.insertProxy.Execute(ctx, px); err != nil {
	// 	panic(err)
	// }
	//
	// if err := a.ports.insertTask.Execute(ctx, tk); err != nil {
	// 	panic(err)
	// }
	//
	// a.queue.PushTasks(ctx, queue.Task{
	//
	// })
}
