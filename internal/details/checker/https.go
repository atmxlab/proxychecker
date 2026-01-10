package checker

import (
	"context"
	"strings"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type HTTPSChecker struct {
	clientFactory  ClientFactory
	httpBinFactory HTTPBinFactory
}

func NewHTTPSChecker(clientFactory ClientFactory, httpBinFactory HTTPBinFactory) *HTTPSChecker {
	return &HTTPSChecker{clientFactory: clientFactory, httpBinFactory: httpBinFactory}
}

func (c *HTTPSChecker) Run(ctx context.Context, agg *aggregate.Task) (task.Result, error) {
	cl := c.clientFactory.Create(client.WithProxyURL(agg.Proxy().URL()))
	httpBin := c.httpBinFactory.Create(cl)

	_, err := httpBin.Get(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Method not allowed") {
			return task.Result{
				HTTPSResult: &task.HTTPSResult{
					IsAvailable: false,
				},
			}, nil
		}
		
		return task.Result{}, errors.Wrap(err, "httpBin.Get")
	}

	return task.Result{
		HTTPSResult: &task.HTTPSResult{
			IsAvailable: true,
		},
	}, nil
}
