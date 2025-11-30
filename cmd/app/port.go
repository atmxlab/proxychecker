package app

import (
	"github.com/atmxlab/proxychecker/internal/details/repository/inmemory"
	"github.com/atmxlab/proxychecker/internal/details/scheduler"
	"github.com/atmxlab/proxychecker/internal/port"
)

type Ports struct {
	insertProxy  port.InsertProxy
	getProxy     port.GetProxy
	insertTask   port.InsertTask
	updateTask   port.UpdateTask
	getTask      port.GetTask
	scheduleTask port.ScheduleTask
	getTaskAgg   port.GetTaskAgg
	saveTaskAgg  port.SaveTaskAgg
	runTx        port.RunTx
}

func (a *App) initPorts() {
	a.ports.runTx = inmemory.NewRunTx()

	proxySharedState := inmemory.NewProxySharedState()
	a.ports.insertProxy = inmemory.NewInsertProxy(proxySharedState)
	a.ports.getProxy = inmemory.NewGetProxy(proxySharedState)

	taskSharedState := inmemory.NewTaskSharedState()
	a.ports.insertTask = inmemory.NewInsertTask(taskSharedState)
	a.ports.updateTask = inmemory.NewUpdateTask(taskSharedState)
	a.ports.getTask = inmemory.NewGetTask(taskSharedState)
	a.ports.getTaskAgg = inmemory.NewGetTaskAgg(
		a.ports.getTask,
		a.ports.getProxy,
	)
	a.ports.saveTaskAgg = inmemory.NewSaveTaskAgg(
		a.ports.updateTask,
	)
	a.ports.scheduleTask = scheduler.NewSchedulerTask(a.queue, a.timeProvider)
}
