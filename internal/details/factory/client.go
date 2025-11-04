package factory

import (
	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/details/client/http"
	"github.com/atmxlab/proxychecker/internal/domain/entity"
)

type ClientFactory struct {
}

func NewClientFactory() *ClientFactory {
	return &ClientFactory{}
}

func (c ClientFactory) Create(p *entity.Proxy) client.Client {
	// return http.NewClient(p.URL())
	return http.NewClient("") // TODO: remove
}
