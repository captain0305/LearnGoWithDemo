package main

import "fmt"

type myType struct {
	a bool
	c string
}

func main() {
	mystruct := myType{c: "1323"}

	fmt.Println(mystruct.a)
	fmt.Println(mystruct.c)
}
