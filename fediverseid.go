package fediverseid

import (
	"encoding"
	"fmt"
	"strings"
	"unsafe"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-opt"
)

const (
	errBadHost              = erorr.Error("fediverseid: bad host")
	errEmpty                = erorr.Error("fediverseid: empty")
	errEmptyHost            = erorr.Error("fediverseid: empty host")
	errHostNotFound         = erorr.Error("fediverseid: host not found")
	errNilReceiver          = erorr.Error("fediverseid: nil receiver")
	errNotFediverseID       = erorr.Error("fediverseid: not fediverse-id")
)

// FediverseID represents a Fediverse-ID.
//
// A (serialized) Fediverse-ID looks similar to this:
//
//	@joeblow@host.example
//
// To create a FediverseID use [CreateFediverseID] or [EmptyFediverseID].
//
// To create a pointer to a FediverseID use [NewFediverseID] or new(fediverseid.FediverseID).
//
// To create a FediverseID from a serialized fediverse-id use [ParseFediverseIDBytes] or [ParseFediverseIDString].
//
//To serialize a FediverseID (to a string) use [FediverseID.MarshalText] or [FediverseID.String].
type FediverseID struct {
	name opt.Optional[string]
	host  opt.Optional[string]
}

var _ fmt.Stringer = FediverseID{}
var _ fmt.GoStringer = FediverseID{}
var _ encoding.TextMarshaler = FediverseID{}
var _ encoding.TextUnmarshaler = new(FediverseID)

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

// NewFediverseID returns a new *[FediverseID] with the `name` and `host` specified.
//
// For example:
//
//	var name string = "joeblow"
//	var host string = "host.example"
//	
//	fid := fediverseid.NewFediverseID(name, host)
func NewFediverseID(name string, host string) *FediverseID {
	fid := CreateFediverseID(name, host)
	return &fid
}

// EmptyFediverseID returns an empty [FediverseID].
//
// For example:
//
//	fid := fediverseid.EmptyFediverseID()
func EmptyFediverseID() FediverseID {
	return FediverseID{}
}

// ParseFediverseIDBytes parses a []byte and (if valid) returns a [FediverseID].
// If not valid, returns an error.
//
// For example:
//
//	var value []byte = []byte("@joeblow@host.example")
//	
//	fid, err := fediverseid.ParseFediverseIDBytes(value)
//
// See also: [ParseFediverseIDString]
func ParseFediverseIDBytes(id []byte) (FediverseID, error) {
	if len(id) <= 0 {
		var nada FediverseID
		return nada, errNotFediverseID
	}

	var str string = unsafe.String(unsafe.SliceData(id), len(id))

	return ParseFediverseIDString(str)
}

// ParseFediverseIDString parses a string and (if valid) returns a [FediverseID].
// If not valid, returns an error.
//
// For example:
//
//	var value string = "@joeblow@host.example"
//	
//	fid, err := fediverseid.ParseFediverseIDString(value)
//
// See also: [ParseFediverseIDBytes]
func ParseFediverseIDString(id string) (FediverseID, error) {
	if "" == id {
		var nada FediverseID
		return nada, errNotFediverseID
	}

	{
		var b0 byte = id[0]
		if '@' != b0 {
			var nada FediverseID
			return nada, errNotFediverseID
		}

		id = id[1:]
	}

	var atindex int = strings.LastIndexByte(id, '@')
	if atindex < 0 {
		var nada FediverseID
		return nada, errHostNotFound
	}

	var host string = id[1+atindex:]
	if badHost(host) {
		var nada FediverseID
		return nada, errBadHost
	}

	var name string = id[:atindex]

	return FediverseID{
		name: opt.Something(name),
		host: opt.Something(host),
	}, nil
}

// ChainSetName sets the (raw) 'name' of the [FediverseID], and returns the receiver.
//
// This is useful for chaining.
func (receiver *FediverseID) ChainSetName(value string) *FediverseID {
	receiver.SetName(value)
	return receiver
}

// ChainSetHost sets the (raw) 'host' of the [FediverseID], and returns the receiver.
//
// This is useful for chaining.
func (receiver *FediverseID) ChainSetHost(value string) *FediverseID {
	receiver.SetHost(value)
	return receiver
}

// FediverseID turns a  *[FediverseID] (i.e., a pointer to a [FediverseID]) back into a [FediverseID].
//
// For example:
//
//	var fid fediveseid.FediverseID = fediveseid.EmptyFediverseID().ChainSetHost("example.com").FediverseID()
func (receiver *FediverseID) FediverseID() FediverseID {
	if nil == receiver {
		return EmptyFediverseID()
	}
	return *receiver
}

