package httpbin

import (
	"context"
	"encoding/json"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

const Nonce = "XqA6aewNkQtF"

type Output struct {
	Args struct {
		Nonce string `json:"nonce"`
	} `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

type Service struct {
	client client.Client
}

func New(client client.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Get(ctx context.Context) (Output, error) {
	bytes, err := s.client.Get(ctx, "http://httpbin.org/get?nonce="+Nonce)
	if err != nil {
		return Output{}, errors.Wrap(err, "s.client.Get")
	}

	var output Output
	if err = json.Unmarshal(bytes, &output); err != nil {
		return Output{}, errors.Wrap(err, "json.Unmarshal")
	}

	return output, nil
}
