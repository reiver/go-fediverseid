package fediverseid_test

import (
	"fmt"

	"github.com/reiver/go-fediverseid"
)

func ExampleFediverseID_newFediverseID_keyword() {

	const name string = "joeblow"
	const host string = "host.example"

	var fid *fediverseid.FediverseID = new(fediverseid.FediverseID).ChainSetName(name).ChainSetHost(host)

	fmt.Printf("fediverse-id: %s", fid)

	// Output:
	// fediverse-id: @joeblow@host.example
}
