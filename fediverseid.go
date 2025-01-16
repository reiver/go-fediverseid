package fediverseid

import (
	"strings"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-opt"
)

const (
	errAtSignNotFound       = erorr.Error("fediverseid: at-sign not found")
	errAtSignPrefixNotFound = erorr.Error("fediverseid: at-sign not found")
	errBadHost              = erorr.Error("fediverseid: bad host")
	errEmptyFediverseID     = erorr.Error("fediverseid: empty fediverse-id")
)

// FediverseID represents a Fediverse-ID.
//
// A (serialized) Fediverse-ID looks similar to this:
//
//	@joeblow@host.example
type FediverseID struct {
	name opt.Optional[string]
	host  opt.Optional[string]
}

// CreateFediverseID creates a [FediverseID].
//
// For example:
//
//	var name string = "joeblow"
//	var host string = "host.example"
//	
//	fid := fediverseid.CreateFediverseID(name, host)
func CreateFediverseID(name string, host string) FediverseID {
	return FediverseID{
		name: opt.Something(name),
		host: opt.Something(host),
	}
}

// ParseFediverseID parses a string and (if value) returns a [FediverseID].
//
// For example:
//
//	var value string = "@joeblow@host.example"
//	
//	fediverseid.ParseFediverseID(value)
func ParseFediverseID(id string) (FediverseID, error) {
	if "" == id {
		var nada FediverseID
		return nada, errEmptyFediverseID
	}

	{
		var b0 byte = id[0]
		if '@' != b0 {
			var nada FediverseID
			return nada, errAtSignPrefixNotFound
		}

		id = id[1:]
	}

	var atindex int = strings.LastIndexByte(id, '@')
	if atindex < 0 {
		var nada FediverseID
		return nada, errAtSignNotFound
	}

	var host string = id[1+atindex:]
	if badHost(host) {
		var nada FediverseID
		return nada, errBadHost
	}
	host = strings.ToLower(host)

	var name string = id[:atindex]

	return FediverseID{
		name: opt.Something(name),
		host: opt.Something(host),
	}, nil
}

// Name returns the (raw) 'name' of a Fediverse-ID.
func (receiver FediverseID) Name() (string, bool) {
	return receiver.name.Get()
}

// Host returns the (raw) 'host' of a Fediverse-ID.
func (receiver FediverseID) Host() (string, bool) {
	return receiver.host.Get()
}

// SetName sets the (raw) 'name' of a [FediverseID]..
func (receiver *FediverseID) SetName(value string) {
	receiver.name = opt.Something(value)
}

// SetHost sets the (raw) 'host' of a [FediverseID]..
func (receiver *FediverseID) SetHost(value string) {
	receiver.host = opt.Something(value)
}

// String returns the (serialized) Fediverse-ID, if value.
// Else returns an empty string.
func (receiver FediverseID) String() string {
	var name string
	{
		var found bool
		name, found = receiver.name.Get()
		if !found {
			return ""
		}
		if "" == name {
			return ""
		}
	}

	var host string
	{
		var found bool
		host, found = receiver.host.Get()
		if !found {
			return ""
		}
		if "" == host {
			return ""
		}
	}
	host = strings.ToLower(host)
	if badHost(host) {
		return ""
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	{
		p = append(p, '@')
		p = append(p, name...)
		p = append(p, '@')
		p = append(p, host...)
	}

	return string(p)
}
