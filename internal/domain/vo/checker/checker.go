package checker

type Kind string

const (
	KindUnknown    Kind = "Unknown"
	KindLatency    Kind = "Latency"
	KindGEO        Kind = "GEO"
	KindExternalIP Kind = "ExternalIP"
)
