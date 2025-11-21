package task

import "github.com/atmxlab/proxychecker/pkg/uuid"

type ID string

func NewID() ID {
	return ID(uuid.MustV7().String())
}

func (id ID) String() string {
	return string(id)
}

type GroupID string

func NewGroupID() GroupID {
	return GroupID(uuid.MustV7().String())
}

func (id GroupID) String() string {
	return string(id)
}

type Status string

const (
	StatusUnknown Status = "unknown"
	StatusPending Status = "pending"
	StatusSuccess Status = "success"
	StatusFailure Status = "failure"
)
