package factory

import (
	"github.com/atmxlab/proxychecker/internal/details/checker/geo"
	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/details/service/ipapi"
)

type IPApiFactory struct {
}

func NewIPApiFactory() *IPApiFactory {
	return &IPApiFactory{}
}

func (I IPApiFactory) Create(cl client.Client) geo.IPApi {
	return ipapi.New(cl)
}
