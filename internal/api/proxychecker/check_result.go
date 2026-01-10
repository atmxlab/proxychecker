package proxychecker

import (
	"context"

	desc "github.com/atmxlab/proxychecker/gen/proto/api/proxychecker"
	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/pkg/conv"
	"github.com/atmxlab/proxychecker/internal/service/query"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/samber/lo"
)

func (s *Service) CheckResult(ctx context.Context, req *desc.CheckResultRequest) (*desc.CheckResultResponse, error) {
	res, err := s.c.Queries().CheckResult().Execute(ctx, query.CheckResultInput{
		GroupID: task.GroupID(req.GetGroupId()),
	})
	if err != nil {
		return nil, errors.Wrap(err, "check result")
	}

	return &desc.CheckResultResponse{
		Statistic: &desc.CheckResultResponse_Statistic{
			Proxies: &desc.CheckResultResponse_ProxiesStatistic{
				Count:        int64(res.ProxiesStatistic.Count),
				CheckedCount: int64(res.ProxiesStatistic.CheckedCount),
				PendingCount: int64(res.ProxiesStatistic.PendingCount),
			},
			Tasks: &desc.CheckResultResponse_TasksStatistic{
				Count:        int64(res.TasksStatistic.Count),
				SuccessCount: int64(res.TasksStatistic.SuccessCount),
				FailureCount: int64(res.TasksStatistic.FailureCount),
				PendingCount: int64(res.TasksStatistic.PendingCount),
			},
		},
		Proxies: lo.Map(res.Proxies, func(
			item query.CheckResultOutputProxyInfo,
			_ int,
		) *desc.CheckResultResponse_Proxy {
			return s.mapProxiesResult(item)
		}),
		IsChecked: res.IsChecked,
	}, nil
}

func (s *Service) mapProxiesResult(pxInfo query.CheckResultOutputProxyInfo) *desc.CheckResultResponse_Proxy {
	return &desc.CheckResultResponse_Proxy{
		Id:        pxInfo.Proxy.ID().String(),
		Url:       pxInfo.Proxy.URL(),
		IsChecked: pxInfo.IsChecked,
		TasksStatistic: &desc.CheckResultResponse_TasksStatistic{
			Count:        int64(pxInfo.TasksStatistic.Count),
			SuccessCount: int64(pxInfo.TasksStatistic.SuccessCount),
			FailureCount: int64(pxInfo.TasksStatistic.FailureCount),
			PendingCount: int64(pxInfo.TasksStatistic.PendingCount),
		},
		Tasks: lo.Map(pxInfo.Tasks, func(item *entity.Task, _ int) *desc.Task {
			return s.mapTask(item)
		}),
	}
}

func (s *Service) mapTask(tk *entity.Task) *desc.Task {
	pbtk := &desc.Task{
		CheckerKind: conv.FromCheckerKind(tk.CheckerKind()),
		Status:      conv.FromTaskStatus(tk.Status()),
	}

	if e := tk.State().Result().ErrorResult; e != nil && s.IsProd() {
		pbtk.Result = &desc.Task_Error{
			Error: &desc.Task_ResultError{
				Message: e.Code.String(),
			},
		}
	}

	if e := tk.State().Result().ErrorResult; e != nil && !s.IsProd() {
		pbtk.Result = &desc.Task_Error{
			Error: &desc.Task_ResultError{
				Message: e.String(),
			},
		}
	}

	if geo := tk.State().Result().GEOResult; geo != nil {
		pbtk.Result = &desc.Task_Geo{
			Geo: &desc.Task_ResultGEO{
				CountryCode: geo.CountryCode,
				Region:      geo.Region,
				City:        geo.City,
				Timezone:    geo.Timezone,
			},
		}
	}

	if latency := tk.State().Result().LatencyResult; latency != nil {
		pbtk.Result = &desc.Task_Latency{
			Latency: &desc.Task_ResultLatency{
				FromHostToProxyRoundTrip:   latency.FromHostToProxyRoundTrip.Microseconds(),
				FromHostToTargetRoundTrip:  latency.FromHostToTargetRoundTrip.Microseconds(),
				FromProxyToTargetRoundTrip: latency.FromProxyToTargetRoundTrip.Microseconds(),
			},
		}
	}

	if res := tk.State().Result().ExternalIPResult; res != nil {
		pbtk.Result = &desc.Task_ExternalIp{
			ExternalIp: &desc.Task_ResultExternalIP{
				Ip: res.IP,
			},
		}
	}

	if res := tk.State().Result().URLResult; res != nil {
		pbtk.Result = &desc.Task_Url{
			Url: &desc.Task_ResultURL{
				Url:         res.URL,
				IsAvailable: res.IsAvailable,
				StatusCode:  int64(res.StatusCode),
			},
		}
	}

	if res := tk.State().Result().HTTPSResult; res != nil {
		pbtk.Result = &desc.Task_Https{
			Https: &desc.Task_ResultHTTPS{
				IsAvailable: res.IsAvailable,
			},
		}
	}

	if res := tk.State().Result().MITMResult; res != nil {
		pbtk.Result = &desc.Task_Mitm{
			Mitm: &desc.Task_ResultMITM{
				HasMitm: res.HasMITM,
			},
		}
	}

	if res := tk.State().Result().TypeResult; res != nil {
		pbtk.Result = &desc.Task_Type{
			Type: &desc.Task_ResultType{
				Type: mapProxyType(res.Type),
			},
		}
	}

	if res := tk.State().Result().AnonymousResult; res != nil {
		pbtk.Result = &desc.Task_Anonymous{
			Anonymous: &desc.Task_ResultAnonymous{
				Kind: mapAnonymousKind(res.Kind),
				SuspiciousHeaders: lo.Map(res.SuspiciousHeaders, func(
					item task.Header,
					_ int,
				) *desc.Task_ResultAnonymous_Header {
					return &desc.Task_ResultAnonymous_Header{
						Key:   item.Key,
						Value: item.Value,
					}
				}),
			},
		}
	}

	return pbtk
}

func mapProxyType(t proxy.Type) desc.Task_ResultType_Type {
	m := map[proxy.Type]desc.Task_ResultType_Type{
		proxy.TypeUnknown:     desc.Task_ResultType_UNKNOWN,
		proxy.TypeDatacenter:  desc.Task_ResultType_DATACENTER,
		proxy.TypeResidential: desc.Task_ResultType_RESIDENTIAL,
		proxy.TypeMobile:      desc.Task_ResultType_MOBILE,
	}

	return m[t]
}

func mapAnonymousKind(kind proxy.AnonymouseKind) desc.Task_ResultAnonymous_Kind {
	m := map[proxy.AnonymouseKind]desc.Task_ResultAnonymous_Kind{
		proxy.AnonymousKindUnknown:     desc.Task_ResultAnonymous_UNKNOWN,
		proxy.AnonymousKindTransparent: desc.Task_ResultAnonymous_TRANSPARENT,
		proxy.AnonymousKindMiddle:      desc.Task_ResultAnonymous_MIDDLE,
		proxy.AnonymousKindHigh:        desc.Task_ResultAnonymous_HIGH,
	}

	return m[kind]
}
