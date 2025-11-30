package command

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/port"
	stask "github.com/atmxlab/proxychecker/internal/service/task"
	"github.com/atmxlab/proxychecker/internal/service/task/payload"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/validator"
)

type CheckInput struct {
	OperationTime time.Time
	Proxies       []string
	Checkers      []checker.Kind
}

func (i CheckInput) Validate() error {
	v := validator.New()

	if len(i.Proxies) == 0 {
		v.Failed("empty proxies")
	}
	if len(i.Checkers) == 0 {
		v.Failed("empty checkers")
	}
	if i.OperationTime.IsZero() || i.OperationTime.Unix() <= 0 {
		v.Failed("invalid operation time")
	}

	for _, pUrl := range i.Proxies {
		u, err := url.Parse(pUrl)
		if err != nil {
			v.AddErr(errors.Wrapf(err, "invalid proxy URL: %s", pUrl))
		}

		if _, ok := proxy.TryProtocolFromString(u.Scheme); !ok {
			v.Failed(fmt.Sprintf("invalid proxy protocol: %s", u.Scheme))
		}
	}

	return v.Err()
}

type CheckOutput struct {
	taskGroupID task.GroupID
}

type CheckCommand struct {
	runTx        port.RunTx
	insertProxy  port.InsertProxy
	insertTask   port.InsertTask
	scheduleTask port.ScheduleTask
}

func NewCheckCommand(
	runTx port.RunTx,
	insertProxy port.InsertProxy,
	insertTask port.InsertTask,
	scheduleTask port.ScheduleTask,
) *CheckCommand {
	return &CheckCommand{
		runTx:        runTx,
		insertProxy:  insertProxy,
		insertTask:   insertTask,
		scheduleTask: scheduleTask,
	}
}

func (c *CheckCommand) Execute(ctx context.Context, input CheckInput) (*CheckOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid input")
	}

	proxies, err := c.makeProxies(input.Proxies, input.OperationTime)
	if err != nil {
		return nil, errors.Wrap(err, "c.makeProxies")
	}

	tasks, groupID := c.makeTasks(proxies, input.Checkers, input.OperationTime)

	qTasks, err := c.makeQueueTasks(tasks)
	if err != nil {
		return nil, errors.Wrap(err, "c.makeQueueTasks")
	}

	err = c.runTx.Execute(ctx, func(ctx context.Context) error {
		if err = c.insertProxy.Execute(ctx, proxies...); err != nil {
			return errors.Wrap(err, "c.insertProxy.Execute")
		}

		if err = c.insertTask.Execute(ctx, tasks...); err != nil {
			return errors.Wrap(err, "c.insertTask.Execute")
		}

		if err = c.scheduleTask.Execute(ctx, qTasks...); err != nil {
			return errors.Wrap(err, "c.scheduleTask.Execute")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "c.repo.InsertProxy")
	}

	return &CheckOutput{taskGroupID: groupID}, nil
}

func (c *CheckCommand) makeProxies(proxyUrls []string, ot time.Time) ([]*entity.Proxy, error) {
	proxies := make([]*entity.Proxy, 0, len(proxyUrls))

	for _, pUrl := range proxyUrls {
		u, err := url.Parse(pUrl)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid proxy url: %s", pUrl)
		}

		proxies = append(proxies, c.buildProxy(ot, u))
	}

	return proxies, nil
}

func (c *CheckCommand) makeTasks(
	proxies []*entity.Proxy,
	checkerKinds []checker.Kind,
	ot time.Time,
) ([]*entity.Task, task.GroupID) {
	tasks := make([]*entity.Task, 0, len(proxies)*len(checkerKinds))

	groupID := task.NewGroupID()
	for _, p := range proxies {
		for _, kind := range checkerKinds {
			tk := c.buildTask(
				ot,
				groupID,
				p.ID(),
				kind,
			)

			tasks = append(tasks, tk)
		}
	}

	return tasks, groupID
}

func (c *CheckCommand) makeQueueTasks(tasks []*entity.Task) (
	[]stask.Task,
	error,
) {
	qTasks := make([]stask.Task, 0, len(tasks))

	for _, tk := range tasks {
		qTasks = append(qTasks, &payload.Task{
			ID:          tk.ID(),
			CheckerKind: tk.CheckerKind(),
		})
	}

	return qTasks, nil
}

func (c *CheckCommand) buildProxy(operationTime time.Time, url *url.URL) *entity.Proxy {
	p := entity.
		NewProxyBuilder().
		ID(proxy.NewID()).
		URL(url.String()).
		Protocol(proxy.ProtocolFromString(url.Scheme)).
		CreatedAt(operationTime).
		UpdatedAt(operationTime).
		Build()

	return p
}

func (c *CheckCommand) buildTask(
	operationTime time.Time,
	groupID task.GroupID,
	proxyID proxy.ID,
	checkerKind checker.Kind,
) *entity.Task {
	t := entity.
		NewTaskBuilder().
		ID(task.NewID()).
		GroupID(groupID).
		ProxyID(proxyID).
		CheckerKind(checkerKind).
		CreatedAt(operationTime).
		UpdatedAt(operationTime).
		Build()

	return t
}
