package checker

import (
	"context"
	"io"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/sirupsen/logrus"
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

	res, err := cl.Do(ctx, targetURL.URL)
	if err != nil {
		return task.Result{}, errors.Wrap(err, "client.Get")
	}
	defer res.Body.Close()

	if c.is2xxStatus(res.StatusCode) {
		logrus.Warn(io.ReadAll(res.Body))
		return task.Result{
			URLResult: &task.URLResult{
				IsAvailable: true,
				URL:         targetURL.URL,
				StatusCode:  res.StatusCode,
			},
		}, nil
	}

	return task.Result{
		URLResult: &task.URLResult{
			IsAvailable: false,
			URL:         targetURL.URL,
			StatusCode:  res.StatusCode,
		},
	}, nil
}

func (c *URLChecker) is2xxStatus(status int) bool {
	return status >= 200 && status < 300
}
