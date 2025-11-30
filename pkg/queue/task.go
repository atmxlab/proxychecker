package queue

import (
	"time"

	"github.com/atmxlab/proxychecker/pkg/uuid"
)

type Status int

const (
	StatusUnknown Status = iota
	StatusPending
	StatusRunning
	StatusSuccess
	StatusFailure
)

func (s Status) String() string {
	m := map[Status]string{
		StatusUnknown: "unknown",
		StatusPending: "pending",
		StatusRunning: "running",
		StatusSuccess: "success",
		StatusFailure: "failure",
	}

	return m[s]
}

type Kind string

type ID string

func NewID() ID {
	return ID(uuid.MustV7().String())
}

func (id ID) String() string {
	return string(id)
}

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

func (t Task) SetStatus(status Status) Task {
	t.status = status // TODO: refactor: use modifier
	return t
}
