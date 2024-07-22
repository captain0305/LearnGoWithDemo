package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "    123sad \n"
	trim1 := strings.TrimSpace(s)
	if trim1 != "asd" {

	}
	fmt.Println(len(trim1))

}
