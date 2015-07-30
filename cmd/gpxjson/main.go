package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/revh/gpxjson"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "You must provide a gpx file in input\n")
		os.Exit(1)
	}

	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	j, err := gpxjson.Convert(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%s", j)
}
