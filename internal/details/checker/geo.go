package checker

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type GeoChecker struct {
	clientFactory ClientFactory
	ipApiFactory  IPApiFactory
}

func NewGeoChecker(clientFactory ClientFactory, ipApiFactory IPApiFactory) *GeoChecker {
	return &GeoChecker{clientFactory: clientFactory, ipApiFactory: ipApiFactory}
}

func (c *GeoChecker) Run(ctx context.Context, agg *aggregate.Task) (task.Result, error) {
	cl := c.clientFactory.Create(client.WithProxyURL(agg.Proxy().URL()))
	ipApi := c.ipApiFactory.Create(cl)

	output, err := ipApi.Get(ctx)
	if err != nil {
		return task.Result{}, errors.Wrap(err, "ipApi.Get")
	}

	return task.Result{
		GEOResult: &task.GEOResult{
			ContinentCode: output.ContinentCode,
			Continent:     output.Continent,
			CountryCode:   output.CountryCode,
			Country:       output.Country,
			Region:        output.Country,
			City:          output.City,
			Timezone:      output.Timezone,
		},
	}, nil
}
