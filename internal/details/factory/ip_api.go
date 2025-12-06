package factory

import (
	"github.com/atmxlab/proxychecker/internal/details/checker"
	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/details/service/ipapi"
)

type IPApiFactory struct {
}

func NewIPApiFactory() *IPApiFactory {
	return &IPApiFactory{}
}

func (I IPApiFactory) Create(cl client.Client) checker.IPApi {
	return ipapi.New(cl)
}
