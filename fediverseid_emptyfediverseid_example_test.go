package fediverseid_test

import (
	"fmt"

	"github.com/reiver/go-fediverseid"
)

func ExampleFediverseID_emptyFediverseID() {

	var fid fediverseid.FediverseID = fediverseid.EmptyFediverseID()

	fmt.Printf("fediverse-id (golang code): %#v", fid)

	// Output:
	// fediverse-id (golang code): fediverseid.EmptyFediverseID()
}
