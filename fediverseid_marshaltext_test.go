package fediverseid_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-fediverseid"
)

func TestFediverseID_MarshalText(t *testing.T) {

	tests := []struct{
		FediverseID fediverseid.FediverseID
		Expected []byte
	}{
		{
			FediverseID: new(fediverseid.FediverseID).ChainSetHost("host.example").FediverseID(),
			Expected:   []byte(`@@host.example`),
		},



		{
			FediverseID: new(fediverseid.FediverseID).ChainSetHost("HOST.EXAMPLE").FediverseID(),
			Expected:   []byte(`@@host.example`),
		},



		{
			FediverseID: fediverseid.CreateFediverseID("joeblow", "host.example"),
			Expected:   []byte(`@joeblow@host.example`),
		},
		{
			FediverseID: fediverseid.CreateFediverseID("joeblow", "HOST.EXAMPLE"),
			Expected:   []byte(`@joeblow@host.example`),
		},
		{
			FediverseID: fediverseid.CreateFediverseID("JOEBLOW", "host.example"),
			Expected:   []byte(`@JOEBLOW@host.example`),
		},
		{
			FediverseID: fediverseid.CreateFediverseID("JOEBLOW", "HOST.EXAMPLE"),
			Expected:   []byte(`@JOEBLOW@host.example`),
		},
	}

	for testNumber, test := range tests {

		actual, err := test.FediverseID.MarshalText()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual go-string is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}

func TestFediverseID_MarshalText_error(t *testing.T) {

	tests := []struct{
		FediverseID fediverseid.FediverseID
		ExpectedError string
	}{
		{
			ExpectedError: "fediverseid: empty",
		},
		{
			FediverseID: fediverseid.EmptyFediverseID(),
			ExpectedError: "fediverseid: empty",
		},



		{
			FediverseID: new(fediverseid.FediverseID).ChainSetName("joeblow").FediverseID(),
			ExpectedError: "fediverseid: empty host",
		},



		{
			FediverseID: new(fediverseid.FediverseID).ChainSetName("JOEBLOW").FediverseID(),
			ExpectedError: "fediverseid: empty host",
		},

	}

	for testNumber, test := range tests {

		actualBytes, err := test.FediverseID.MarshalText()
		if nil == err {
			t.Errorf("For test #%d, expected an error but actually get one.", testNumber)
			continue
		}

		{
			actual := actualBytes

			if nil != actual {
				t.Errorf("For test #%d, the actual bytes is not what was expected (to be empty).", testNumber)
				t.Logf("EXPECTED: %#v", nil)
				t.Logf("ACTUAL: %#v", actual)
				continue
			}
		}

		{
			actual := err.Error()
			expected := test.ExpectedError

			if expected != actual {
				t.Errorf("For test #%d, the actual error-message is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL: %q", actual)
				continue
			}
		}
	}
}
