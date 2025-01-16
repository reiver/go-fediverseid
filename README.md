# go-fediverseid

Package **fediverseid** implements tools for working with **Fediverse-ID**s, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-fediverseid

[![GoDoc](https://godoc.org/github.com/reiver/go-fediverseid?status.svg)](https://godoc.org/github.com/reiver/go-fediverseid)

## Examples

To parse a Fediverse-ID and split it into its **name** and **host** you can do something similar to the following:

```golang
import "github.com/reiver/go-fediverseid"

// ...

fediverseID, err := fediverseid.ParseFediverseID("@joeblow@host.example")
if nil != err {
	fmt.Printf("ERROR: problem parsing fediverse-id: %s\n", err)
	return
}

name, found := fediverseID.Name()
if !found {
	fmt.Println("ERROR: missing name")
	return
}

host, found := fediverseID.Host()
if !found {
	fmt.Println("ERROR: missing host")
	return
}
```

And, to generate a Fediverse-ID from a **name** and a **host** you can do something similar to the following:

```golang
import "github.com/reiver/go-fediverseid"

// ...

fid := fediverseid.CreateFediverseID("joeblow", "host.example")

var serializedFediverseID string = fediverseID.String()
```

## Import

To import package **fediverseid** use `import` code like the following:
```
import "github.com/reiver/go-fediverseid"
```

## Installation

To install package **fediverseid** do the following:
```
GOPROXY=direct go get github.com/reiver/go-fediverseid
```

## Author

Package **fediverseid** was written by [Charles Iliya Krempeaux](http://reiver.link)
