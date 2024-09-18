package timeseries1

import (
	"context"
	"errors"
	"github.com/advanced-go/events/common"
	"github.com/advanced-go/events/module"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

// put - function to Put a set of entries into a datastore
func put[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, body []T) (h2 http.Header, status *core.Status) {
	var e E

	h2 = httpx.SetHeader(h2, httpx.ContentType, httpx.ContentTypeText)
	if len(body) == 0 {
		status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: no entries found"))
		return nil, status
	}
	h = httpx.SetHeader(h, core.XFrom, module.Authority)
	_, status = pgxsql.InsertT[T](ctx, h, common.AccessLogResource, common.AccessLogInsert, body)
	if !status.OK() {
		e.Handle(status.WithRequestId(h))
	}
	return
}
