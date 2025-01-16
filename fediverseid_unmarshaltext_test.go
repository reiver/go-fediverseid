package fediverseid_test

import (
	"testing"

	"github.com/reiver/go-fediverseid"
)

func TestFediverseID_UnmarshalText(t *testing.T) {

	tests := []struct{
		Text []byte
		Expected fediverseid.FediverseID
	}{
		{
			Text:                              []byte(`@@host.example`),
			Expected: fediverseid.CreateFediverseID("", "host.example"),
		},



		{
			Text:                              []byte(`@@HOST.EXAMPLE`),
			Expected: fediverseid.CreateFediverseID("", "HOST.EXAMPLE"),
		},



		{
			Text:                           []byte(`@joeblow@host.example`),
			Expected: fediverseid.CreateFediverseID("joeblow", "host.example"),
		},
		{
			Text:                           []byte(`@joeblow@HOST.EXAMPLE`),
			Expected: fediverseid.CreateFediverseID("joeblow", "HOST.EXAMPLE"),
		},
		{
			Text:                           []byte(`@JOEBLOW@host.example`),
			Expected: fediverseid.CreateFediverseID("JOEBLOW", "host.example"),
		},
		{
			Text:                           []byte(`@JOEBLOW@HOST.EXAMPLE`),
			Expected: fediverseid.CreateFediverseID("JOEBLOW", "HOST.EXAMPLE"),
		},
	}

	for testNumber, test := range tests {
		var actual fediverseid.FediverseID
		err := actual.UnmarshalText(test.Text)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("TEXT: %q", test.Text)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual unmarshaled-text is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("TEXT: %q", test.Text)
			continue
		}
	}
}

func TestFediverseID_UnmarshalText_error(t *testing.T) {

	tests := []struct{
		Text []byte
		ExpectedError string
	}{
		{
			ExpectedError: "fediverseid: not fediverse-id",
		},
		{
			Text: []byte(nil),
			ExpectedError: "fediverseid: not fediverse-id",
		},
		{
			Text: []byte(""),
			ExpectedError: "fediverseid: not fediverse-id",
		},



		{
			Text: []byte("@joeblow"),
			ExpectedError: "fediverseid: host not found",
		},



		{
			Text: []byte("@JOEBLOW"),
			ExpectedError: "fediverseid: host not found",
		},



		{
			Text: []byte("joeblow"),
			ExpectedError: "fediverseid: not fediverse-id",
		},
	}

	for testNumber, test := range tests {

		var actualFediverseID fediverseid.FediverseID
		err := actualFediverseID.UnmarshalText(test.Text)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("TEXT: %q", test.Text)
			continue
		}

		{
			actual := actualFediverseID
			expected := fediverseid.EmptyFediverseID()

			if expected != actual {
				t.Errorf("For test #%d, the actual (unmarshaled) fediverse-id is not what was expected (to be empty).", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("TEXT: %q", test.Text)
				continue
			}
		}

		{
			actual := err.Error()
			expected := test.ExpectedError

			if expected != actual {
				t.Errorf("For test #%d, the actual error-message is not what was expected (to be empty).", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("TEXT: %q", test.Text)
				continue
			}
		}
	}
}
