package timeseries2

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"strings"
)

const (
	PkgPath         = "github/advanced-go/events/timeseries2"
	EgressResource  = "egress"
	IngressResource = "ingress"
)

// Get - timeseries2 resource GET
func Get(r *http.Request, path string) (entries []Entry, h2 http.Header, status *core.Status) {
	if r == nil {
		return entries, h2, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is nil"))
	}
	rsc := ""
	if strings.Contains(path, EgressResource) {
		rsc = EgressResource
	} else {
		if strings.Contains(path, IngressResource) {
			rsc = IngressResource
		} else {
			return nil, h2, core.NewStatusError(http.StatusBadRequest, errors.New("error: resource is not ingress or egress"))
		}
	}
	return get[core.Log, Entry](r.Context(), core.AddRequestId(r.Header), rsc, r.URL.Query())
}
