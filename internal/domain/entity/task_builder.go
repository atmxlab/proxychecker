package entity

import (
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
)

type TaskBuilder struct {
	id        task.ID
	groupID   task.GroupID
	proxyID   proxy.ID
	checkerID checker.ID
	createdAt time.Time
	updatedAt time.Time
}

func NewTaskBuilder() *TaskBuilder {
	return &TaskBuilder{}
}

func (b *TaskBuilder) ID(id task.ID) *TaskBuilder {
	b.id = id
	return b
}

func (b *TaskBuilder) GroupID(id task.GroupID) *TaskBuilder {
	b.groupID = id
	return b
}

func (b *TaskBuilder) ProxyID(id proxy.ID) *TaskBuilder {
	b.proxyID = id
	return b
}

func (b *TaskBuilder) CheckerID(id checker.ID) *TaskBuilder {
	b.checkerID = id
	return b
}

func (b *TaskBuilder) CreatedAt(t time.Time) *TaskBuilder {
	b.createdAt = t
	return b
}

func (b *TaskBuilder) UpdatedAt(t time.Time) *TaskBuilder {
	b.updatedAt = t
	return b
}

func (b *TaskBuilder) Build() *Task {
	return &Task{
		id:           b.id,
		groupID:      b.groupID,
		proxyID:      b.proxyID,
		checkerID:    b.checkerID,
		status:       task.StatusPending,
		state:        task.State{},
		lockDeadline: nil,
		createdAt:    b.createdAt,
		updatedAt:    b.updatedAt,
	}
}
