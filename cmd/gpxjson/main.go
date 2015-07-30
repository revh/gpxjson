package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/revh/gpxjson"
)

func main() {
	f, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		panic(err)
	}

	p, err := gpxjson.JSON(f)
	fmt.Printf("%s", p)
}
