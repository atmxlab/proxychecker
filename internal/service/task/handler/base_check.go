package handler

import (
	"context"
	"crypto/x509"
	"net"
	"net/url"
	"strings"

	"github.com/atmxlab/proxychecker/internal/domain/aggregate"
	"github.com/atmxlab/proxychecker/internal/domain/vo/task"
	"github.com/atmxlab/proxychecker/internal/port"
	"github.com/atmxlab/proxychecker/internal/service/task/payload"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/atmxlab/proxychecker/pkg/queue"
)

//go:generate mock Checker
type Checker interface {
	Run(ctx context.Context, t *aggregate.Task) (task.Result, error)
}

type BaseCheckHandler struct {
	checker     Checker
	getTaskAgg  port.GetTaskAgg
	saveTaskAgg port.SaveTaskAgg
}

func NewBaseCheckHandler(
	checker Checker,
	getTaskAgg port.GetTaskAgg,
	saveTaskAgg port.SaveTaskAgg,
) *BaseCheckHandler {
	return &BaseCheckHandler{
		checker:     checker,
		getTaskAgg:  getTaskAgg,
		saveTaskAgg: saveTaskAgg,
	}
}

func (c *BaseCheckHandler) Handle(ctx context.Context, qt queue.Task) error {
	p, err := payload.NewTaskFromBytes(qt.Payload())
	if err != nil {
		return errors.Wrap(err, "payload.NewTaskFromBytes")
	}

	t, err := c.getTaskAgg.Execute(ctx, p.ID)
	if err != nil {
		return errors.Wrap(err, "getTaskAgg.Execute")
	}

	res, err := c.checker.Run(ctx, t)
	if err != nil {
		if failureErr := t.Failure(task.Result{
			ErrorResult: &task.ErrorResult{
				Code:    c.detectErrorCode(err),
				Message: err.Error(),
			},
		}); failureErr != nil {
			return errors.Wrap(failureErr, "checker.Run")
		}
		return errors.Wrap(err, "checker.Run")
	}

	if err = t.Success(res); err != nil {
		return errors.Wrap(err, "taskAgg.Success")
	}

	if err = c.saveTaskAgg.Execute(ctx, t); err != nil {
		return errors.Wrap(err, "saveTaskAgg.Execute")
	}

	return nil
}

func (c *BaseCheckHandler) detectErrorCode(err error) task.ErrCode {
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		switch e := urlErr.Err.(type) {
		case *net.OpError:
			if e.Op == "dial" {
				if strings.Contains(e.Err.Error(), "refused") {
					return task.ErrCodeConnectionRefused
				}
				if strings.Contains(e.Err.Error(), "i/o timeout") {
					return task.ErrCodeConnectionTimeout
				}
				return task.ErrCodeNetwork
			}
			return task.ErrCodeNetwork
		case net.Error:
			if e.Timeout() {
				return task.ErrCodeTimeout
			}
			return task.ErrCodeNetwork

		case *x509.UnknownAuthorityError:
			return task.ErrCodeTLSCertUnknownAuthority

		case *x509.HostnameError:
			return task.ErrCodeTLSCertHostnameMismatch

		case *x509.CertificateInvalidError:
			return task.ErrCodeTLSCertInvalid

		default:
			msg := e.Error()
			if strings.Contains(msg, "EOF") {
				return task.ErrCodeConnectionClosedUnexpectedly
			}
			if strings.Contains(msg, "timeout") {
				return task.ErrCodeTimeout
			}
			if strings.Contains(msg, "header") {
				return task.ErrCodeHeaderReadTimeout
			}

			return task.ErrCodeUnknown
		}
	}

	msg := err.Error()
	if strings.Contains(msg, "proxyconnect") {
		return task.ErrCodeProxyConnectionFailed
	}
	if strings.Contains(msg, "certificate") {
		return task.ErrCodeTLSCertInvalid
	}

	return task.ErrCodeUnknown
}
