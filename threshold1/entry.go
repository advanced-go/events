package threshold1

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

type Entry struct {
	Origin  core.Origin `json:"origin"`
	Percent int         `json:"percent"` // Used for latency, traffic, status codes, counter, profile
	Value   int         `json:"value"`   // Used for latency, saturation duration or traffic
	Minimum int         `json:"minimum"` // Used for status codes to attenuate underflow, applied to the window interval
}

func (Entry) Scan(columnNames []string, values []any) (e Entry, err error) {
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

func (e Entry) Values() []any {
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

func (Entry) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
