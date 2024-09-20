package log1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"strings"
)

const (
	PkgPath         = "github/advanced-go/events/log1"
	Route           = "timeseries"
	EgressResource  = "egress"
	IngressResource = "ingress"
)

// Get - timeseries1 GET
func Get(r *http.Request, path string) (entries []Entry, h2 http.Header, status *core.Status) {
	if r == nil {
		return entries, h2, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
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

// Put - timeseries1 PUT, with optional content override
func Put(r *http.Request, path string, body []Entry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status.WithRequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}

func IngressQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}

func EgressQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}
