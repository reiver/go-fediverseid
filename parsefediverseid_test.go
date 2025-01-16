package fediverseid_test

import (
	"testing"

	"github.com/reiver/go-fediverseid"
)

func TestParseFediverseID(t *testing.T) {

	tests := []struct{
		Value string
		Expected fediverseid.FediverseID
	}{
		{
			Value: "@joeblow@example.com",
			Expected: fediverseid.CreateFediverseID("joeblow", "example.com"),
		},



		{
			Value: "@joeblow@EXAMPLE.COM",
			Expected: fediverseid.CreateFediverseID("joeblow", "example.com"),
		},
		{
			Value: "@JOEBLOW@example.com",
			Expected: fediverseid.CreateFediverseID("JOEBLOW", "example.com"),
		},
		{
			Value: "@JOEBLOW@EXAMPLE.COM",
			Expected: fediverseid.CreateFediverseID("JOEBLOW", "example.com"),
		},
	}

	for testNumber, test := range tests {

		actual, err := fediverseid.ParseFediverseID(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual (parsed) fediverse-id is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}
