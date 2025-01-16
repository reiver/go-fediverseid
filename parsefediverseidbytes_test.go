package fediverseid_test

import (
	"testing"

	"github.com/reiver/go-fediverseid"
)

func TestParseFediverseIDBytes(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected fediverseid.FediverseID
	}{
		{
			Value: []byte("@joeblow@example.com"),
			Expected: fediverseid.CreateFediverseID("joeblow", "example.com"),
		},



		{
			Value: []byte("@joeblow@EXAMPLE.COM"),
			Expected: fediverseid.CreateFediverseID("joeblow", "EXAMPLE.COM"),
		},
		{
			Value: []byte("@JOEBLOW@example.com"),
			Expected: fediverseid.CreateFediverseID("JOEBLOW", "example.com"),
		},
		{
			Value: []byte("@JOEBLOW@EXAMPLE.COM"),
			Expected: fediverseid.CreateFediverseID("JOEBLOW", "EXAMPLE.COM"),
		},
	}

	for testNumber, test := range tests {

		actual, err := fediverseid.ParseFediverseIDBytes(test.Value)
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
