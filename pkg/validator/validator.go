package validator

import (
	"fmt"
	"strings"

	"github.com/atmxlab/proxychecker/pkg/errors"
)

type Validator struct {
	msgs []string
}

func New() *Validator {
	return &Validator{
		msgs: []string{},
	}
}

func (v *Validator) Failed(msg string) {
	if msg != "" {
		v.msgs = append(v.msgs, msg)
	}
}

func (v *Validator) Failedf(format string, args ...any) {
	v.Failed(fmt.Sprintf(format, args...))
}

func (v *Validator) AddErr(err error) {
	if err != nil {
		v.Failed(err.Error())
	}
}

func (v *Validator) WrapErr(err error, msg string) {
	v.AddErr(errors.Wrap(err, msg))
}

func (v *Validator) Err() error {
	if len(v.msgs) == 0 {
		return nil
	}

	return errors.Wrap(
		errors.ErrInvalidArgument,
		strings.Join(v.msgs, ";"),
	)
}
