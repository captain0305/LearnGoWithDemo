package main

import "fmt"

func containsKey() {
	a := map[string]int{"a": 1, "b": 2}
	value, ok := a["b"]
	fmt.Println(value, ok)
}

func main() {

	containsKey()

}
