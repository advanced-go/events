package threshold1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

func ExampleGet_Test() {
	values := make(url.Values)
	//ctx := core.NewExchangeOverrideContext(nil, core.NewExchangeOverride("", testrsc.TS1EgressEntry, ""))

	values.Add(core.RegionKey, "*")
	entries, _, status := get[core.Output, Threshold](nil, nil, "", values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	values.Set(core.RegionKey, "us-west")
	values.Add(core.SubZoneKey, "dc1")
	entries, _, status = get[core.Output, Threshold](nil, nil, "", values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]
	//test: get() -> [status:OK] [entries:1]

}
