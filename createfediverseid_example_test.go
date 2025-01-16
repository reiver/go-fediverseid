package fediverseid_test

import (
	"fmt"

	"github.com/reiver/go-fediverseid"
)

func ExampleCreateFediverseID() {

	const name string = "joeblow"
	const host string = "host.example"

	fid := fediverseid.CreateFediverseID(name, host)

	var fediverseID string = fid.String()

	fmt.Printf("fediverse-id: %s", fediverseID)

	// Output:
	// fediverse-id: @joeblow@host.example
}
