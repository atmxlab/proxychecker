package conv

import (
	desc "github.com/atmxlab/proxychecker/gen/proto/api/proxychecker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
)

func FromCheckerKind(kind checker.Kind) desc.CheckKind {
	m := map[checker.Kind]desc.CheckKind{
		checker.KindUnknown:    desc.CheckKind_CHECK_KIND_UNKNOWN,
		checker.KindGEO:        desc.CheckKind_CHECK_KIND_GEO,
		checker.KindLatency:    desc.CheckKind_CHECK_KIND_LATENCY,
		checker.KindExternalIP: desc.CheckKind_CHECK_KIND_EXTERNAL_IP,
		checker.KindURL:        desc.CheckKind_CHECK_KIND_URL,
	}

	return m[kind]
}

func FromTaskStatus(status task.Status) desc.Task_Status {
	m := map[task.Status]desc.Task_Status{
		task.StatusUnknown: desc.Task_STATUS_UNKNOWN,
		task.StatusPending: desc.Task_STATUS_PENDING,
		task.StatusSuccess: desc.Task_STATUS_SUCCESS,
		task.StatusFailure: desc.Task_STATUS_FAILURE,
	}

	return m[status]
}
