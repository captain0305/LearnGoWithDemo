package main

import "fmt"

func main() {
	hello := []byte("Hello ")
	world := "world!"

	// helloWorld := append(hello, []byte(world)...) // 正常的语法
	//必须加三个点，否则编译不通过
	helloWorld := append(hello, world...) // 语法糖
	fmt.Println(string(helloWorld))

	helloWorld2 := make([]byte, len(hello)+len(world))
	copy(helloWorld2, hello)
	// copy(helloWorld2[len(hello):], []byte(world)) // 正常的语法
	copy(helloWorld2[len(hello):], world) // 语法糖
	fmt.Println(string(helloWorld2))
}
