package checker

import (
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/validator"
)

type Kind string

const (
	KindUnknown    Kind = "unknown"
	KindLatency    Kind = "latency"
	KindGEO        Kind = "geo"
	KindExternalIP Kind = "externalIP"
	KindURL        Kind = "url"
	KindHTTPS      Kind = "https"
	KindMITM       Kind = "mitm"
	KindType       Kind = "type"
)

func KindFromString(kind string) Kind {
	m := map[string]Kind{
		"unknown":    KindUnknown,
		"latency":    KindLatency,
		"geo":        KindGEO,
		"externalIP": KindExternalIP,
		"url":        KindURL,
		"https":      KindHTTPS,
		"mitm":       KindMITM,
		"type":       KindType,
	}

	if k, ok := m[kind]; ok {
		return k
	}

	return KindUnknown
}

type KindWithPayload struct {
	payload task.Payload
	kind    Kind
}

func NewKindWithPayload(payload task.Payload, kind Kind) KindWithPayload {
	return KindWithPayload{payload: payload, kind: kind}
}

func (c KindWithPayload) Payload() task.Payload {
	return c.payload
}

func (c KindWithPayload) Kind() Kind {
	return c.kind
}

func (c KindWithPayload) Validate() error {
	v := validator.New()

	if c.kind == KindUnknown {
		v.Failed("kind is unknown")
	}
	if c.kind == KindURL {
		if c.payload.TargetURL == nil {
			v.Failedf("checker [%s] must have a target URL", c.kind)
		} else {
			v.AddErr(c.payload.TargetURL.Validate())
		}
	}

	return v.Err()
}
