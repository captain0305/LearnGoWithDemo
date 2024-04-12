package main

import (
	"container/list"
	"fmt"
)

func string1() {
	const World = "world"
	var hello = "hello"

	// 衔接字符串。
	var helloWorld = hello + " " + World
	helloWorld += "!"
	fmt.Println(helloWorld) // hello world!

	// 比较字符串。
	fmt.Println(hello == "hello")   // true
	fmt.Println(hello > helloWorld) // false
}

func main() {
	l := list.List{}
	l.PushFront(1)
	fmt.Println(l.Len())

	l.PushFront(2)
	fmt.Println(l.Len())

	l.Remove(l.Front())
	fmt.Println(l.Len())
	//string1()
}
