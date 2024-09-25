package log1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

func ExampleGetIngress_Test() {
	values := make(url.Values)
	//ctx := core.NewExchangeOverrideContext(nil, core.NewExchangeOverride("", testrsc.LOG1IngressEntry, ""))

	values.Add(core.RegionKey, "*")
	entries, status := get[core.Output, Entry](nil, nil, IngressResource, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	values.Set(core.RegionKey, "us-west")
	values.Add(core.SubZoneKey, "dc1")
	entries, status = get[core.Output, Entry](nil, nil, IngressResource, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]
	//test: get() -> [status:OK] [entries:1]

}

func ExampleGetEgress_Test() {
	values := make(url.Values)
	//ctx := core.NewExchangeOverrideContext(nil, core.NewExchangeOverride("", testrsc.LOG1EgressEntry, ""))

	values.Add(core.RegionKey, "*")
	entries, status := get[core.Output, Entry](nil, nil, EgressResource, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	values.Set(core.RegionKey, "us-west")
	values.Add(core.SubZoneKey, "dc1")
	entries, status = get[core.Output, Entry](nil, nil, EgressResource, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]
	//test: get() -> [status:OK] [entries:1]

}
