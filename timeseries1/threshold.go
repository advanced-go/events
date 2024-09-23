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

// Threshold - struct for thresholds
// Templates are of the form: 4xx, 5xx
type Threshold struct {
	Origin core.Origin `json:"origin"`
	//StatusCodes string      `json:"codes"`   // String of comma seperated status codes, and supporting templates
	Percent int `json:"percent"` // Used for latency, traffic, status codes, counter, profile
	Value   int `json:"value"`   // Used for latency, saturation duration or traffic
	Minimum int `json:"minimum"` // Used for status codes to attenuate underflow, applied to the window interval
}

type PercentileThreshold struct {
	T Threshold
}

func NewPercentileThreshold() *PercentileThreshold {
	p := new(PercentileThreshold)
	p.T.Minimum = PercentileMinimum
	p.T.Percent = PercentilePercent
	p.T.Value = PercentileValue
	return p
}

func (p *PercentileThreshold) Latency() int {
	return p.T.Value
}

func (p *PercentileThreshold) Percent() int {
	return p.T.Percent
}

type StatusCodeThreshold struct {
	T Threshold
}

func NewStatusCodeThreshold() *StatusCodeThreshold {
	p := new(StatusCodeThreshold)
	p.T.Minimum = StatusCodeMinimum
	p.T.Percent = StatusCodePercent
	p.T.Value = StatusCodeValue
	return p
}

func (s *StatusCodeThreshold) Percent() int {
	return s.T.Percent
}

func (s *StatusCodeThreshold) Minimum() int {
	return s.T.Minimum
}

/*
func InitPercentileThreshold(t *Threshold) {
	if t != nil {
		t.Minimum = PercentileMinimum
		t.Percent = PercentilePercent
		t.Value = PercentileValue
	}
}

func InitStatusCodeThreshold(t *Threshold) {
	if t != nil {
		t.Minimum = StatusCodeMinimum
		t.Percent = StatusCodePercent
		t.Value = StatusCodeValue
	}
}


*/

func (Threshold) Scan(columnNames []string, values []any) (e Threshold, err error) {
	for i, name := range columnNames {
		switch name {
		//case common.CreatedTSName:
		//	e.CreatedTS = values[i].(time.Time)

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

		case percentName:
			e.Percent = values[i].(int)
		case valueName:
			e.Value = values[i].(int)
		case minimumName:
			e.Minimum = values[i].(int)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e Threshold) Values() []any {
	return []any{
		//e.CreatedTS,

		e.Origin.Region,
		e.Origin.Zone,
		e.Origin.SubZone,
		e.Origin.Host,
		e.Origin.Route,

		e.Percent,
		e.Value,
		e.Minimum,
	}
}

func (Threshold) Rows(entries []Threshold) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
