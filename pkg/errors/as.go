package errors

import "errors"

func As[T any](err error, target *T) bool {
	return errors.As(err, target)
}
