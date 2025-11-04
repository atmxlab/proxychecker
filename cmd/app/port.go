package app

import (
	"github.com/atmxlab/proxychecker/internal/details/repository/inmemory"
	"github.com/atmxlab/proxychecker/internal/port"
)

type Ports struct {
	insertProxy port.InsertProxy
	getProxy    port.GetProxy
	insertTask  port.InsertTask
	updateTask  port.UpdateTask
	getTask     port.GetTask
}

func (a *App) initPorts() {
	proxySharedState := inmemory.NewProxySharedState()
	a.ports.insertProxy = inmemory.NewInsertProxy(proxySharedState)
	a.ports.getProxy = inmemory.NewGetProxy(proxySharedState)

	taskSharedState := inmemory.NewTaskSharedState()
	a.ports.insertTask = inmemory.NewInsertTask(taskSharedState)
	a.ports.updateTask = inmemory.NewUpdateTask(taskSharedState)
	a.ports.getTask = inmemory.NewGetTask(taskSharedState)
}
