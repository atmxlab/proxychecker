package checker

import (
	"context"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/details/service/httpbin"
	"github.com/atmxlab/proxychecker/internal/details/service/ipapi"
)

//go:generate mock Client
type Client = client.Client

//go:generate mock ClientFactory
type ClientFactory = client.Factory

//go:generate mock IPApi
type IPApi interface {
	Get(ctx context.Context) (ipapi.Output, error)
}

//go:generate mock IPApiFactory
type IPApiFactory interface {
	Create(client Client) IPApi
}

//go:generate mock IPApi
type HTTPBin interface {
	Get(ctx context.Context) (httpbin.Output, error)
}

//go:generate mock IPApiFactory
type HTTPBinFactory interface {
	Create(client Client) IPApi
}
