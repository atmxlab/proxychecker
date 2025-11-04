package task

type ErrCode string

const (
	ErrCodeUnknown     ErrCode = "unknown"
	ErrCodeConnTimeout ErrCode = "connection timeout"
)

type ErrorResult struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message,omitempty"`
}
