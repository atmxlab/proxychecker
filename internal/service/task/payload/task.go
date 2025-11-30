package payload

import (
	"encoding/json"

	"github.com/atmxlab/proxychecker/internal/domain/vo/checker"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	stask "github.com/atmxlab/proxychecker/internal/service/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/validator"
)

type Task struct {
	ID          task.ID
	CheckerKind checker.Kind
}

func NewTaskFromBytes(b []byte) (Task, error) {
	var t Task
	if err := t.Unmarshal(b); err != nil {
		return Task{}, errors.Wrap(err, "t.Unmarshal")
	}

	return t, nil
}

func (t Task) Key() string {
	return t.ID.String()
}

func (t Task) Kind() stask.Kind {
	return stask.FromDomainTask(t.CheckerKind)
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

func (t Task) Validate() error {
	v := validator.New()

	if t.ID == "" {
		v.Failed("ID is required")
	}
	if t.CheckerKind == checker.KindUnknown {
		v.Failed("CheckerKind is unknown")
	}

	return v.Err()
}
