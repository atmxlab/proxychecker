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
	lockDeadline time.Time
	createdAt    time.Time
	updatedAt    time.Time
}
