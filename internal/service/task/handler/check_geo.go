package handler

import (
	"context"

	"github.com/atmxlab/proxychecker/pkg/queue"
)

type CheckGEOHandler struct{}

func NewCheckGEOHandler() *CheckGEOHandler {
	return &CheckGEOHandler{}
}

func (c CheckGEOHandler) Handle(ctx context.Context, task queue.Task) error {
	return nil
}
