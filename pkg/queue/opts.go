package queue

type options struct {
	workerCount uint
}

func newOptions() options {
	return options{}
}

type Option func(o *options)

func WithWorkerCount(count uint) Option {
	return func(o *options) {
		o.workerCount = count
	}
}
