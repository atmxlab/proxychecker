package checker

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type ExternalIPChecker struct {
	clientFactory ClientFactory
	ipApiFactory  IPApiFactory
}

func NewExternalIPChecker(clientFactory ClientFactory, ipApiFactory IPApiFactory) *ExternalIPChecker {
	return &ExternalIPChecker{clientFactory: clientFactory, ipApiFactory: ipApiFactory}
}

func (c *ExternalIPChecker) Run(ctx context.Context, agg *aggregate.Task) (task.Result, error) {
	cl := c.clientFactory.Create(client.WithProxyURL(agg.Proxy().URL()))
	ipApi := c.ipApiFactory.Create(cl)

	output, err := ipApi.Get(ctx)
	if err != nil {
		return task.Result{}, errors.Wrap(err, "ipApi.Get")
	}

	if output.Query == "" {
		return task.Result{}, errors.New("ipApi.Get returned empty query result")
	}

	return task.Result{
		ExternalIPResult: &task.ExternalIPResult{
			IP: output.Query,
		},
	}, nil
}
