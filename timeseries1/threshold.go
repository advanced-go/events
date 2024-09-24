package timeseries1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/events/common"
	"github.com/advanced-go/stdlib/core"
)

const (
	PercentilePercent = 95
	PercentileValue   = 3000 // milliseconds
	PercentileMinimum = 0

	StatusCodePercent = 10
	StatusCodeValue   = 0
	StatusCodeMinimum = 100

	percentName = "percent"
	valueName   = "value"
	minimumName = "minimum"
)

// threshold - struct for thresholds
// Templates are of the form: 4xx, 5xx
type threshold struct {
	origin core.Origin `json:"origin"`
	//StatusCodes string      `json:"codes"`   // String of comma seperated status codes, and supporting templates
	percent int `json:"percent"` // Used for latency, traffic, status codes, counter, profile
	value   int `json:"value"`   // Used for latency, saturation duration or traffic
	minimum int `json:"minimum"` // Used for status codes to attenuate underflow, applied to the window interval
}

type PercentileThreshold struct {
	t threshold
}

func NewPercentileThreshold() PercentileThreshold {
	p := PercentileThreshold{}
	p.t.minimum = PercentileMinimum
	p.t.percent = PercentilePercent
	p.t.value = PercentileValue
	return p
}

func (p PercentileThreshold) Latency() int {
	return p.t.value
}

func (p PercentileThreshold) Percent() int {
	return p.t.percent
}

type StatusCodeThreshold struct {
	t threshold
}

func NewStatusCodeThreshold() StatusCodeThreshold {
	p := StatusCodeThreshold{}
	p.t.minimum = StatusCodeMinimum
	p.t.percent = StatusCodePercent
	p.t.value = StatusCodeValue
	return p
}

func (s StatusCodeThreshold) Percent() int {
	return s.t.percent
}

func (s StatusCodeThreshold) Minimum() int {
	return s.t.minimum
}

func (threshold) Scan(columnNames []string, values []any) (e threshold, err error) {
	for i, name := range columnNames {
		switch name {
		//case common.CreatedTSName:
		//	e.CreatedTS = values[i].(time.Time)

		case common.RegionName:
			e.origin.Region = values[i].(string)
		case common.ZoneName:
			e.origin.Zone = values[i].(string)
		case common.SubZoneName:
			e.origin.SubZone = values[i].(string)
		case common.HostName:
			e.origin.Host = values[i].(string)
		case common.RouteName:
			e.origin.Route = values[i].(string)

		case percentName:
			e.percent = values[i].(int)
		case valueName:
			e.value = values[i].(int)
		case minimumName:
			e.minimum = values[i].(int)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e threshold) Values() []any {
	return []any{
		//e.CreatedTS,

		e.origin.Region,
		e.origin.Zone,
		e.origin.SubZone,
		e.origin.Host,
		e.origin.Route,

		e.percent,
		e.value,
		e.minimum,
	}
}

func (threshold) Rows(entries []threshold) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
