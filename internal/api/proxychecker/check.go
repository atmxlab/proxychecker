package proxychecker

import (
	"context"

	desc "github.com/atmxlab/proxychecker/gen/proto/api/proxychecker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/service/command"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/samber/lo"
)

func (s *Service) Check(ctx context.Context, req *desc.CheckRequest) (*desc.CheckResponse, error) {
	res, err := s.c.Commands().Check().Execute(ctx, command.CheckInput{
		OperationTime: s.c.Entities().TimeProvider().CurrentTime(ctx),
		Proxies:       req.GetProxies(),
		Checkers: lo.Map(req.GetKinds(), func(item *desc.CheckRequest_Kind, _ int) checker.KindWithPayload {
			getKind := func() checker.Kind {
				switch item.GetKind().(type) {
				case *desc.CheckRequest_Kind_Geo_:
					return checker.KindGEO
				case *desc.CheckRequest_Kind_Latency_:
					return checker.KindLatency
				case *desc.CheckRequest_Kind_Url_:
					return checker.KindURL
				case *desc.CheckRequest_Kind_ExternalIp_:
					return checker.KindExternalIP
				default:
					return checker.KindUnknown
				}
			}

			getPayload := func() task.Payload {
				switch v := item.GetKind().(type) {
				case *desc.CheckRequest_Kind_Url_:
					return task.Payload{
						TargetURL: &task.TargetURL{URL: v.Url.GetUrl()},
					}
				default:
					return task.NewEmptyPayload()
				}
			}

			return checker.NewKindWithPayload(getPayload(), getKind())
		}),
	})
	if err != nil {
		return nil, errors.Wrap(err, "check.Execute")
	}

	return &desc.CheckResponse{
		GroupId: res.TaskGroupID.String(),
	}, nil
}
