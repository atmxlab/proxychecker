package waiter

import (
	"time"

	"github.com/atmxlab/proxychecker/pkg/errors"
)

const (
	DefaultTimeout = 10 * time.Second
	DefaultTick    = 300 * time.Millisecond
)

type Config struct {
	timeout time.Duration
	tick    time.Duration
}

type Option func(*Config)

func WithTimeout(timeout time.Duration) Option {
	return func(opt *Config) {
		opt.timeout = timeout
	}
}

func WithTick(tick time.Duration) Option {
	return func(opt *Config) {
		opt.tick = tick
	}
}

func Wait(cb func() (bool, error), opts ...Option) error {
	config := Config{
		timeout: DefaultTimeout,
		tick:    DefaultTick,
	}
	for _, opt := range opts {
		opt(&config)
	}

	timer := time.NewTimer(config.timeout)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			return errors.New("wait timeout")
		default:
			done, err := cb()
			if err != nil {
				return errors.Wrap(err, "waiting for callback")
			}
			if done {
				return nil
			}
			time.Sleep(config.tick)
		}
	}
}
