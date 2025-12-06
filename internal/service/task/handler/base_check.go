package handler

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/port"
	"github.com/atmxlab/proxychecker/internal/service/task/payload"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/queue"
)

//go:generate mock Checker
type Checker interface {
	Run(ctx context.Context, t *aggregate.Task) (task.Result, error)
}

type BaseCheckHandler struct {
	checker     Checker
	getTaskAgg  port.GetTaskAgg
	saveTaskAgg port.SaveTaskAgg
}

func NewBaseCheckHandler(
	checker Checker,
	getTaskAgg port.GetTaskAgg,
	saveTaskAgg port.SaveTaskAgg,
) *BaseCheckHandler {
	return &BaseCheckHandler{
		checker:     checker,
		getTaskAgg:  getTaskAgg,
		saveTaskAgg: saveTaskAgg,
	}
}

func (c *BaseCheckHandler) Handle(ctx context.Context, qt queue.Task) error {
	p, err := payload.NewTaskFromBytes(qt.Payload())
	if err != nil {
		return errors.Wrap(err, "payload.NewTaskFromBytes")
	}

	t, err := c.getTaskAgg.Execute(ctx, p.ID)
	if err != nil {
		return errors.Wrap(err, "getTaskAgg.Execute")
	}

	res, err := c.checker.Run(ctx, t)
	if err != nil {
		return errors.Wrap(err, "checker.Run")
	}

	if err = t.Success(res); err != nil {
		return errors.Wrap(err, "taskAgg.Success")
	}

	if err = c.saveTaskAgg.Execute(ctx, t); err != nil {
		return errors.Wrap(err, "saveTaskAgg.Execute")
	}

	return nil
}
