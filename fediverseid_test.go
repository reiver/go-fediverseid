package fediverseid_test

import (
	"testing"

	"github.com/reiver/go-fediverseid"
)

func TestFediverseID(t *testing.T) {

	tests := []struct{
		Name string
		Host string
		ExpectedName string
		ExpectedNameFound bool
		ExpectedHost string
		ExpectedHostFound bool
		ExpectedString string
	}{
		{

		},



		{
			Name:         "joeblow",
			ExpectedName: "joeblow",
		},
		{
			Host:         "host.example",
			ExpectedHost: "host.example",
		},



		{
			Name:            "joeblow",
			Host:                    "host.example",
			ExpectedName:    "joeblow",
			ExpectedNameFound: true,
			ExpectedHost:            "host.example",
			ExpectedHostFound: true,
			ExpectedString: "@joeblow@host.example",
		},
		{
			Name:            "joeblow",
			Host:                    "HOST.EXAMPLE",
			ExpectedName:    "joeblow",
			ExpectedNameFound: true,
			ExpectedHost:            "HOST.EXAMPLE",
			ExpectedHostFound: true,
			ExpectedString: "@joeblow@host.example",
		},
		{
			Name:            "joeblow",
			Host:                    "hOST.eXAMpLE",
			ExpectedName:    "joeblow",
			ExpectedNameFound: true,
			ExpectedHost:            "hOST.eXAMpLE",
			ExpectedHostFound: true,
			ExpectedString: "@joeblow@host.example",
		},



		{
			Name:            "JOEBLOW",
			Host:                    "host.example",
			ExpectedName:    "JOEBLOW",
			ExpectedNameFound: true,
			ExpectedHost:            "host.example",
			ExpectedHostFound: true,
			ExpectedString: "@JOEBLOW@host.example",
		},
		{
			Name:            "JOEBLOW",
			Host:                    "HOST.EXAMPLE",
			ExpectedName:    "JOEBLOW",
			ExpectedNameFound: true,
			ExpectedHost:            "HOST.EXAMPLE",
			ExpectedHostFound: true,
			ExpectedString: "@JOEBLOW@host.example",
		},
		{
			Name:            "JOEBLOW",
			Host:                    "hOST.eXAMpLE",
			ExpectedName:    "JOEBLOW",
			ExpectedNameFound: true,
			ExpectedHost:            "hOST.eXAMpLE",
			ExpectedHostFound: true,
			ExpectedString: "@JOEBLOW@host.example",
		},



		{
			Name:            "JoeBLoW",
			Host:                    "host.example",
			ExpectedName:    "JoeBLoW",
			ExpectedNameFound: true,
			ExpectedHost:            "host.example",
			ExpectedHostFound: true,
			ExpectedString: "@JoeBLoW@host.example",
		},
		{
			Name:            "JoeBLoW",
			Host:                    "HOST.EXAMPLE",
			ExpectedName:    "JoeBLoW",
			ExpectedNameFound: true,
			ExpectedHost:            "HOST.EXAMPLE",
			ExpectedHostFound: true,
			ExpectedString: "@JoeBLoW@host.example",
		},
		{
			Name:            "JoeBLoW",
			Host:                    "hOST.eXAMpLE",
			ExpectedName:    "JoeBLoW",
			ExpectedNameFound: true,
			ExpectedHost:            "hOST.eXAMpLE",
			ExpectedHostFound: true,
			ExpectedString: "@JoeBLoW@host.example",
		},



		{
			Name:            "joeblow",
			Host:                    "first.example@second.example",
			ExpectedName:    "joeblow",
			ExpectedNameFound: true,
			ExpectedHost:            "first.example@second.example",
			ExpectedHostFound: true,
			ExpectedString: "",
		},



		{
			Name:            "joeblow@first.example",
			Host:                    "second.example",
			ExpectedName:    "joeblow@first.example",
			ExpectedNameFound: true,
			ExpectedHost:            "second.example",
			ExpectedHostFound: true,
			ExpectedString: "@joeblow@first.example@second.example",
		},
	}

	for testNumber, test := range tests {

		var actualFediverseID fediverseid.FediverseID = fediverseid.CreateFediverseID(test.Name, test.Host)

		{
			actual := actualFediverseID.String()
			expected := test.ExpectedString

			if expected != actual {
				t.Errorf("For test #%d, the (serialized) 'fediverse-id' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("NAME: %q", test.Name)
				t.Logf("Host: %q", test.Host)
				continue
			}
		}

		{
			actual, _ := actualFediverseID.Name()
			expected := test.ExpectedName

			if expected != actual {
				t.Errorf("For test #%d, the fediverse-id 'name' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("NAME: %q", test.Name)
				t.Logf("Host: %q", test.Host)
				continue
			}
		}

		{
			actual, _ := actualFediverseID.Host()
			expected := test.ExpectedHost

			if expected != actual {
				t.Errorf("For test #%d, the fediverse-id 'host' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("NAME: %q", test.Name)
				t.Logf("Host: %q", test.Host)
				continue
			}
		}
	}
}
