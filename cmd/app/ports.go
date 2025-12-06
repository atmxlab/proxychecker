package app

import (
	"github.com/atmxlab/proxychecker/internal/port"
)

type Ports struct {
	insertProxy             port.InsertProxy
	getProxy                port.GetProxy
	getProxies              port.GetProxies
	getProxiesByTaskGroupID port.GetProxiesByTaskGroupID
	insertTask              port.InsertTask
	updateTask              port.UpdateTask
	getTask                 port.GetTask
	getTasks                port.GetTasks
	getTasksByGroupID       port.GetTasksByGroupID
	scheduleTask            port.ScheduleTask
	getTaskAgg              port.GetTaskAgg
	saveTaskAgg             port.SaveTaskAgg
	runTx                   port.RunTx
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

func (p Ports) GetProxiesByTaskGroupID() port.GetProxiesByTaskGroupID {
	return p.getProxiesByTaskGroupID
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

type PortsBuilder struct {
	c *Container
}

func newPortsBuilder(c *Container) *PortsBuilder {
	return &PortsBuilder{c: c}
}

func (pb *PortsBuilder) Container() *Container {
	return pb.c
}

func (pb *PortsBuilder) InsertProxy(p port.InsertProxy) *PortsBuilder {
	pb.c.ports.insertProxy = p
	return pb
}

func (pb *PortsBuilder) GetProxy(p port.GetProxy) *PortsBuilder {
	pb.c.ports.getProxy = p
	return pb
}

func (pb *PortsBuilder) GetProxies(p port.GetProxies) *PortsBuilder {
	pb.c.ports.getProxies = p
	return pb
}

func (pb *PortsBuilder) GetProxiesByTaskGroupID(p port.GetProxiesByTaskGroupID) *PortsBuilder {
	pb.c.ports.getProxiesByTaskGroupID = p
	return pb
}

func (pb *PortsBuilder) InsertTask(p port.InsertTask) *PortsBuilder {
	pb.c.ports.insertTask = p
	return pb
}

func (pb *PortsBuilder) UpdateTask(p port.UpdateTask) *PortsBuilder {
	pb.c.ports.updateTask = p
	return pb
}

func (pb *PortsBuilder) GetTask(p port.GetTask) *PortsBuilder {
	pb.c.ports.getTask = p
	return pb
}

func (pb *PortsBuilder) GetTasks(p port.GetTasks) *PortsBuilder {
	pb.c.ports.getTasks = p
	return pb
}

func (pb *PortsBuilder) GetTasksByGroupID(p port.GetTasksByGroupID) *PortsBuilder {
	pb.c.ports.getTasksByGroupID = p
	return pb
}

func (pb *PortsBuilder) ScheduleTask(p port.ScheduleTask) *PortsBuilder {
	pb.c.ports.scheduleTask = p
	return pb
}

func (pb *PortsBuilder) GetTaskAgg(p port.GetTaskAgg) *PortsBuilder {
	pb.c.ports.getTaskAgg = p
	return pb
}

func (pb *PortsBuilder) SaveTaskAgg(p port.SaveTaskAgg) *PortsBuilder {
	pb.c.ports.saveTaskAgg = p
	return pb
}

func (pb *PortsBuilder) RunTx(p port.RunTx) *PortsBuilder {
	pb.c.ports.runTx = p
	return pb
}
