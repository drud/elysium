# Elysium #

Elysium is a GO client library for accessing the Patheon Systems [Terminus API](https://github.com/pantheon-systems/terminus).


## Usage ##

You'll need to get a machine token from Pantheon. Please see [Creating and Revoking Machine Tokens](https://pantheon.io/docs/machine-tokens/) from Pantheon documentation for instructions on how to manage your machine tokens. It is recommended you set this value as an environment variable. The following examples will assume your token has been set as an environment variable named `TERMINUS_API_TOKEN`.

You can import Elysium for use in Go by importing the following package:

```go
import "github.com/drud/elysium/pkg/elysium"
```



### Authentication

To use Elysium, you construct a new  session, then use the various services on the client to
access the Terminus API. For example:

```go
// Create a new session for your API token.
session := elysium.NewAuthSession(os.Getenv("TERMINUS_API_TOKEN"))

```

The session object is responsible for managing API sessions and requesting new session tokens as needed. To prevent requesting new sessions to often, it supports reading and writing session state to disk.

```go
sessionLocation := "/home/user/.elysium/savedsession"

// Create a new session for your API token.
session := elysium.NewAuthSession(os.Getenv("TERMINUS_API_TOKEN"))

// Make an authentication call.
err := session.Auth()
if err != nil {
    log.Fatal(err)
}

// Write session info to disk.
err := session.Write(sessionLocation)
if err != nil {
    log.Fatal(err)
}

// Read session info from disk.
err := session.Read(sessionLocation)
if err != nil {
    log.Fatal(err)
}
```
