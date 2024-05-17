package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(recover()) // 3
	}()

	defer panic(3) // 将替换恐慌2
	defer panic(2) // 将替换恐慌1
	defer panic(1) // 将替换恐慌0
	panic(0)
	fmt.Println("end")
}
