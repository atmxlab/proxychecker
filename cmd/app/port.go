package app

import (
	"github.com/atmxlab/proxychecker/internal/details/repository/inmemory"
	"github.com/atmxlab/proxychecker/internal/details/scheduler"
	"github.com/atmxlab/proxychecker/internal/port"
)

type Ports struct {
	insertProxy       port.InsertProxy
	getProxy          port.GetProxy
	getProxies        port.GetProxies
	insertTask        port.InsertTask
	updateTask        port.UpdateTask
	getTask           port.GetTask
	getTasks          port.GetTasks
	getTasksByGroupID port.GetTasksByGroupID
	scheduleTask      port.ScheduleTask
	getTaskAgg        port.GetTaskAgg
	saveTaskAgg       port.SaveTaskAgg
	runTx             port.RunTx
}

func (p Ports) GetTasksByGroupID() port.GetTasksByGroupID {
	return p.getTasksByGroupID
}

func (p Ports) GetTasks() port.GetTasks {
	return p.getTasks
}

func (p Ports) GetProxies() port.GetProxies {
	return p.getProxies
}

func (p Ports) InsertTask() port.InsertTask {
	return p.insertTask
}

func (p Ports) InsertProxy() port.InsertProxy {
	return p.insertProxy
}

func (p Ports) GetProxy() port.GetProxy {
	return p.getProxy
}

func (p Ports) UpdateTask() port.UpdateTask {
	return p.updateTask
}

func (p Ports) GetTask() port.GetTask {
	return p.getTask
}

func (p Ports) ScheduleTask() port.ScheduleTask {
	return p.scheduleTask
}

func (p Ports) GetTaskAgg() port.GetTaskAgg {
	return p.getTaskAgg
}

func (p Ports) SaveTaskAgg() port.SaveTaskAgg {
	return p.saveTaskAgg
}

func (p Ports) RunTx() port.RunTx {
	return p.runTx
}

func (a *App) initPorts() {
	a.ports.runTx = inmemory.NewRunTx()

	proxySharedState := inmemory.NewProxySharedState()
	a.ports.insertProxy = inmemory.NewInsertProxy(proxySharedState)
	a.ports.getProxy = inmemory.NewGetProxy(proxySharedState)
	a.ports.getProxies = inmemory.NewGetProxies(proxySharedState)

	taskSharedState := inmemory.NewTaskSharedState()
	a.ports.insertTask = inmemory.NewInsertTask(taskSharedState)
	a.ports.updateTask = inmemory.NewUpdateTask(taskSharedState)
	a.ports.getTask = inmemory.NewGetTask(taskSharedState)
	a.ports.getTasks = inmemory.NewGetTasks(taskSharedState)
	a.ports.getTasksByGroupID = inmemory.NewGetTasksByGroupID(taskSharedState)
	a.ports.getTaskAgg = inmemory.NewGetTaskAgg(
		a.ports.getTask,
		a.ports.getProxy,
	)
	a.ports.saveTaskAgg = inmemory.NewSaveTaskAgg(
		a.ports.updateTask,
	)
	a.ports.scheduleTask = scheduler.NewSchedulerTask(a.queue, a.timeProvider)
}
