package main

import (
	"fmt"
)

type APIType int

const (
	None APIType = iota
	Universal
	HTTP
	TRPC
	WUJI
	DB
)

func main() {
	apiTypeNames := []string{
		"None",
		"Universal",
		"HTTP",
		"TRPC",
		"WUJI",
		"DB",
	}

	apiType := HTTP
	fmt.Println("API type name:", apiTypeNames[apiType])
}
