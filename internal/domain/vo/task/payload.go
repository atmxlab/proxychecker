package task

import (
	"net/url"

	"github.com/atmxlab/proxychecker/pkg/validator"
)

type Payload struct {
	TargetURL *TargetURL `json:"target_url"`
}

func NewEmptyPayload() Payload {
	return Payload{}
}

type TargetURL struct {
	URL string `json:"url"`
}

func (u TargetURL) Validate() error {
	v := validator.New()

	if u.URL == "" {
		v.Failed("empty target url")
	} else {
		_, err := url.Parse(u.URL)
		v.AddErr(err)
	}

	return v.Err()
}
