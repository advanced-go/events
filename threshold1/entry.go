package threshold1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/events/common"
	"github.com/advanced-go/stdlib/core"
	"time"
)

const (
	percentName = "percent"
	valueName   = "value"
	minimumName = "minimum"
)

type Entry struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	Route     string    `json:"route"`
	CreatedTS time.Time `json:"created-ts"`

	Percent int `json:"percent"` // Used for latency, traffic, status codes, counter, profile
	Value   int `json:"value"`   // Used for latency, saturation duration or traffic
	Minimum int `json:"minimum"` // Used for status codes to attenuate underflow, applied to the window interval
}

func (Entry) Scan(columnNames []string, values []any) (e Entry, err error) {
	for i, name := range columnNames {
		switch name {
		case common.CreatedTSName:
			e.CreatedTS = values[i].(time.Time)

		case common.RegionName:
			e.Region = values[i].(string)
		case common.ZoneName:
			e.Zone = values[i].(string)
		case common.SubZoneName:
			e.SubZone = values[i].(string)
		case common.HostName:
			e.Host = values[i].(string)
		case common.RouteName:
			e.Route = values[i].(string)

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
		e.CreatedTS,

		e.Region,
		e.Zone,
		e.SubZone,
		e.Host,
		e.Route,

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

func (e Entry) Origin() core.Origin {
	return core.Origin{
		Region:     e.Region,
		Zone:       e.Zone,
		SubZone:    e.SubZone,
		Host:       e.Host,
		InstanceId: "",
		Route:      e.Route,
	}
}
