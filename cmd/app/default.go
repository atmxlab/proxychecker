package app

import (
	"github.com/atmxlab/proxychecker/internal/details/checker/geo"
	"github.com/atmxlab/proxychecker/internal/details/checker/latency"
	"github.com/atmxlab/proxychecker/internal/details/factory"
	"github.com/atmxlab/proxychecker/internal/details/repository/inmemory"
	"github.com/atmxlab/proxychecker/internal/details/scheduler"
	"github.com/atmxlab/proxychecker/internal/service/command"
	"github.com/atmxlab/proxychecker/pkg/queue"
	queuerepo "github.com/atmxlab/proxychecker/pkg/queue/inmemory"
	"github.com/atmxlab/proxychecker/pkg/time"
)

func SetupContainerBuilder(cfg Config) *ContainerBuilder {
	return NewContainerBuilder().
		WithConfig(cfg).
		WithEntities(func(eb *EntitiesBuilder) {
			eb.
				TimeProvider(time.NewNowProvider()).
				Queue(queue.NewQueue(queuerepo.New(), cfg.Queue.QueueBufferSize)).
				ClientFactory(factory.NewClientFactory()).
				IpAPIFactory(factory.NewIPApiFactory())
		}).
		WithPorts(func(pb *PortsBuilder) {
			pb.RunTx(inmemory.NewRunTx())

			taskSharedState := inmemory.NewTaskSharedState()
			getTasksByGroupID := inmemory.NewGetTasksByGroupID(taskSharedState)
			pb.
				InsertTask(inmemory.NewInsertTask(taskSharedState)).
				UpdateTask(inmemory.NewUpdateTask(taskSharedState)).
				GetTask(inmemory.NewGetTask(taskSharedState)).
				GetTasks(inmemory.NewGetTasks(taskSharedState)).
				GetTasksByGroupID(getTasksByGroupID)

			proxySharedState := inmemory.NewProxySharedState()
			pb.
				InsertProxy(inmemory.NewInsertProxy(proxySharedState)).
				GetProxy(inmemory.NewGetProxy(proxySharedState)).
				GetProxies(inmemory.NewGetProxies(proxySharedState)).
				GetProxiesByTaskGroupID(inmemory.NewGetProxiesByTaskGroupID(
					proxySharedState,
					getTasksByGroupID,
				))

			pb.
				GetTaskAgg(inmemory.NewGetTaskAgg(
					pb.Container().Ports().GetTask(),
					pb.Container().Ports().GetProxy(),
				)).
				SaveTaskAgg(
					inmemory.NewSaveTaskAgg(pb.Container().Ports().UpdateTask()),
				)

			pb.ScheduleTask(scheduler.NewSchedulerTask(
				pb.Container().Entities().Queue(),
				pb.Container().Entities().TimeProvider(),
			))
		}).
		WithCheckers(func(cb *CheckersBuilder) {
			cb.
				GEO(geo.New(
					cb.Container().Entities().ClientFactory(),
					cb.Container().Entities().IpApiFactory(),
				)).
				Latency(latency.New(
					cb.Container().Entities().ClientFactory(),
					cb.Container().Entities().IpApiFactory(),
					cb.Container().Entities().TimeProvider(),
				))
		}).
		WithCommands(func(pb *CommandsBuilder) {
			pb.Check(command.NewCheckCommand(
				pb.Container().Ports().RunTx(),
				pb.Container().Ports().InsertProxy(),
				pb.Container().Ports().InsertTask(),
				pb.Container().Ports().ScheduleTask(),
			))
		})
}
