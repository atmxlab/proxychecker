package checker

import (
	"context"
	"strings"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/samber/lo"
)

type AnonymousChecker struct {
	serverIP       string
	clientFactory  ClientFactory
	httpBinFactory HTTPBinFactory
}

func NewAnonymousChecker(
	serverIP string,
	clientFactory ClientFactory,
	httpBinFactory HTTPBinFactory,
) *AnonymousChecker {
	return &AnonymousChecker{serverIP: serverIP, clientFactory: clientFactory, httpBinFactory: httpBinFactory}
}

func (c *AnonymousChecker) Run(ctx context.Context, agg *aggregate.Task) (task.Result, error) {
	cl := c.clientFactory.Create(client.WithProxyURL(agg.Proxy().URL()))
	httpBin := c.httpBinFactory.Create(cl)

	output, err := httpBin.Get(ctx)
	if err != nil {
		return task.Result{}, errors.Wrap(err, "httpBin.Get")
	}

	if output.Origin == "" {
		return task.Result{}, errors.New("httpbin.Get returned empty result")
	}

	if output.Origin == c.serverIP {
		return task.Result{
			AnonymousResult: &task.AnonymousResult{
				Kind:              proxy.AnonymousKindTransparent,
				SuspiciousHeaders: nil,
			},
		}, nil
	}

	normalizedSuspiciousHeaderNames := lo.Map(task.SuspiciousHeaderNames(), func(item string, _ int) string {
		return strings.ToLower(item)
	})

	suspiciousHeaders := make([]task.Header, 0)

	for key, value := range output.Headers {
		if lo.Contains(normalizedSuspiciousHeaderNames, strings.ToLower(key)) ||
			strings.Contains(value, c.serverIP) ||
			strings.Contains(key, c.serverIP) {
			suspiciousHeaders = append(suspiciousHeaders, task.Header{Key: key, Value: value})
		}
	}

	return task.Result{
		AnonymousResult: &task.AnonymousResult{
			Kind:              lo.Ternary(len(suspiciousHeaders) == 0, proxy.AnonymousKindHigh, proxy.AnonymousKindMiddle),
			SuspiciousHeaders: suspiciousHeaders,
		},
	}, nil
}
