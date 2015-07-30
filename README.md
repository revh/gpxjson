GPXJSON
=======

gpxjson is command line utility and a little go library aims to convert a gpx file to a simplified version in JSON

##Getting Started

###Installing

To start using gpxjson, install Go and run `go get`:

```sh
$ go get github.com/revh/gpxjson/...
```

This will retrieve the library and install the gpxjson command line utility into your `$GOBIN` path.

###Library

You can use gpxjson package in your go program:

```go
package main

import (
    "log"

    "github.com/revh/gpxjson"
)

func main() {
  ...

  json, err := gpxjson.Convert([]byte(gpxXml))
  if err != nil {
    log.Fatal(err)
  }

  ...
}
```

###CLI

You can invoke gpxjson as a command line utility:

```sh
$ gpxjson input.gpx > output.json
```

###About

This utility was made for study purpose to improve my Go skills. Any contributions are welcome.
