package proxychecker

import (
	"context"

	desc "github.com/atmxlab/proxychecker/gen/proto/api/proxychecker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/service/command"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/samber/lo"
)

func (s *Service) Check(ctx context.Context, req *desc.CheckRequest) (*desc.CheckResponse, error) {
	res, err := s.c.Commands().Check().Execute(ctx, command.CheckInput{
		OperationTime: s.c.Entities().TimeProvider().CurrentTime(ctx),
		Proxies:       req.GetProxies(),
		Checkers: lo.Map(req.GetKinds(), func(item desc.CheckKind, _ int) checker.Kind {
			m := map[desc.CheckKind]checker.Kind{
				desc.CheckKind_CHECK_KIND_UNKNOWN: checker.KindUnknown,
				desc.CheckKind_CHECK_KIND_GEO:     checker.KindGEO,
				desc.CheckKind_CHECK_KIND_LATENCY: checker.KindLatency,
			}

			return m[item]
		}),
	})
	if err != nil {
		return nil, errors.Wrap(err, "check.Execute")
	}
	
	return &desc.CheckResponse{
		GroupId: res.TaskGroupID.String(),
	}, nil
}
