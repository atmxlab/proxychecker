package task

import (
	"fmt"
	"time"
)

type Result struct {
	ErrorResult   *ErrorResult   `json:"errorResult,omitempty"`
	LatencyResult *LatencyResult `json:"latencyResult,omitempty"`
	GEOResult     *GEOResult     `json:"geoResult,omitempty"`
}

func (r Result) String() string {
	if res := r.LatencyResult; res != nil {
		return res.String()
	}
	if res := r.GEOResult; res != nil {
		return res.String()
	}
	if r.ErrorResult != nil {
		return r.ErrorResult.String()
	}

	return ""
}

type LatencyResult struct {
	LatencyToProxy  time.Duration `json:"latencyToProxy"`
	LatencyToTarget time.Duration `json:"latencyToTarget"`
}

func (r *LatencyResult) String() string {
	return fmt.Sprintf(
		"latencyToProxy: [%s], latencyToTarget: [%s]",
		r.LatencyToProxy, r.LatencyToTarget,
	)
}

type GEOResult struct {
	ContinentCode string `json:"continentCode"`
	Continent     string `json:"continent"`
	CountryCode   string `json:"countryCode"`
	Country       string `json:"country"`
	Region        string `json:"region"`
	City          string `json:"city"`
	Timezone      string `json:"timezone"`
}

func (r *GEOResult) String() string {
	return fmt.Sprintf(
		"continentCode: [%s], continent: [%s], countryCode: [%s], region: [%s], city: [%s], timezone: [%s]",
		r.ContinentCode,
		r.Continent,
		r.CountryCode,
		r.Region,
		r.City,
		r.Timezone,
	)
}
