package aggregate

import (
	"github.com/atmxlab/proxychecker/internal/domain/entity"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type Task struct {
	task  *entity.Task
	proxy *entity.Proxy
}

func NewTask(task *entity.Task, proxy *entity.Proxy) *Task {
	return &Task{task: task, proxy: proxy}
}

func (t *Task) Task() *entity.Task {
	return t.task
}

func (t *Task) Proxy() *entity.Proxy {
	return t.proxy
}

func (t *Task) Success(res task.Result) error {
	err := t.task.Modify(func(m *entity.TaskModifier) error {
		m.Success(res)
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "t.Modify")
	}

	return nil
}

func (t *Task) Failure(res task.Result) error {
	err := t.task.Modify(func(m *entity.TaskModifier) error {
		m.Failure(res)
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "t.Modify")
	}

	return nil
}
