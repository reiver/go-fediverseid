package fediverseid_test

import (
	"fmt"

	"github.com/reiver/go-fediverseid"
)

func ExampleFediverseID_newFediverseID() {

	const name string = "joeblow"
	const host string = "host.example"

	var fid *fediverseid.FediverseID = fediverseid.NewFediverseID(name, host)

	fmt.Printf("fediverse-id: %s", fid)

	// Output:
	// fediverse-id: @joeblow@host.example
}
