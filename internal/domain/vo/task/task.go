package task

type ID string

func NewID() ID {
	return ID("") // TODO: uuid
}

type GroupID string

func NewGroupID() GroupID {
	return GroupID("") // TODO: uuid
}

type Status string

const (
	StatusUnknown Status = "unknown"
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusSuccess Status = "success"
	StatusFailure Status = "failure"
)
