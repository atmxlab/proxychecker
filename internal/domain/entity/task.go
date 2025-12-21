package entity

import (
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type Task struct {
	id          task.ID
	groupID     task.GroupID
	proxyID     proxy.ID
	checkerKind checker.Kind
	status      task.Status
	payload     task.Payload
	state       task.State
	createdAt   time.Time
	updatedAt   time.Time
}

func (t *Task) GroupID() task.GroupID {
	return t.groupID
}

func (t *Task) ID() task.ID {
	return t.id
}

func (t *Task) ProxyID() proxy.ID {
	return t.proxyID
}

func (t *Task) CheckerKind() checker.Kind {
	return t.checkerKind
}

func (t *Task) Status() task.Status {
	return t.status
}

func (t *Task) State() task.State {
	return t.state
}

func (t *Task) TargetURL() (task.TargetURL, bool) {
	if t.payload.TargetURL != nil {
		return *t.payload.TargetURL, true
	}

	return task.TargetURL{}, false
}

func (t *Task) Payload() task.Payload {
	return t.payload
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Task) Modify(cb func(m *TaskModifier) error) error {
	if err := cb(NewTaskModifier(t)); err != nil {
		return errors.Wrap(err, "callback task modifier")
	}

	return nil
}

type TaskModifier struct {
	t *Task
}

func NewTaskModifier(t *Task) *TaskModifier {
	return &TaskModifier{t: t}
}

func (m *TaskModifier) Success(res task.Result) {
	m.t.status = task.StatusSuccess
	m.t.state = m.t.state.SetResult(res)
}

func (m *TaskModifier) Failure(res task.Result) {
	m.t.status = task.StatusFailure
	m.t.state = m.t.state.SetResult(res)
}
