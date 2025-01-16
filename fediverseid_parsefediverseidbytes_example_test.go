package fediverseid_test

import (
	"fmt"

	"github.com/reiver/go-fediverseid"
)

func ExampleFediverseID_parseFediverseIDBytes() {

	var bytes []byte = []byte("@joeblow@host.example")

	var fid fediverseid.FediverseID
	var err error

	fid, err = fediverseid.ParseFediverseIDBytes(bytes)
	if nil != err {
		fmt.Printf("ERROR: problem parsing (serialized) fediverse-id %q: %s", bytes, err)
		return
	}

	fmt.Printf("fediverse-id name: %s\n", fid.NameElse(""))
	fmt.Printf("fediverse-id host: %s\n", fid.HostElse(""))

	// Output:
	// fediverse-id name: joeblow
	// fediverse-id host: host.example
}
