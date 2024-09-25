package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/events/module"
	"github.com/advanced-go/events/timeseries1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

func timeseriesExchange[E core.ErrorHandler](r *http.Request, p *uri.Parsed) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if p == nil {
		p1, status := httpx.ValidateURL(r.URL, module.Authority)
		if !status.OK() {
			return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
		}
		p = p1
	}

	switch r.Method {
	case http.MethodGet:
		return timeseriesGet[E](r, p)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
	}
}

func timeseriesGet[E core.ErrorHandler](r *http.Request, p *uri.Parsed) (resp *http.Response, status *core.Status) {
	var entries any
	var h2 http.Header

	switch p.Version {
	case ver1, "":
		entries, h2, status = timeseries1.Get[core.Log](r, p.Path)
	case ver2:
		//entries, h2, status = timeseries2.Get(r, p.Path)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
	}
	if !status.OK() {
		resp, _ = httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
		return resp, status
	}
	return httpx.NewResponse[E](status.HttpCode(), h2, entries)

}
