# go-fediverseid

Package **fediverseid** implements tools for working with **Fediverse-ID**s, for the Go programming language.

**Fediverse-ID**s look like these:

* `@reiver@mastodon.social`
* `@joeblow@example.com`
* `@dariush@host.example`
* `@malekeh@host.example`

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-fediverseid

[![GoDoc](https://godoc.org/github.com/reiver/go-fediverseid?status.svg)](https://godoc.org/github.com/reiver/go-fediverseid)

## Origin

**Fediverse-ID** are derived from **Twitter-IDs**, such:

* `@reiver`
* `@twitter`
* `@golanggo`

**Twitter-IDs** were **_not_** invented by the Twitter company.
Instead, users of Twitter started using them on Twitter as a way of replying to other people on Twitter.
The Twitter company noticed this, and then (the Twitter company) built this into Twitter.

Twitter users started using **Twitter-IDs** because this convention of putting an at-symbol ("@") in front of someone _handle_ to reply to them already existing on blogs, and web-based bulletin boards / forums.
Although on blogs, and web-based bulletin boards / forums, the `@handle` convention was often hyper-linked to the post that was being replied to.

Some people try to trace a history for **Fediverse-ID**s to e-mail addresses — _these people are wrong_.
E-Mail addresses are compared to **Fediverse-ID**s as a way of helping people understand what **Fediverse-ID**s are.
The origin of the **Fediverse-ID**s does **_not_** go back to e-mail addresses — it instead goes back to Twitter, blogs, and web-based bulletin boards / forums.

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
