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
		return nil, errors.Wrap(err, "get task")
	}

	p, err := g.getProxy.Execute(ctx, t.ProxyID())
	if err != nil {
		return nil, errors.Wrap(err, "get proxy")
	}

	return aggregate.NewTask(t, p), nil
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
