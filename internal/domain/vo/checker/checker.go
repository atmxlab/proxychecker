package checker

type ID string

type Kind string

const (
	KindUnknown    Kind = "unknown"
	KindLatency    Kind = "latency"
	KindGEO        Kind = "geo"
	KindExternalIP Kind = "external_ip"
)
