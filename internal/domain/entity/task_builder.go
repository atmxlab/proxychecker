package entity

import (
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
)

type TaskBuilder struct {
	id          task.ID
	groupID     task.GroupID
	proxyID     proxy.ID
	checkerKind checker.Kind
	payload     task.Payload
	createdAt   time.Time
	updatedAt   time.Time
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

func (b *TaskBuilder) CheckerKind(kind checker.Kind) *TaskBuilder {
	b.checkerKind = kind
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

func (b *TaskBuilder) TargetURL(targetURL task.TargetURL) *TaskBuilder {
	b.payload.TargetURL = &targetURL
	return b
}

func (b *TaskBuilder) Build() *Task {
	return &Task{
		id:          b.id,
		groupID:     b.groupID,
		proxyID:     b.proxyID,
		checkerKind: b.checkerKind,
		status:      task.StatusPending,
		state:       task.State{},
		createdAt:   b.createdAt,
		updatedAt:   b.updatedAt,
	}
}
