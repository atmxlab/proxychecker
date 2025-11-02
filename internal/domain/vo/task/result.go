package task

import (
	"time"
)

type Result struct {
	LatencyResult *LatencyResult `json:"latencyResult,omitempty"`
	GEOResult     *GEOResult     `json:"geoResult,omitempty"`
}

type LatencyResult struct {
	LatencyToProxy  time.Duration `json:"latencyToProxy"`
	LatencyToTarget time.Duration `json:"latencyToTarget"`
}

type GEOResult struct {
	CountryCode string `json:"countryCode"`
	State       string `json:"state"`
	City        string `json:"city"`
}
