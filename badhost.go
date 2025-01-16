package fediverseid

import (
	"strings"
)

func badHost(host string) bool {
	if "" == host {
		return true
	}

	if strings.Contains(host, "@") {
		return true
	}

	return false
}
