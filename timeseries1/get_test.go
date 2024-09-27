package timeseries1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

func ExampleGetPercentile_Test() {
	values := make(url.Values)
	//ctx := core.NewExchangeOverrideContext(nil, core.NewExchangeOverride("", testrsc.TS1PercentileThresholdTest, ""))

	values.Add(core.RegionKey, "*")
	entries, status := get[core.Output, PercentileThreshold](nil, nil, percentileResource, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	values.Set(core.RegionKey, "us-west")
	values.Add(core.SubZoneKey, "dc1")
	entries, status = get[core.Output, PercentileThreshold](nil, nil, percentileResource, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]
	//test: get() -> [status:OK] [entries:1]

}

func ExampleGetStatusCode_Test() {
	values := make(url.Values)
	//ctx := core.NewExchangeOverrideContext(nil, core.NewExchangeOverride("", testrsc.TS1StatusCodeThresholdTest, ""))

	values.Add(core.RegionKey, "*")
	entries, status := get[core.Output, PercentileThreshold](nil, nil, statusCodeResource, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	values.Set(core.RegionKey, "us-west")
	values.Add(core.SubZoneKey, "dc1")
	entries, status = get[core.Output, PercentileThreshold](nil, nil, statusCodeResource, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]
	//test: get() -> [status:OK] [entries:1]

}
