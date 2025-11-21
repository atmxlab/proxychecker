package task

type Kind int16

const (
	KindUnknown Kind = iota
	KindCheckLatency
	KindCheckGEO
)

func (t Kind) String() string {
	m := map[Kind]string{
		KindUnknown:      "Unknown",
		KindCheckLatency: "CheckLatency",
		KindCheckGEO:     "CheckGEO",
	}

	return m[t]
}
