package inmemory

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/port"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type GetTaskAgg struct {
	getTask  port.GetTask
	getProxy port.GetProxy
}

func NewGetTaskAgg(getTask port.GetTask, getProxy port.GetProxy) *GetTaskAgg {
	return &GetTaskAgg{getTask: getTask, getProxy: getProxy}
}

func (g *GetTaskAgg) Execute(ctx context.Context, id task.ID) (*aggregate.Task, error) {
	t, err := g.getTask.Execute(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "getTask.Execute")
	}

	p, err := g.getProxy.Execute(ctx, t.ProxyID())
	if err != nil {
		return nil, errors.Wrap(err, "getProxy.Execute")
	}

	return aggregate.NewTask(t, p), nil
}

type GetTaskAggsByGroupID struct {
	getTasksByGroupID port.GetTasksByGroupID
	getProxy          port.GetProxy
}

func NewGetTaskAggsByGroupID(getTasksByGroupID port.GetTasksByGroupID, getProxy port.GetProxy) *GetTaskAggsByGroupID {
	return &GetTaskAggsByGroupID{getTasksByGroupID: getTasksByGroupID, getProxy: getProxy}
}

func (g *GetTaskAggsByGroupID) Execute(ctx context.Context, groupID task.GroupID) ([]*aggregate.Task, error) {
	tasks, err := g.getTasksByGroupID.Execute(ctx, groupID)
	if err != nil {
		return nil, errors.Wrap(err, "getTasksByGroupID.Execute")
	}

	aggs := make([]*aggregate.Task, 0, len(tasks))
	for _, tk := range tasks {
		px, err := g.getProxy.Execute(ctx, tk.ProxyID())
		if err != nil {
			return nil, errors.Wrap(err, "getProxy.Execute")
		}
		aggs = append(aggs, aggregate.NewTask(tk, px))
	}

	return aggs, nil
}

type SaveTaskAgg struct {
	updateTask port.UpdateTask
}

func NewSaveTaskAgg(updateTask port.UpdateTask) *SaveTaskAgg {
	return &SaveTaskAgg{updateTask: updateTask}
}

func (s *SaveTaskAgg) Execute(ctx context.Context, agg *aggregate.Task) error {
	if err := s.updateTask.Execute(ctx, agg.Task()); err != nil {
		return errors.Wrap(err, "s.updateTask.Execute")
	}

	return nil
}
