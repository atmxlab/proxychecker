package entity

import (
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
)

type Task struct {
	id           task.ID
	groupID      task.GroupID
	proxyID      proxy.ID
	checkerID    checker.ID
	status       task.Status
	state        task.State
	lockDeadline *time.Time
	createdAt    time.Time
	updatedAt    time.Time
}

func (t Task) GroupID() task.GroupID {
	return t.groupID
}

func (t Task) ID() task.ID {
	return t.id
}

func (t Task) ProxyID() proxy.ID {
	return t.proxyID
}

func (t Task) CheckerID() checker.ID {
	return t.checkerID
}

func (t Task) Status() task.Status {
	return t.status
}

func (t Task) State() task.State {
	return t.state
}

func (t Task) LockDeadline() *time.Time {
	return t.lockDeadline
}

func (t Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t Task) UpdatedAt() time.Time {
	return t.updatedAt
}
