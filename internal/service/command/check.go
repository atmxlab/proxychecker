package command

import (
	"context"
	"fmt"
	url2 "net/url"
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/validator"
)

type CheckInput struct {
	operationTime time.Time
	proxies       []string
	checkerIDs    []checker.ID
}

func (i CheckInput) Validate() error {
	v := validator.New()

	if len(i.proxies) == 0 {
		v.Failed("empty proxies")
	}
	if len(i.checkerIDs) == 0 {
		v.Failed("empty checkers")
	}
	if i.operationTime.IsZero() || i.operationTime.Unix() <= 0 {
		v.Failed("invalid operation time")
	}

	for _, url := range i.proxies {
		u, err := url2.Parse(url)
		if err != nil {
			v.AddErr(errors.Wrapf(err, "invalid proxy URL: %s", url))
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

type TxRunner interface {
	Run(context.Context, func(ctx context.Context) error) error
}

type Repository interface {
	InsertProxy(ctx context.Context, proxy ...*entity.Proxy) error
	InsertTask(ctx context.Context, task ...*entity.Task) error
}

type CheckCommand struct {
	repo Repository
	tx   TxRunner
}

func (c *CheckCommand) Execute(ctx context.Context, input CheckInput) (*CheckOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid input")
	}

	proxies := make([]*entity.Proxy, len(input.proxies), len(input.proxies))

	for _, url := range input.proxies {
		u, err := url2.Parse(url)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid proxy url: %s", url)
		}

		proxies = append(proxies, c.buildProxy(input.operationTime, u))
	}

	tasks := make([]*entity.Task, len(proxies)*len(input.checkerIDs), len(proxies)*len(input.checkerIDs))

	groupID := task.NewGroupID()
	for _, p := range proxies {
		for _, checkerID := range input.checkerIDs {
			tasks = append(tasks, c.buildTask(
				input.operationTime,
				groupID,
				p.ID(),
				checkerID,
			))
		}
	}

	err := c.tx.Run(ctx, func(ctx context.Context) error {
		if err := c.repo.InsertProxy(ctx, proxies...); err != nil {
			return errors.Wrap(err, "c.repo.InsertProxy")
		}

		if err := c.repo.InsertTask(ctx, tasks...); err != nil {
			return errors.Wrap(err, "c.repo.InsertTask")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "c.repo.InsertProxy")
	}

	return &CheckOutput{taskGroupID: groupID}, nil
}

func (c *CheckCommand) buildProxy(operationTime time.Time, url *url2.URL) *entity.Proxy {
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
	checkerID checker.ID,
) *entity.Task {
	t := entity.
		NewTaskBuilder().
		ID(task.NewID()).
		GroupID(groupID).
		ProxyID(proxyID).
		CheckerID(checkerID).
		CreatedAt(operationTime).
		UpdatedAt(operationTime).
		Build()

	return t
}
