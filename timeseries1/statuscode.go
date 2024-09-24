package timeseries1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/events/common"
	"github.com/advanced-go/stdlib/core"
)

const (
	percentageName = "percentage"
	countName      = "count"
)

// StatusCodeThreshold - struct for status code thresholds
type StatusCodeThreshold struct {
	Origin     core.Origin `json:"origin"`
	Percentage int         `json:"percentage"`
	Count      int         `json:"count"`
}

func NewStatusCodeThreshold() StatusCodeThreshold {
	s := StatusCodeThreshold{}
	s.Percentage = 0
	s.Count = 0
	return s
}

func (StatusCodeThreshold) Scan(columnNames []string, values []any) (e StatusCodeThreshold, err error) {
	for i, name := range columnNames {
		switch name {
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

		case percentageName:
			e.Percentage = values[i].(int)
		case countName:
			e.Count = values[i].(int)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e StatusCodeThreshold) Values() []any {
	return []any{
		e.Origin.Region,
		e.Origin.Zone,
		e.Origin.SubZone,
		e.Origin.Host,
		e.Origin.Route,

		e.Percentage,
		e.Count,
	}
}

func (StatusCodeThreshold) Rows(entries []StatusCodeThreshold) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
