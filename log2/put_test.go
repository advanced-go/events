package log2

import (
	"fmt"
	"github.com/advanced-go/events/testrsc"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
)

func ExamplePut() {
	entries, _ := json.New[[]Entry](testrsc.LOG2EgressEntry, nil)

	ex := core.NewExchangeOverride("", "", json.StatusTimeoutUri)
	ctx := core.NewExchangeOverrideContext(nil, ex)

	_, status := put[core.Output, Entry](ctx, nil, entries)
	fmt.Printf("test: put() -> [status:%v]\n", status)

	//Output:
	//test: put() -> [status:Timeout]

}
