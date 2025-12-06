package task

import "fmt"

type ErrCode string

const (
	ErrCodeUnknown     ErrCode = "unknown"
	ErrCodeConnTimeout ErrCode = "connection timeout"
)

type ErrorResult struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message,omitempty"`
}

func (e *ErrorResult) String() string {
	return fmt.Sprintf("code: [%s], message: [%s]", e.Code, e.Message)
}
