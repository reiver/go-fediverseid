package fediverseid_test

import (
	"fmt"

	"github.com/reiver/go-fediverseid"
)

func ExampleFediverseID_AcctURI() {

	const name string = "joeblow"
	const host string = "host.example"

	fid := fediverseid.CreateFediverseID(name, host)

	var acctURI string = fid.AcctURI()

	fmt.Printf("acct-uri: %s", acctURI)

	// Output:
	// acct-uri: acct:joeblow@host.example
}
