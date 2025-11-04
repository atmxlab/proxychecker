package handler

import (
	"context"
	"encoding/json"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/port"
	"github.com/atmxlab/proxychecker/internal/service/task/payload"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/queue"
)

type Checker interface {
	Run(ctx context.Context, p *entity.Proxy) (task.Result, error)
}

type BaseCheckHandler struct {
	checker    Checker
	getProxy   port.GetProxy
	getTask    port.GetTask
	updateTask port.UpdateTask
}

func NewBaseCheckHandler(
	checker Checker,
	getProxy port.GetProxy,
	getTask port.GetTask,
	updateTask port.UpdateTask,
) *BaseCheckHandler {
	return &BaseCheckHandler{
		checker:    checker,
		getProxy:   getProxy,
		getTask:    getTask,
		updateTask: updateTask,
	}
}

func (c *BaseCheckHandler) Handle(ctx context.Context, qt queue.Task) error {
	var p payload.Task
	if err := json.Unmarshal(qt.Payload(), &p); err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}

	t, err := c.getTask.Execute(ctx, p.ID)
	if err != nil {
		return errors.Wrap(err, "getTask.Execute")
	}

	px, err := c.getProxy.Execute(ctx, t.ProxyID())
	if err != nil {
		return errors.Wrap(err, "getProxy.Execute")
	}

	res, err := c.checker.Run(ctx, px)
	if err != nil {
		return errors.Wrap(err, "checker.Run")
	}

	err = t.Modify(func(m *entity.TaskModifier) error {
		m.Success(res)
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "t.Modify")
	}

	if err = c.updateTask.Execute(ctx, t); err != nil {
		return errors.Wrap(err, "updateTask.Execute")
	}

	return nil
}
