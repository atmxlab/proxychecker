package entity

import (
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
)

type ProxyBuilder struct {
	id        proxy.ID
	url       string
	protocol  proxy.Protocol
	createdAt time.Time
	updatedAT time.Time
}

func NewProxyBuilder() *ProxyBuilder {
	return &ProxyBuilder{}
}

func (b *ProxyBuilder) ID(id proxy.ID) *ProxyBuilder {
	b.id = id
	return b
}

func (b *ProxyBuilder) URL(url string) *ProxyBuilder {
	b.url = url
	return b
}

func (b *ProxyBuilder) Protocol(proto proxy.Protocol) *ProxyBuilder {
	b.protocol = proto
	return b
}

func (b *ProxyBuilder) CreatedAt(t time.Time) *ProxyBuilder {
	b.createdAt = t
	return b
}

func (b *ProxyBuilder) UpdatedAt(t time.Time) *ProxyBuilder {
	b.updatedAT = t
	return b
}

func (b *ProxyBuilder) Build() *Proxy {
	return &Proxy{
		id:        b.id,
		url:       b.url,
		protocol:  b.protocol,
		createdAt: b.createdAt,
		updatedAT: b.updatedAT,
	}
}
