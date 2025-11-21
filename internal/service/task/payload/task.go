package payload

import (
	"encoding/json"

	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type Task struct {
	ID task.ID
}

func NewTaskFromBytes(b []byte) (Task, error) {
	var t Task
	if err := t.Unmarshal(b); err != nil {
		return Task{}, errors.Wrap(err, "t.Unmarshal")
	}

	return t, nil
}

func (t Task) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return nil, errors.Wrap(err, "json.Marshal")
	}

	return bytes, nil
}

func (t *Task) Unmarshal(bytes []byte) error {
	if err := json.Unmarshal(bytes, t); err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}

	return nil
}
