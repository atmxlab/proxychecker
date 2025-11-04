package geo

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/details/service/ipapi"
	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type IPApi interface {
	Get(ctx context.Context) (ipapi.Output, error)
}

type IPApiFactory interface {
	Create(client client.Client) IPApi
}

type Checker struct {
	clientFactory client.Factory
	ipApiFactory  IPApiFactory
}

func New(clientFactory client.Factory, ipApiFactory IPApiFactory) *Checker {
	return &Checker{clientFactory: clientFactory, ipApiFactory: ipApiFactory}
}

func (c *Checker) Run(ctx context.Context, p *entity.Proxy) (task.Result, error) {
	cl := c.clientFactory.Create(p)
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
