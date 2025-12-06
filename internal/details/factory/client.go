package factory

import (
	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/details/client/http"
)

type ClientFactory struct {
}

func NewClientFactory() *ClientFactory {
	return &ClientFactory{}
}

func (c ClientFactory) Create(opts ...client.Option) client.Client {
	cfg := client.Config{}
	for _, opt := range opts {
		opt(&cfg)
	}
	return http.NewClient(cfg)
}
