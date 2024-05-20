package main

import "fmt"

func deffer1() {
	s := []string{"a", "demo", "c", "d"}
	defer fmt.Println(s) // [a x y d]
	// defer append(s[:1], "x", "y") // 编译错误
	defer func() {
		_ = append(s[:1], "x", "y")
	}()
}

func main() {

	s := []string{"a", "demo", "c", "d"}
	strings := append(s[:1], "x", "y")
	fmt.Println(strings) // [a x y]
	fmt.Println(s)

	//deffer1()
}
