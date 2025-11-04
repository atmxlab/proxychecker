package queue

import "context"

type Handler interface {
	Handle(ctx context.Context, task Task) error
}

type handler struct {
	kind    Kind
	handler Handler
	options options
}

func newHandler(k Kind, h Handler, opts options) handler {
	return handler{kind: k, handler: h, options: opts}
}
