package proxychecker

import (
	"github.com/atmxlab/proxychecker/cmd/app"
	desc "github.com/atmxlab/proxychecker/gen/proto/api/proxychecker"
)

type Service struct {
	desc.UnimplementedProxycheckerServer
	c *app.Container
}

func New(c *app.Container) *Service {
	return &Service{
		c: c,
	}
}
