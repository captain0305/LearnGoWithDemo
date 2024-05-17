package main

import (
	"flag"
	"fmt"
)

var debugMode bool

func init() {
	flag.BoolVar(&debugMode, "debug", true, "enable debug mode")
	flag.Parse()
}

func main() {
	if debugMode {
		fmt.Println("Debug mode is on.")
	} else {
		fmt.Println("Debug mode is off.")
	}
}
