package task

import (
	"time"
)

type Result struct {
	ErrorResult   *ErrorResult   `json:"errorResult,omitempty"`
	LatencyResult *LatencyResult `json:"latencyResult,omitempty"`
	GEOResult     *GEOResult     `json:"geoResult,omitempty"`
}

type LatencyResult struct {
	LatencyToProxy  time.Duration `json:"latencyToProxy"`
	LatencyToTarget time.Duration `json:"latencyToTarget"`
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
