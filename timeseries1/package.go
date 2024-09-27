package timeseries1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"time"
)

const (
	PkgPath        = "github/advanced-go/events/timeseries1"
	Route          = "timeseries"
	percentilePath = "timeseries/percentile-threshold"
	statusCodePath = "timeseries/status-code-threshold"
)

type TimeUTC time.Time

// Get - timeseries1 GET
func Get(r *http.Request, path string) ([]byte, http.Header, *core.Status) {
	if r == nil {
		status := core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
		return nil, nil, status
	}
	if r.Header.Get(core.XFrom) == "" {
		return httpGet[core.Log](r, path)
	}
	return httpGet[core.Output](r, path)
}

func httpGet[E core.ErrorHandler](r *http.Request, path string) ([]byte, http.Header, *core.Status) {
	var e E

	h2 := httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeText)
	switch path {
	case percentilePath:
		t, status := get[E, PercentileThreshold](r.Context(), core.AddRequestId(r.Header), percentileResource, r.URL.Query())
		if !status.OK() {
			return nil, h2, status
		}
		buf, status1 := json2.Marshal(t)
		if !status1.OK() {
			e.Handle(status1)
			return nil, h2, status1
		}
		return buf, httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeJson), status1
	case statusCodePath:
		t, status := get[E, StatusCodeThreshold](r.Context(), core.AddRequestId(r.Header), statusCodeResource, r.URL.Query())
		if !status.OK() {
			return nil, h2, status
		}
		buf, status1 := json2.Marshal(t)
		if !status1.OK() {
			e.Handle(status1)
			return nil, h2, status1
		}
		return buf, httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeJson), status1
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New("error: resource is not ingress or egress"))
		return nil, nil, status
	}
}

// PercentileThresholdSLO - ingress host, pre-calculated percentile thresholds
func PercentileThresholdSLO(ctx context.Context, origin core.Origin) (PercentileThreshold, *core.Status) {
	return NewPercentileThreshold(), core.StatusOK()
}

// PercentileThresholdQuery - ingress host, queryable percentile thresholds
func PercentileThresholdQuery(ctx context.Context, origin core.Origin, from, to TimeUTC) (PercentileThreshold, *core.Status) {
	return NewPercentileThreshold(), core.StatusOK()
}

// StatusCodeThresholdQuery - egress route, queryable status code thresholds
func StatusCodeThresholdQuery(ctx context.Context, origin core.Origin, from, to TimeUTC, statusCodes string) (StatusCodeThreshold, *core.Status) {
	return NewStatusCodeThreshold(), core.StatusOK()
}

// GetProfile - retrieve traffic profile
func GetProfile(ctx context.Context) (*Profile, *core.Status) {
	return NewProfile(), core.StatusOK()
}
