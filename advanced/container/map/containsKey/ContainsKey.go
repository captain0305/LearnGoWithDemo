package main

import "fmt"

func containsKey() {
	a := map[string]int{"a": 1, "demo": 2}
	value, ok := a["demo"]
	fmt.Println(value, ok)
}

func containsKey2() {
	a := map[string]mytype{"a": {MyBool: false}, "demo": {MyBool: true}}
	value, ok := a["d"]
	fmt.Println(value.MyBool)
	fmt.Println(value, ok)
}

type mytype struct {
	MyBool bool
}

func main() {

	containsKey2()

}
