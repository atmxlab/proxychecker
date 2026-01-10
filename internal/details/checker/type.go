package checker

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type TypeChecker struct {
	clientFactory ClientFactory
	ipApiFactory  IPApiFactory
}

func NewTypeChecker(clientFactory ClientFactory, ipApiFactory IPApiFactory) *TypeChecker {
	return &TypeChecker{clientFactory: clientFactory, ipApiFactory: ipApiFactory}
}

func (c *TypeChecker) Run(ctx context.Context, agg *aggregate.Task) (task.Result, error) {
	cl := c.clientFactory.Create(client.WithProxyURL(agg.Proxy().URL()))
	ipApi := c.ipApiFactory.Create(cl)

	output, err := ipApi.Get(ctx)
	if err != nil {
		return task.Result{}, errors.Wrap(err, "ipApi.Get")
	}

	if output.Query == "" {
		return task.Result{}, errors.New("ipApi.Get returned empty query result")
	}

	t := proxy.TypeResidential

	if output.Mobile {
		t = proxy.TypeMobile
	}
	if output.Hosting {
		t = proxy.TypeDatacenter
	}

	return task.Result{
		TypeResult: &task.TypeResult{
			Type: t,
		},
	}, nil
}
