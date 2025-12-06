package checker

type Kind string

const (
	KindUnknown    Kind = "unknown"
	KindLatency    Kind = "latency"
	KindGEO        Kind = "geo"
	KindExternalIP Kind = "externalIP"
)

func KindFromString(kind string) Kind {
	m := map[string]Kind{
		"unknown":    KindUnknown,
		"latency":    KindLatency,
		"geo":        KindGEO,
		"externalIP": KindExternalIP,
	}

	if k, ok := m[kind]; ok {
		return k
	}

	return KindUnknown
}
