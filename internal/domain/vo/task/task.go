package task

type ID string

type GroupID string

type Status string

const (
	StatusUnknown Status = "unknown"
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusSuccess Status = "success"
	StatusFailure Status = "failure"
)
