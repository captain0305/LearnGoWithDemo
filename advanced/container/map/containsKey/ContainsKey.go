package main

import "fmt"

func containsKey() {
	a := map[string]int{"a": 1, "demo": 2}
	value, ok := a["demo"]
	fmt.Println(value, ok)
}

func main() {

	containsKey()

}
