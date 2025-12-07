package proxychecker

import (
	"github.com/atmxlab/proxychecker/cmd/app"
	desc "github.com/atmxlab/proxychecker/gen/proto/api/proxychecker"
)

type Service struct {
	desc.UnimplementedProxycheckerServer
	c app.Commands
}

func New(c app.Commands) *Service {
	return &Service{
		c: c,
	}
}
