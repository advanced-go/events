package timeseries1

import (
	"fmt"
	"github.com/advanced-go/events/testrsc"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

func _ExampleGet() {
	values := make(url.Values)
	ctx := core.NewExchangeOverrideContext(nil, core.NewExchangeOverride("", testrsc.TS1EgressEntry, ""))

	values.Add(core.RegionKey, "us-west")
	entries, _, status := get[core.Output, Entry](ctx, nil, "", values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	values.Add(core.SubZoneKey, "dc1")
	entries, _, status = get[core.Output, Entry](ctx, nil, "", values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]
	//test: get() -> [status:OK] [entries:1]

}

func ExampleGet_Test() {
	values := make(url.Values)
	//ctx := core.NewExchangeOverrideContext(nil, core.NewExchangeOverride("", testrsc.TS1EgressEntry, ""))

	values.Add(core.RegionKey, "*")
	entries, _, status := get[core.Output, Entry](nil, nil, "", values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//values.Add(core.SubZoneKey, "dc1")
	entries, _, status = get[core.Output, Entry](nil, nil, "", values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]
	//test: get() -> [status:OK] [entries:2]

}
