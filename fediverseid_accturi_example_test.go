package fediverseid_test

import (
	"fmt"

	"github.com/reiver/go-fediverseid"
)

func ExampleFediverseID_AcctURI() {

	const name string = "joeblow"
	const host string = "host.example"

	fid := fediverseid.CreateFediverseID(name, host)

	var fediverseID string = fid.String()
	var acctURI string = fid.AcctURI() // <---------

	fmt.Printf("fediverse-id: %s\n", fediverseID)
	fmt.Printf("acct-uri: %s\n", acctURI)

	// Output:
	// fediverse-id: @joeblow@host.example
	// acct-uri: acct:joeblow@host.example
}
