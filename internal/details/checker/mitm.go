package checker

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/details/service/httpbin"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type MITMChecker struct {
	clientFactory  ClientFactory
	httpBinFactory HTTPBinFactory
}

func NewMITMChecker(clientFactory ClientFactory, httpBinFactory HTTPBinFactory) *MITMChecker {
	return &MITMChecker{clientFactory: clientFactory, httpBinFactory: httpBinFactory}
}

func (c *MITMChecker) Run(ctx context.Context, agg *aggregate.Task) (task.Result, error) {
	cl := c.clientFactory.Create(client.WithProxyURL(agg.Proxy().URL()))
	httpBin := c.httpBinFactory.Create(cl)

	output, err := httpBin.Get(ctx)
	if err != nil {
		return task.Result{}, errors.Wrap(err, "httpBin.Get")
	}

	return task.Result{
		MITMResult: &task.MITMResult{
			HasMITM: output.Args.Nonce != httpbin.Nonce,
		},
	}, nil
}
