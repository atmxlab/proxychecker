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

	bytes, err := cl.Get(ctx, "https://google.com")
	if err != nil {
		for _, msg := range c.connectToProxyErrMessages() {
			if strings.Contains(err.Error(), msg) {
				return task.Result{
					HTTPSResult: &task.HTTPSResult{
						IsAvailable: false,
					},
				}, nil
			}
		}

		return task.Result{}, errors.Wrap(err, "client.Get")
	}

	if !strings.Contains(string(bytes), "itemscope") {
		return task.Result{
			HTTPSResult: &task.HTTPSResult{
				IsAvailable: false,
			},
		}, nil
	}

	return task.Result{
		HTTPSResult: &task.HTTPSResult{
			IsAvailable: true,
		},
	}, nil
}

func (c *HTTPSChecker) connectToProxyErrMessages() []string {
	return []string{
		"Method not allowed",
		"Bad Request",
		"Bad Gateway",
		"Service Unavailable",
		"Internal Server Error",
	}
}
