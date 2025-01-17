package fediverseid_test

import (
	"testing"

	"github.com/reiver/go-fediverseid"
)

func TestFediverseID_Unserialize(t *testing.T) {

	tests := []struct{
		Text string
		Expected fediverseid.FediverseID
	}{
		{
			Text:                                     `@@host.example`,
			Expected: fediverseid.CreateFediverseID("", "host.example"),
		},



		{
			Text:                                     `@@HOST.EXAMPLE`,
			Expected: fediverseid.CreateFediverseID("", "HOST.EXAMPLE"),
		},



		{
			Text:                                     `@joeblow@host.example`,
			Expected: fediverseid.CreateFediverseID("joeblow", "host.example"),
		},
		{
			Text:                                     `@joeblow@HOST.EXAMPLE`,
			Expected: fediverseid.CreateFediverseID("joeblow", "HOST.EXAMPLE"),
		},
		{
			Text:                                     `@JOEBLOW@host.example`,
			Expected: fediverseid.CreateFediverseID("JOEBLOW", "host.example"),
		},
		{
			Text:                                     `@JOEBLOW@HOST.EXAMPLE`,
			Expected: fediverseid.CreateFediverseID("JOEBLOW", "HOST.EXAMPLE"),
		},
	}

	for testNumber, test := range tests {
		var actual fediverseid.FediverseID
		err := actual.Unserialize(test.Text)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("TEXT: %q", test.Text)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual unserialized-text is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("TEXT: %q", test.Text)
			continue
		}
	}
}

func TestFediverseID_Unserialize_error(t *testing.T) {

	tests := []struct{
		Text string
		ExpectedError string
	}{
		{
			ExpectedError: "fediverseid: not fediverse-id",
		},
		{
			Text: "",
			ExpectedError: "fediverseid: not fediverse-id",
		},



		{
			Text: "@joeblow",
			ExpectedError: "fediverseid: host not found",
		},



		{
			Text: "@JOEBLOW",
			ExpectedError: "fediverseid: host not found",
		},



		{
			Text: "joeblow",
			ExpectedError: "fediverseid: not fediverse-id",
		},
	}

	for testNumber, test := range tests {

		var actualFediverseID fediverseid.FediverseID
		err := actualFediverseID.Unserialize(test.Text)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("TEXT: %q", test.Text)
			continue
		}

		{
			actual := actualFediverseID
			expected := fediverseid.EmptyFediverseID()

			if expected != actual {
				t.Errorf("For test #%d, the actual (unserialized) fediverse-id is not what was expected (to be empty).", testNumber)
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
