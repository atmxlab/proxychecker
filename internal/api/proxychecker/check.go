package proxychecker

import (
	"context"

	desc "github.com/atmxlab/proxychecker/gen/proto/api/proxychecker"
)

func (s *Service) Check(ctx context.Context, req *desc.CheckRequest) (*desc.CheckResponse, error) {
	return &desc.CheckResponse{}, nil
}