// GoString returns Go code (as a string) that could be used to create this [FediverseID].
//
// GoString also makes [FediverseID] fit the [fmt.GoStringer] interface.
// (Which is used by [fmt.Errorf], [fmt.Fprint], [fmt.Fprintf], [fmt.Fprintln], [fmt.Print], [fmt.Printf], [fmt.Println], and other similar functions, with the "%#v" format.)
func (receiver FediverseID) GoString() string {
	name, nameFound := receiver.Name()
	host, hostFound := receiver.Host()

	switch {
	case !nameFound && !hostFound:
		return "fediverseid.EmptyFediverseID()"
	case !nameFound &&  hostFound:
		return fmt.Sprintf("new(fediverseid.FediverseID).ChainSetHost(%q).FediverseID()", host)
	case  nameFound && !hostFound:
		return fmt.Sprintf("new(fediverseid.FediverseID).ChainSetName(%q).FediverseID()", name)
	default: // case nameFound && hostFound:
		return fmt.Sprintf("fediverseid.CreateFediverseID(%q, %q)", name, host)
	}
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

// SetName sets the (raw) 'name' of the [FediverseID].
func (receiver *FediverseID) SetName(value string) {
	if nil == receiver {
		return
	}
	receiver.name = opt.Something(value)
}

// SetHost sets the (raw) 'host' of the [FediverseID].
func (receiver *FediverseID) SetHost(value string) {
	if nil == receiver {
		return
	}
	receiver.host = opt.Something(value)
}

// Serialize returns the (serialized) Fediverse-ID, if valid.
// Else returns an error.
//
// Serialize is similar to [FediverseID.String] except that it returns an error if it is invalid.
//
// Serialize is also similar to [FediverseID.MarshalText] except that is returns a string rather than a []byte.
func (receiver FediverseID) Serialize() (string, error) {
	bytes, err := receiver.MarshalText()
	if nil != err {
		var nada string
		return nada, err
	}

	return string(bytes), nil
}

// String returns the (serialized) Fediverse-ID, if valid.
// Else returns an empty string.
//
// String also makes [FediverseID] fit the [fmt.Stringer] interface.
// (Which is used by [fmt.Errorf], [fmt.Fprint], [fmt.Fprintf], [fmt.Fprintln], [fmt.Print], [fmt.Printf], [fmt.Println], and other similar functions.)
//
// See also: [FediverseID.Serialize].
func (receiver FediverseID) String() string {
	str, err := receiver.Serialize()
	if nil != err {
		return ""
	}

	return str
}

// MarshalText returns the (serialized) Fediverse-ID, if valid.
// Else returns an error.
//
// MarshalText is similar to [FediverseID.Serialize] except that is returns a []byte rather than a string.
//
// MarshalText is also similar to [FediverseID.String] except that it returns a []byte and an error if it is invalid.
//
// MarshalText also makes [FediverseID] fit the [encoding.TextMarshaler] interface.
// And thus, among other things, is an alternative to [json.Marshaler].
func (receiver FediverseID) MarshalText() ([]byte, error) {
	if EmptyFediverseID() == receiver {
		var nada []byte
		return nada, errEmpty
	}

	var name string = receiver.name.GetElse("")

	var host string
	{
		var found bool
		host, found = receiver.host.Get()
		if !found {
			var nada []byte
			return nada, errEmptyHost
		}
		if "" == host {
			var nada []byte
			return nada, errEmptyHost
		}
	}
	host = strings.ToLower(host)
	if badHost(host) {
		var nada []byte
		return nada, errBadHost
	}

	var buffer [128]byte
	var p []byte = buffer[0:0]

	{
		p = append(p, '@')
		p = append(p, name...)
		p = append(p, '@')
		p = append(p, host...)
	}

	return p, nil
}

// UnmarshalText unserializes Fediverse-ID, if valid, and sets the value of the receiver to it.
// Else returns an error.
//
// UnmarshalText is similar to [ParseFediverseIDBytes], except UnmarshalText is a method of [FediverseID], where [ParseFediverseIDBytes] is a standalone function.
// And because UnmarshalText is similar to [ParseFediverseIDBytes], it is also similar to [ParseFediverseIDString].
//
// UnmarshalText also makes [FediverseID] fit the [encoding.TextUnmarshaler] interface.
// And thus, among other things, is an alternative to [json.Unmarshaler].
func (receiver *FediverseID) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	fid, err := ParseFediverseIDBytes(text)
	if nil != err {
		return err
	}

	*receiver = fid
	return nil
}
