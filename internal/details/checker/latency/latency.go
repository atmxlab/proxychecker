package latency

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/details/checker/geo"
	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/details/service/ipapi"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/time"
)

type IPApi interface {
	Get(ctx context.Context) (ipapi.Output, error)
}

type IPApiFactory interface {
	Create(client client.Client) geo.IPApi
}

//go:generate mock TimeProvider
type TimeProvider time.Provider

type Checker struct {
	clientFactory client.Factory
	ipApiFactory  IPApiFactory
	timeProvider  TimeProvider
}

func New(clientFactory client.Factory, ipApiFactory IPApiFactory, timeProvider TimeProvider) *Checker {
	return &Checker{clientFactory: clientFactory, ipApiFactory: ipApiFactory, timeProvider: timeProvider}
}

func (c *Checker) Run(ctx context.Context, tAgg *aggregate.Task) (task.Result, error) {
	cl := c.clientFactory.Create(tAgg.Proxy())
	ipAPI := c.ipApiFactory.Create(cl)

	now := c.timeProvider.CurrentTime(ctx)

	_, err := ipAPI.Get(ctx)
	if err != nil {
		return task.Result{}, errors.Wrap(err, "ipAPI.Get")
	}

	since := c.timeProvider.Since(ctx, now)
	return task.Result{
		LatencyResult: &task.LatencyResult{
			LatencyToProxy:  since / 2, // TODO: ping proxy
			LatencyToTarget: since,
		},
	}, nil
}
