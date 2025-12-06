package scheduler

import (
	"context"

	stask "github.com/atmxlab/proxychecker/internal/service/task"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/queue"
	"github.com/atmxlab/proxychecker/pkg/time"
)

type Queue interface {
	PushTasks(ctx context.Context, task ...queue.Task) error
}

type ScheduleTask struct {
	queue        Queue
	timeProvider time.Provider
}

func NewSchedulerTask(queue Queue, timeProvider time.Provider) *ScheduleTask {
	return &ScheduleTask{queue: queue, timeProvider: timeProvider}
}

func (s *ScheduleTask) Execute(ctx context.Context, tasks ...stask.Task) error {
	qTasks := make([]queue.Task, 0, len(tasks))
	for _, t := range tasks {
		if err := t.Validate(); err != nil {
			return errors.Wrap(err, "t.Validate")
		}
		bytes, err := t.Marshal()
		if err != nil {
			return errors.Wrap(err, "t.Marshal")
		}

		qt := queue.
			NewTaskBuilder().
			ID(queue.NewID()).
			Kind(t.Kind().ToQueue()).
			ExternalID(t.Key()).
			Status(queue.StatusPending). // TODO: не разрешать управлять статусом снаружи
			Payload(bytes).
			CreatedAt(s.timeProvider.CurrentTime(ctx)).
			UpdatedAt(s.timeProvider.CurrentTime(ctx)).
			Build()

		qTasks = append(qTasks, qt)
	}

	if err := s.queue.PushTasks(ctx, qTasks...); err != nil {
		return errors.Wrap(err, "s.queue.PushTasks")
	}

	return nil
}
