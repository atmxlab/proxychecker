package task

import "fmt"

type ErrCode string

const (
	ErrCodeUnknown                      ErrCode = "unknown"
	ErrCodeNetwork                      ErrCode = "network: unknown"
	ErrCodeTimeout                      ErrCode = "timeout: unknown"
	ErrCodeHeaderReadTimeout            ErrCode = "timeout: header read"
	ErrCodeConnectionRefused            ErrCode = "connection: refused"
	ErrCodeConnectionTimeout            ErrCode = "connection: timeout"
	ErrCodeConnectionClosedUnexpectedly ErrCode = "connection: closed unexpectedly"
	ErrCodeTLSCertUnknownAuthority      ErrCode = "tls certificate: unknown authority"
	ErrCodeTLSCertHostnameMismatch      ErrCode = "tls certificate: hostname mismatch"
	ErrCodeTLSCertInvalid               ErrCode = "tls certificate: invalid"
	ErrCodeProxyConnectionFailed        ErrCode = "proxy: connection failed"
)

func (e ErrCode) String() string {
	return string(e)
}

type ErrorResult struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message,omitempty"`
}

func (e *ErrorResult) String() string {
	return fmt.Sprintf("code: [%s], message: [%s]", e.Code, e.Message)
}
