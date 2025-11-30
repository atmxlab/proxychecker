package task

import (
	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/pkg/queue"
)

type Task interface {
	Kind() Kind
	Key() string
	Unmarshal([]byte) error
	Marshal() ([]byte, error)
	Validate() error
}
type Kind string

const (
	KindUnknown         Kind = "Unknown"
	KindCheckLatency    Kind = "CheckLatency"
	KindCheckGEO        Kind = "CheckGEO"
	KindCheckExternalIP Kind = "CheckExternalIP"
)

func (t Kind) String() string {
	return string(t)
}

func (t Kind) ToQueue() queue.Kind {
	return queue.Kind(t)
}

func FromDomainTask(kind checker.Kind) Kind {
	m := map[checker.Kind]Kind{
		checker.KindUnknown:    KindUnknown,
		checker.KindLatency:    KindCheckLatency,
		checker.KindGEO:        KindCheckGEO,
		checker.KindExternalIP: KindCheckExternalIP,
	}

	return m[kind]
}
