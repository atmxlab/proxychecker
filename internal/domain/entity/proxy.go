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
