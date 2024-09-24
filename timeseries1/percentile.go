package timeseries1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/events/common"
	"github.com/advanced-go/stdlib/core"
)

const (
	percentileName = "percentile"
	latencyName    = "latency"
)

// PercentileThreshold - struct for percentile thresholds
type PercentileThreshold struct {
	Origin     core.Origin `json:"origin"`
	Percentile int         `json:"percentile"` // Used for latency, traffic, status codes, counter, profile
	Latency    int         `json:"latency"`    // Used for latency, saturation duration or traffic
}

func NewPercentileThreshold() PercentileThreshold {
	p := PercentileThreshold{}
	p.Percentile = 95
	p.Latency = 3000 // milliseconds
	return p
}

func (PercentileThreshold) Scan(columnNames []string, values []any) (e PercentileThreshold, err error) {
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

		case percentileName:
			e.Percentile = values[i].(int)
		case latencyName:
			e.Latency = values[i].(int)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e PercentileThreshold) Values() []any {
	return []any{
		e.Origin.Region,
		e.Origin.Zone,
		e.Origin.SubZone,
		e.Origin.Host,
		e.Origin.Route,

		e.Percentile,
		e.Latency,
	}
}

func (PercentileThreshold) Rows(entries []PercentileThreshold) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
