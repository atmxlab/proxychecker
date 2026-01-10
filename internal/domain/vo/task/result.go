package task

import (
	"fmt"
	"time"

	"github.com/atmxlab/proxychecker/internal/domain/vo/proxy"
)

type Result struct {
	ErrorResult      *ErrorResult      `json:"errorResult,omitempty"`
	LatencyResult    *LatencyResult    `json:"latencyResult,omitempty"`
	GEOResult        *GEOResult        `json:"geoResult,omitempty"`
	ExternalIPResult *ExternalIPResult `json:"externalIPResult,omitempty"`
	URLResult        *URLResult        `json:"urlResult,omitempty"`
	HTTPSResult      *HTTPSResult      `json:"httpsResult,omitempty"`
	MITMResult       *MITMResult       `json:"hasMITM,omitempty"`
	TypeResult       *TypeResult       `json:"typeResult,omitempty"`
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
	if r.ExternalIPResult != nil {
		return r.ExternalIPResult.String()
	}
	if r.URLResult != nil {
		return r.URLResult.String()
	}
	if r.HTTPSResult != nil {
		return r.HTTPSResult.String()
	}
	if r.MITMResult != nil {
		return r.MITMResult.String()
	}
	if r.TypeResult != nil {
		return r.TypeResult.String()
	}

	return ""
}

type LatencyResult struct {
	FromHostToProxyRoundTrip   time.Duration `json:"fromHostToProxyRoundTrip"`
	FromHostToTargetRoundTrip  time.Duration `json:"fromHostToTargetRoundTrip"`
	FromProxyToTargetRoundTrip time.Duration `json:"fromProxyToTargetRoundTrip"`
}

func (r *LatencyResult) String() string {
	return fmt.Sprintf(
		"fromHostToProxyRoundTrip: [%s], fromHostToTargetRoundTrip: [%s], fromProxyToTargetRoundTrip: [%s]",
		r.FromHostToProxyRoundTrip, r.FromHostToTargetRoundTrip, r.FromProxyToTargetRoundTrip,
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

type ExternalIPResult struct {
	IP string
}

func (r *ExternalIPResult) String() string {
	return fmt.Sprintf("external ip: [%s]", r.IP)
}

type URLResult struct {
	IsAvailable bool
	URL         string
	StatusCode  int
}

func (r *URLResult) String() string {
	return fmt.Sprintf("url: [%s], is available: [%t], status: [%d]", r.URL, r.IsAvailable, r.StatusCode)
}

type HTTPSResult struct {
	IsAvailable bool `json:"isAvailable"`
}

func (r *HTTPSResult) String() string {
	return fmt.Sprintf("isAvailable: [%t]", r.IsAvailable)
}

type MITMResult struct {
	HasMITM bool `json:"hasMITM"`
}

func (r *MITMResult) String() string {
	return fmt.Sprintf("hasMIT: [%t]", r.HasMITM)
}

type TypeResult struct {
	Type proxy.Type
}

func (r *TypeResult) String() string {
	return fmt.Sprintf("type: [%s]", r.Type)
}
