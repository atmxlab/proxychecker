package factory

import (
	"github.com/atmxlab/proxychecker/internal/details/checker"
	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/internal/details/service/httpbin"
)

type HTTPBinFactory struct {
}

func NewHTTPBinFactory() *HTTPBinFactory {
	return &HTTPBinFactory{}
}

func (I HTTPBinFactory) Create(cl client.Client) checker.HTTPBin {
	return httpbin.New(cl)
}
