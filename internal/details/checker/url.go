package checker

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type URLChecker struct {
	clientFactory ClientFactory
}

func NewURLChecker(clientFactory ClientFactory) *URLChecker {
	return &URLChecker{clientFactory: clientFactory}
}

func (c *URLChecker) Run(ctx context.Context, agg *aggregate.Task) (task.Result, error) {
	cl := c.clientFactory.Create(client.WithProxyURL(agg.Proxy().URL()))

	targetURL, ok := agg.Task().TargetURL()
	if !ok {
		return task.Result{}, errors.New("empty target URL")
	}

	_, err := cl.Get(ctx, targetURL.URL)
	if err != nil {
		return task.Result{
			URLResult: &task.URLResult{
				IsAvailable: false,
			},
		}, nil
	}

	return task.Result{
		URLResult: &task.URLResult{
			IsAvailable: true,
		},
	}, nil
}
