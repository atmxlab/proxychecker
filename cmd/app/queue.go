package app

import (
	"context"
	"time"

	"github.com/atmxlab/proxychecker/internal/details/checker/geo"
	"github.com/atmxlab/proxychecker/internal/details/checker/latency"
	"github.com/atmxlab/proxychecker/internal/details/factory"
	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	stask "github.com/atmxlab/proxychecker/internal/service/task"
	"github.com/atmxlab/proxychecker/internal/service/task/payload"

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
				a.ports.getProxy,
				a.ports.getTask,
				a.ports.updateTask,
			),
			queue.WithWorkerCount(a.cfg.Queue.QueueWorkerCount),
		)
	}

	addHandler(stask.KindCheckLatency, latency.New())
	addHandler(stask.KindCheckGEO, geo.New(
		factory.NewClientFactory(),
		factory.NewIPApiFactory(),
	))

	ctx := context.Background()

	px := entity.
		NewProxyBuilder().
		ID(proxy.NewID()).
		URL("http://checker.pro/proxychecker/proxy").
		CreatedAt(time.Now()).
		Protocol("http").
		Build()

	gid := task.NewGroupID()
	tk := entity.
		NewTaskBuilder().
		ID(task.NewID()).
		GroupID(gid).
		CreatedAt(time.Now()).
		ProxyID(px.ID()).
		TargetURL(task.TargetURL{
			URL: "http://ip-api.com/json",
		}).
		Build()

	if err := a.ports.insertProxy.Execute(ctx, px); err != nil {
		panic(err)
	}

	if err := a.ports.insertTask.Execute(ctx, tk); err != nil {
		panic(err)
	}

	p := payload.Task{
		ID: tk.ID(),
	}
	bytes, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	qTask := queue.
		NewTaskBuilder().
		Kind(queue.Kind(stask.KindCheckGEO)).
		ID(queue.NewID()).
		CreatedAt(time.Now()).
		Status(queue.StatusPending).
		Payload(bytes).
		Build()

	if err = a.queue.PushTasks(ctx, qTask); err != nil {
		panic(err)
	}
}
