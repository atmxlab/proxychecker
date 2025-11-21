package queue

import "time"

type TaskBuilder struct {
	id         ID
	kind       Kind
	status     Status
	externalID string
	payload    []byte
	createdAt  time.Time
	updatedAt  time.Time
}

func NewTaskBuilder() *TaskBuilder {
	return &TaskBuilder{}
}

func (t *TaskBuilder) ID(id ID) *TaskBuilder {
	t.id = id
	return t
}

func (t *TaskBuilder) Kind(kind Kind) *TaskBuilder {
	t.kind = kind
	return t
}

func (t *TaskBuilder) Status(status Status) *TaskBuilder {
	t.status = status
	return t
}

func (t *TaskBuilder) ExternalID(externalID string) *TaskBuilder {
	t.externalID = externalID
	return t
}

func (t *TaskBuilder) Payload(payload []byte) *TaskBuilder {
	t.payload = payload
	return t
}

func (t *TaskBuilder) CreatedAt(createdAt time.Time) *TaskBuilder {
	t.createdAt = createdAt
	return t
}

func (t *TaskBuilder) UpdatedAt(updatedAt time.Time) *TaskBuilder {
	t.updatedAt = updatedAt
	return t
}

func (t *TaskBuilder) Build() Task {
	return Task{
		id:         t.id,
		kind:       t.kind,
		status:     t.status,
		externalID: t.externalID,
		payload:    t.payload,
		createdAt:  t.createdAt,
		updatedAt:  t.updatedAt,
	}
}
