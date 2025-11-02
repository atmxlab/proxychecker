package proxy

type ID string

type Protocol string

const (
	ProtocolUnknown Protocol = "unknown"
	ProtocolHTTP    Protocol = "http"
	ProtocolHTTPS   Protocol = "https"
	ProtocolSocks4  Protocol = "socks4"
	ProtocolSocks5  Protocol = "socks5"
)
