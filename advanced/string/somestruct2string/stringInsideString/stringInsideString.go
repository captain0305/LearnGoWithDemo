package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(`This is "studygolang.com" website`)
	fmt.Println("This is \"studygolang.com\" website")
	fmt.Println("This is", strconv.Quote("studygolang.com"), "website")

}
