package timeseries2

import (
	"fmt"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/json"
	"time"
)

var (
	entryData = []Entry{
		{Region: "us-west1", Zone: "a", Host: "www.host1.com", Duration: 100, Traffic: access.IngressTraffic, Route: "host", Timeout: 2000, RateLimit: 98.5, RateBurst: 10, ControllerCode: "RL", StartTime: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-west1", Zone: "a", Host: "www.host2.com", Duration: 85, Traffic: access.IngressTraffic, Route: "host", Timeout: 1500, RateLimit: 100, RateBurst: 10, ControllerCode: "", StartTime: time.Date(2024, 6, 10, 7, 120, 55, 0, time.UTC)},
		{Region: "us-central1", Zone: "c", Host: "www.host1.com", Duration: 200, Traffic: access.IngressTraffic, Route: "host", Timeout: 300, RateLimit: 98.5, RateBurst: 10, ControllerCode: "RL", StartTime: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "c", Host: "www.host2.com", Duration: 750, Traffic: access.IngressTraffic, Route: "host", Timeout: 500, RateLimit: 100, RateBurst: 10, ControllerCode: "TO", StartTime: time.Date(2024, 6, 10, 7, 120, 55, 0, time.UTC)},
	}
)

func _ExampleEntry() {
	buf, status := json.Marshal(entryData)
	if !status.OK() {
		fmt.Printf("test: Entry{} -> [status:%v]\n", status)
	} else {
		fmt.Printf("test: Entry{} -> %v\n", string(buf))
	}

	//Output:
	//fail

}
