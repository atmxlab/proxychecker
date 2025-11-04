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

func (t Task) ID() ID {
	return t.id
}

func (t Task) Kind() Kind {
	return t.kind
}

func (t Task) Status() Status {
	return t.status
}

func (t Task) ExternalID() string {
	return t.externalID
}

func (t Task) Payload() []byte {
	return t.payload
}

func (t Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t Task) UpdatedAt() time.Time {
	return t.updatedAt
}
