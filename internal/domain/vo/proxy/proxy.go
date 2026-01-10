package proxy

import (
	"fmt"

	"github.com/atmxlab/proxychecker/pkg/uuid"
)

type ID string

func NewID() ID {
	return ID(uuid.MustV7().String())
}

func (id ID) String() string {
	return string(id)
}

type Protocol string

const (
	ProtocolUnknown Protocol = "unknown"
	ProtocolHTTP    Protocol = "http"
	ProtocolHTTPS   Protocol = "https"
	ProtocolSocks4  Protocol = "socks4"
	ProtocolSocks5  Protocol = "socks5"
)

func (p Protocol) String() string {
	return string(p)
}

func Protocols() []Protocol {
	return []Protocol{
		ProtocolHTTP,
		ProtocolHTTPS,
		ProtocolSocks4,
		ProtocolSocks5,
	}
}

func ProtocolFromString(s string) Protocol {
	if p, ok := TryProtocolFromString(s); ok {
		return p
	}

	panic(fmt.Sprintf("unknown proxy protocol: %s", s))
}

func TryProtocolFromString(s string) (Protocol, bool) {
	for _, p := range Protocols() {
		if s == p.String() {
			return p, true
		}
	}

	return "", false
}

type Type string

const (
	TypeUnknown     Type = "unknown"
	TypeDatacenter  Type = "datacenter"
	TypeResidential Type = "residential"
	TypeMobile      Type = "mobile"
)
