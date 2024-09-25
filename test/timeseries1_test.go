package test

import (
	http2 "github.com/advanced-go/events/http"
	"github.com/advanced-go/events/log1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/core/coretest"
	httpt "github.com/advanced-go/stdlib/httpx/httpxtest"
	"net/http"
	"reflect"
	"testing"
)

func TestExchange1(t *testing.T) {
	tests := []struct {
		name   string
		req    *http.Request
		resp   *http.Response
		status *core.Status
	}{
		//{name: "read-request-error", req: readRequest("", t), resp: readResponse(testrsc.TS1GetRespURL, t), status: core.StatusOK()},
		//{name: "get-entry", req: httpt.NewRequestTest(testrsc.TS1GetReq, t), resp: httpt.NewResponseTest(testrsc.TS1GetResp, t), status: core.StatusOK()},
	}
	for _, tt := range tests {
		success := true
		t.Run(tt.name, func(t *testing.T) {
			resp, status := http2.Exchange(tt.req)
			if tt.status != nil && status.Code != tt.status.Code {
				t.Errorf("Exchange() got status : %v, want status : %v, error : %v", status.Code, tt.status.Code, status.Err)
				success = false
			}
			if success && resp.StatusCode != tt.resp.StatusCode {
				t.Errorf("Exchange() got status code : %v, want status code : %v", resp.StatusCode, tt.resp.StatusCode)
				success = false
			}
			var gotT []log1.Entry
			var wantT []log1.Entry
			if success {
				gotT, wantT, success = httpt.Deserialize[coretest.Output, []log1.Entry](resp.Body, tt.resp.Body, t)
			}
			if success {
				if !reflect.DeepEqual(gotT, wantT) {
					t.Errorf("Exchange() got = %v, want %v", gotT, wantT)
				}
			}
		})
	}
}

/*
	if !reflect.DeepEqual(got, tt.want) {
		t.Errorf("Exchange() got = %v, want %v", got, tt.want)
	}
	if !reflect.DeepEqual(got1, tt.want1) {
		t.Errorf("Exchange() got1 = %v, want %v", got1, tt.want1)
	}

*/
