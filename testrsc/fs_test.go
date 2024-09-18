package testrsc

import (
	"fmt"
	"github.com/advanced-go/stdlib/io"
)

func ExampleReadFile() {
	name := "file:///f:/files/timeseries1/egress-entry.json"
	bytes, status := io.ReadFile(name)
	fmt.Printf("test: ReadFile() -> [buff:%v] [status:%v]\n", len(bytes), status)

	//Output:
	//test: ReadFile() -> [buff:1512] [status:OK]

}
