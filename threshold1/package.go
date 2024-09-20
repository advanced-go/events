package threshold1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

const (
	PkgPath           = "github/advanced-go/events/threshold1"
	Route             = "threshold"
	thresholdResource = "threshold"
	thresholdSelect   = "select * from threshold"
)

// Get - threshold1 GET
func Get(r *http.Request, path string) (entries []Threshold, h2 http.Header, status *core.Status) {
	if r == nil {
		return entries, h2, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
	}
	return nil, h2, core.StatusNotFound() //OK()get[core.Log, Entry](r.Context(), core.AddRequestId(r.Header), rsc, r.URL.Query())
}

func GetIngressPercentile(ctx context.Context, origin core.Origin) (Threshold, *core.Status) {
	return Threshold{}, core.StatusOK()
}

func GetEgressStatusCode(ctx context.Context, origin core.Origin) (Threshold, *core.Status) {
	return Threshold{}, core.StatusOK()
}

func GetProfile(ctx context.Context) (*Profile, *core.Status) {
	return NewProfile(), core.StatusOK()
}
