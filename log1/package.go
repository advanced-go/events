package log1

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"strings"
)

const (
	PkgPath         = "github/advanced-go/events/log1"
	Route           = "log-events"
	EgressResource  = "egress"
	IngressResource = "ingress"
)

// Get - log1 GET
func Get[E core.ErrorHandler](r *http.Request, path string) ([]byte, http.Header, *core.Status) {
	var e E

	h2 := httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeText)
	if r == nil {
		status := core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
		e.Handle(status)
		return nil, nil, status
	}
	if strings.Contains(path, EgressResource) {
		t, status := get[E, Entry](r.Context(), core.AddRequestId(r.Header), EgressResource, r.URL.Query())
		if !status.OK() {
			return nil, h2, status
		}
		buf, status1 := json2.Marshal(t)
		if !status1.OK() {
			e.Handle(status1)
			return nil, h2, status1
		}
		return buf, httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeJson), status1
	}
	if strings.Contains(path, IngressResource) {
		t, status := get[E, Entry](r.Context(), core.AddRequestId(r.Header), IngressResource, r.URL.Query())
		if !status.OK() {
			return nil, h2, status
		}
		buf, status1 := json2.Marshal(t)
		if !status1.OK() {
			e.Handle(status1)
			return nil, h2, status1
		}
		return buf, httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeJson), status1
	}
	status := core.NewStatusError(http.StatusBadRequest, errors.New("error: resource is not ingress or egress"))
	e.Handle(status)
	return nil, h2, status
}

// Put - log1 PUT, with optional content override
func Put[E core.ErrorHandler](r *http.Request, path string, body []Entry) (http.Header, *core.Status) {
	var e E

	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
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
