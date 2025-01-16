package fediverseid_test

import (
	"testing"

	"github.com/reiver/go-fediverseid"
)

func TestFediverseID_GoString(t *testing.T) {

	tests := []struct{
		FediverseID fediverseid.FediverseID
		Expected string
	}{
		{
			Expected: "fediverseid.EmptyFediverseID()",
		},
		{
			FediverseID: fediverseid.EmptyFediverseID(),
			Expected:   "fediverseid.EmptyFediverseID()",
		},



		{
			FediverseID: new(fediverseid.FediverseID).ChainSetName("joeblow").FediverseID(),
			Expected:   `new(fediverseid.FediverseID).ChainSetName("joeblow").FediverseID()`,
		},
		{
			FediverseID: new(fediverseid.FediverseID).ChainSetHost("host.example").FediverseID(),
			Expected:   `new(fediverseid.FediverseID).ChainSetHost("host.example").FediverseID()`,
		},



		{
			FediverseID: new(fediverseid.FediverseID).ChainSetName("JOEBLOW").FediverseID(),
			Expected:   `new(fediverseid.FediverseID).ChainSetName("JOEBLOW").FediverseID()`,
		},
		{
			FediverseID: new(fediverseid.FediverseID).ChainSetHost("HOST.EXAMPLE").FediverseID(),
			Expected:   `new(fediverseid.FediverseID).ChainSetHost("HOST.EXAMPLE").FediverseID()`,
		},



		{
			FediverseID: fediverseid.CreateFediverseID("joeblow", "host.example"),
			Expected:   `fediverseid.CreateFediverseID("joeblow", "host.example")`,
		},
		{
			FediverseID: fediverseid.CreateFediverseID("joeblow", "HOST.EXAMPLE"),
			Expected:   `fediverseid.CreateFediverseID("joeblow", "HOST.EXAMPLE")`,
		},
		{
			FediverseID: fediverseid.CreateFediverseID("JOEBLOW", "host.example"),
			Expected:   `fediverseid.CreateFediverseID("JOEBLOW", "host.example")`,
		},
		{
			FediverseID: fediverseid.CreateFediverseID("JOEBLOW", "HOST.EXAMPLE"),
			Expected:   `fediverseid.CreateFediverseID("JOEBLOW", "HOST.EXAMPLE")`,
		},
	}

	for testNumber, test := range tests {

		actual := test.FediverseID.GoString()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual go-string is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
