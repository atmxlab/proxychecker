package queue

import "time"

type Kind int16

type ID string

type Task struct {
	id         ID
	kind       Kind
	externalID string
	payload    []byte
	createdAt  time.Time
	updatedAt  time.Time
}
