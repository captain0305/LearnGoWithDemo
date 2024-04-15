package main

import "fmt"

// type switch断言
func switchType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type is int")
	case string:
		fmt.Println("the type is string")
	case float64:
		fmt.Println("the type is float")
	case nil:
		fmt.Println("the type is nil")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	var a int
	a = 10
	switchType(a)

	var i interface{}
	switchType(i)

	switchType(10.5)
}
