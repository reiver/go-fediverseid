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
	errEmptyHost            = erorr.Error("fediverseid: empty host")
	errNameNotFound         = erorr.Error("fediverseid: name not found")
	errHostNotFound         = erorr.Error("fediverseid: host not found")
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

// EmptyFediverseID returns an empty [FediverseID].
//
// For example:
//
//	fid := fediverseid.EmptyFediverseID()
func EmptyFediverseID() FediverseID {
	return FediverseID{}
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

// NameElse returns the (raw) 'name' of a Fediverse-ID if defined, else returns 'alt'.
func (receiver FediverseID) NameElse(alt string) string {
	return receiver.name.GetElse(alt)
}

// Host returns the (raw) 'host' of a Fediverse-ID.
func (receiver FediverseID) Host() (string, bool) {
	return receiver.host.Get()
}

// HostElse returns the (raw) 'host' of a Fediverse-ID if defined, else returns 'alt'.
func (receiver FediverseID) HostElse(alt string) string {
	return receiver.host.GetElse(alt)
}

// SetName sets the (raw) 'name' of a [FediverseID]..
func (receiver *FediverseID) SetName(value string) {
	receiver.name = opt.Something(value)
}

// SetHost sets the (raw) 'host' of a [FediverseID]..
func (receiver *FediverseID) SetHost(value string) {
	receiver.host = opt.Something(value)
}

// Serialize returns the (serialized) Fediverse-ID, if valid.
// Else returns an error.
//
// Serialize is similar to [String] except that it returns an error if it is invalid.
func (receiver FediverseID) Serialize() (string, error) {
	var name string
	{
		var found bool
		name, found = receiver.name.Get()
		if !found {
			var nada string
			return nada, errNameNotFound
		}
	}

	var host string
	{
		var found bool
		host, found = receiver.host.Get()
		if !found {
			var nada string
			return nada, errHostNotFound
		}
		if "" == host {
			var nada string
			return nada, errEmptyHost
		}
	}
	host = strings.ToLower(host)
	if badHost(host) {
		var nada string
		return nada, errBadHost
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	{
		p = append(p, '@')
		p = append(p, name...)
		p = append(p, '@')
		p = append(p, host...)
	}

	return string(p), nil
}

// String returns the (serialized) Fediverse-ID, if valid.
// Else returns an empty string.
//
// String also makes [FediverseID] fit the [fmt.Stringer] interface.
// (Which is used by [fmt.Errorf], [fmt.Fprint], [fmt.Fprintf], [fmt.Fprintln], [fmt.Print], [fmt.Printf], [fmt.Println], and other similar functios.)
//
// See also: [Serialize].
func (receiver FediverseID) String() string {
	str, err := receiver.Serialize()
	if nil != err {
		return ""
	}

	return str
}
