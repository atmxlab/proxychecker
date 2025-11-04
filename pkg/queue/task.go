package queue

import "time"

type Status int

const (
	StatusUnknown Status = iota
	StatusPending
	StatusRunning
	StatusSuccess
	StatusFailure
)

type Kind int16

type ID string

type Task struct {
	id         ID
	kind       Kind
	status     Status
	externalID string
	payload    []byte
	createdAt  time.Time
	updatedAt  time.Time
}
