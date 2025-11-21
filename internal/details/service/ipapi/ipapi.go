package ipapi

import (
	"context"
	"encoding/json"

	"github.com/atmxlab/proxychecker/internal/details/client"
	"github.com/atmxlab/proxychecker/pkg/errors"
)

type Output struct {
	Status        string  `json:"status"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Region        string  `json:"region"`
	RegionName    string  `json:"regionName"`
	City          string  `json:"city"`
	District      string  `json:"district"`
	Zip           string  `json:"zip"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	Timezone      string  `json:"timezone"`
	Offset        int     `json:"offset"`
	Currency      string  `json:"currency"`
	Isp           string  `json:"isp"`
	Org           string  `json:"org"`
	As            string  `json:"as"`
	Asname        string  `json:"asname"`
	Mobile        bool    `json:"mobile"`
	Proxy         bool    `json:"proxy"`
	Hosting       bool    `json:"hosting"`
	Query         string  `json:"query"`
}

type Service struct {
	client client.Client
}

func New(client client.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Get(ctx context.Context) (Output, error) {
	bytes, err := s.client.Get(ctx, "http://ip-api.com/json")
	if err != nil {
		return Output{}, errors.Wrap(err, "s.client.Get")
	}

	var output Output
	if err = json.Unmarshal(bytes, &output); err != nil {
		return Output{}, errors.Wrap(err, "json.Unmarshal")
	}

	return output, nil
}
