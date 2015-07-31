GPXJSON
=======

gpxjson is a command line utility and a little go library which aims you to convert a gpx file to a simplified JSON version

##Getting Started

###Installing

To start using gpxjson, install Go and run `go get`:

```sh
$ go get github.com/revh/gpxjson/...
```

This will retrieve the library and install the gpxjson command line utility into your `$GOBIN` path.

###Library

You can use gpxjson package in your Go programs like this:

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

output.json
{"segments":[[{"lat":43.76319,"lon":11.149139,"time":"2015-07-25T07:17:59Z"}],[{"lat":43.76319,"lon":11.149139,"ele":95.1,"time":"2015-07-25T07:18:00Z"},{"lat":43.76319,"lon":11.149139,"ele":95.2,"time":"2015-07-25T07:18:00Z"}]]}
```

###About

This utility was made for study purpose to improve my Go skills. Any contributions are welcome.
