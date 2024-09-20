package common

import (
	"errors"
	"fmt"
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

func (Threshold) Scan(columnNames []string, values []any) (e Threshold, err error) {
	for i, name := range columnNames {
		switch name {
		//case common.CreatedTSName:
		//	e.CreatedTS = values[i].(time.Time)

		case RegionName:
			e.Origin.Region = values[i].(string)
		case ZoneName:
			e.Origin.Zone = values[i].(string)
		case SubZoneName:
			e.Origin.SubZone = values[i].(string)
		case HostName:
			e.Origin.Host = values[i].(string)
		case RouteName:
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
