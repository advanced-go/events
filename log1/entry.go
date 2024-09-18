package log1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/events/common"
	"github.com/advanced-go/stdlib/core"
	"time"
)

// Entry - timeseries access log struct
type Entry struct {
	CustomerId string    `json:"customer-id"`
	StartTime  time.Time `json:"start-time"`
	Duration   int64     `json:"duration"`
	Traffic    string    `json:"traffic"`
	CreatedTS  time.Time `json:"created-ts"`

	Origin core.Origin `json:"origin"`
	//Region     string `json:"region"`
	//Zone       string `json:"zone"`
	//SubZone    string `json:"sub-zone"`
	//Host       string `json:"host"`
	//Route      string `json:"route"`
	//InstanceId string `json:"instance-id"`

	RequestId string `json:"request-id"`
	RelatesTo string `json:"relates-to"`
	Location  string `json:"location"`
	Protocol  string `json:"proto"`
	Method    string `json:"method"`
	From      string `json:"from"`
	To        string `json:"to"`
	Uri       string `json:"uri"`
	Path      string `json:"path"`
	Query     string `json:"query"`

	StatusCode int32  `json:"status-code"`
	Encoding   string `json:"encoding"`
	Bytes      int64  `json:"bytes"`

	Timeout        int32   `json:"timeout"`
	RateLimit      float64 `json:"rate-limit"`
	RateBurst      int32   `json:"rate-burst"`
	ControllerCode string  `json:"cc"`
}

func (Entry) Scan(columnNames []string, values []any) (e Entry, err error) {
	for i, name := range columnNames {
		switch name {
		case common.CustomerIdName:
			e.CustomerId = values[i].(string)
		case common.StartTimeName:
			e.StartTime = values[i].(time.Time)
		case common.DurationName:
			e.Duration = values[i].(int64)
		case common.TrafficName:
			e.Traffic = values[i].(string)
		case common.CreatedTSName:
			e.CreatedTS = values[i].(time.Time)

		case common.RegionName:
			e.Origin.Region = values[i].(string)
		case common.ZoneName:
			e.Origin.Zone = values[i].(string)
		case common.SubZoneName:
			e.Origin.SubZone = values[i].(string)
		case common.HostName:
			e.Origin.Host = values[i].(string)
		case common.RouteName:
			e.Origin.Route = values[i].(string)
		case common.InstanceIdName:
			e.Origin.InstanceId = values[i].(string)

		case common.RequestIdName:
			e.RequestId = values[i].(string)
		case common.RelatesToName:
			e.RelatesTo = values[i].(string)
		case common.ProtocolName:
			e.Protocol = values[i].(string)
		case common.MethodName:
			e.Method = values[i].(string)
		case common.FromName:
			e.From = values[i].(string)
		case common.ToName:
			e.To = values[i].(string)
		case common.UriName:
			e.Uri = values[i].(string)
		case common.PathName:
			e.Path = values[i].(string)

		case common.StatusCodeName:
			e.StatusCode = values[i].(int32)
		case common.EncodingName:
			e.Encoding = values[i].(string)
		case common.BytesName:
			e.Bytes = values[i].(int64)

		case common.TimeoutName:
			e.Timeout = values[i].(int32)
		case common.RateLimitName:
			e.RateLimit = values[i].(float64)
		case common.RateBurstName:
			e.RateBurst = values[i].(int32)
		case common.ControllerCodeName:
			e.ControllerCode = values[i].(string)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e Entry) Values() []any {
	return []any{
		e.CustomerId,
		e.StartTime,
		e.Duration,
		e.Traffic,
		e.CreatedTS,

		e.Origin.Region,
		e.Origin.Zone,
		e.Origin.SubZone,
		e.Origin.Host,
		e.Origin.Route,
		e.Origin.InstanceId,

		e.RequestId,
		e.RelatesTo,
		e.Protocol,
		e.Method,
		e.From,
		e.To,
		e.Uri,
		e.Path,

		e.StatusCode,
		e.Encoding,
		e.Bytes,

		e.Timeout,
		e.RateLimit,
		e.RateBurst,
		e.ControllerCode,

		//e.RouteTo,
		//e.RoutePercent,
		//e.RouteCode,
	}
}

func (Entry) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
