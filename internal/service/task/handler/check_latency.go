package handler

import (
	"context"

	"github.com/atmxlab/proxychecker/pkg/queue"
)

type CheckLatencyHandler struct{}

func NewCheckLatencyHandler() *CheckLatencyHandler {
	return &CheckLatencyHandler{}
}

func (c CheckLatencyHandler) Handle(ctx context.Context, task queue.Task) error {
}
