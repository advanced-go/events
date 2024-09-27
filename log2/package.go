package log2

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
)

const (
	PkgPath          = "github/advanced-go/events/log2"
	Route            = "log-events"
	egressResource   = "egress"
	ingressResource  = "ingress"
	ingressEntryPath = "log/ingress/entry"
	egressEntryPath  = "log/egress/entry"
)

// Get - log2 GET
func Get[E core.ErrorHandler](r *http.Request, path string) ([]byte, http.Header, *core.Status) {
	var e E

	if r == nil {
		status := core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
		e.Handle(status)
		return nil, nil, status
	}
	h2 := httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeText)
	switch path {
	case ingressEntryPath, egressEntryPath:
		t, status := get[E, Entry](r.Context(), core.AddRequestId(r.Header), path, r.URL.Query())
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
func Put[E core.ErrorHandler](r *http.Request, path string, body []Entry) (http.Header, *core.Status) {
	var e E

	if r == nil {
		status := core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
		e.Handle(status)
		return nil, status
	}
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
