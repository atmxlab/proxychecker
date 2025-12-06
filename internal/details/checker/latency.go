package checker

import (
	"context"
	"fmt"
	"net/url"
	stdtime "time"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/time"
)

//go:generate mock TimeProvider
type TimeProvider time.Provider

type LatencyChecker struct {
	clientFactory ClientFactory
	ipApiFactory  IPApiFactory
	timeProvider  TimeProvider
}

func NewLatencyChecker(
	clientFactory ClientFactory,
	ipApiFactory IPApiFactory,
	timeProvider TimeProvider,
) *LatencyChecker {
	return &LatencyChecker{clientFactory: clientFactory, ipApiFactory: ipApiFactory, timeProvider: timeProvider}
}

func (c *LatencyChecker) Run(ctx context.Context, tAgg *aggregate.Task) (task.Result, error) {
	fromHostToProxyRoundTrip := 0 * stdtime.Second
	fromHostToProxyRoundTripStart := c.timeProvider.CurrentTime(ctx)
	clientWithoutProxy := c.clientFactory.Create(
		client.WithDisableKeepAlives(),
		client.WithAfterDialHook(func() {
			fromHostToProxyRoundTrip = c.timeProvider.Since(ctx, fromHostToProxyRoundTripStart)
		}),
	)

	pxURL, err := url.Parse(tAgg.Proxy().URL())
	if err != nil {
		return task.Result{}, errors.Wrapf(err, "url.Parse: url: [%s]", tAgg.Proxy().URL())
	}
	_, err = clientWithoutProxy.Get(ctx, fmt.Sprintf("http://%s", pxURL.Host))
	if err != nil {
		return task.Result{}, errors.Wrap(err, "clientWithoutProxy.Get")
	}

	clientWithProxy := c.clientFactory.Create(
		client.WithDisableKeepAlives(),
		client.WithProxyURL(tAgg.Proxy().URL()),
	)
	ipAPI := c.ipApiFactory.Create(clientWithProxy)

	fromHostToTargetRoundTrip, err := c.benchmark(ctx, func() error {
		_, err = ipAPI.Get(ctx)
		if err != nil {
			return errors.Wrap(err, "ipAPI.Get")
		}

		return nil
	})
	if err != nil {
		return task.Result{}, errors.Wrap(err, "c.benchmark")
	}

	return task.Result{
		LatencyResult: &task.LatencyResult{
			FromHostToProxyRoundTrip:   fromHostToProxyRoundTrip,
			FromHostToTargetRoundTrip:  fromHostToTargetRoundTrip,
			FromProxyToTargetRoundTrip: fromHostToTargetRoundTrip - fromHostToProxyRoundTrip,
		},
	}, nil
}

func (c *LatencyChecker) benchmark(ctx context.Context, cb func() error) (stdtime.Duration, error) {
	now := c.timeProvider.CurrentTime(ctx)
	if err := cb(); err != nil {
		return 0, errors.Wrap(err, "benchmark callback")
	}

	return c.timeProvider.Since(ctx, now), nil
}
