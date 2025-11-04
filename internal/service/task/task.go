package task

type Kind int16

const (
	KindUnknown Kind = iota
	KindCheckLatency
	KindCheckGEO
)

func (t Kind) String() string {
	switch t {
	case KindCheckLatency:
		return "TaskCheckLatency"
	case KindCheckGEO:
		return "TaskCheckGEO"
	default:
		return "TaskUnknown"
	}
}
