package entity

import (
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
)

type Proxy struct {
	id        proxy.ID
	url       string
	protocol  proxy.Protocol
	createdAt time.Time
	updatedAT time.Time
}

func (p Proxy) URL() string {
	return p.url
}

func (p Proxy) ID() proxy.ID {
	return p.id
}

func (p Proxy) Protocol() proxy.Protocol {
	return p.protocol
}

func (p Proxy) CreatedAt() time.Time {
	return p.createdAt
}

func (p Proxy) UpdatedAT() time.Time {
	return p.updatedAT
}
