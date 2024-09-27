package log2

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	PkgPath          = "github/advanced-go/events/log2"
	Route            = "log-events"
	ingressEntryPath = "log/ingress/entry"
	egressEntryPath  = "log/egress/entry"
)

// Get - log2 GET
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
	case ingressEntryPath, egressEntryPath:
		t, status := get[core.Log, Entry](r.Context(), core.AddRequestId(r.Header), path, r.URL.Query())
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
		return nil, h2, core.NewStatusError(http.StatusBadRequest, errors.New("error: resource is not ingress or egress"))
	}
}

// Put - log2 PUT, with optional content override
func Put(r *http.Request, path string, body []Entry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if r.Header.Get(core.XFrom) == "" {
		return httpPut[core.Log](r, path, body)
	}
	return httpPut[core.Output](r, path, body)
}

func httpPut[E core.ErrorHandler](r *http.Request, path string, body []Entry) (http.Header, *core.Status) {
	var e E

	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			e.Handle(status.WithRequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return put[E](r.Context(), core.AddRequestId(r.Header), body)
}

func IngressQuery(ctx context.Context, h http.Header, values url.Values) ([]Entry, *core.Status) {
	return get[core.Log, Entry](ctx, h, ingressEntryPath, values)
}

func EgressQuery(ctx context.Context, h http.Header, values url.Values) ([]Entry, *core.Status) {
	return get[core.Log, Entry](ctx, h, egressEntryPath, values)
}
