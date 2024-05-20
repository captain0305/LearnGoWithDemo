package main

import "fmt"

//func Add[T any](a, demo T) T {
//	return a + demo
//}

func printStringArray(arr interface{}) {
	for _, a := range arr.([]string) {
		fmt.Println(a)
	}
}

func printArray[T string | int | int64](arr []T) {
	for _, a := range arr {
		fmt.Println(a)
	}
}

// any为所有的基本类型 等价于interface{} 空接口
func printArray1[T any](arr []T) {
	for _, a := range arr {
		fmt.Println(a)
	}
}

// comparable为所有的内置可比较类型
func printArray2[T comparable](arr []T) {
	for _, a := range arr {
		fmt.Println(a)
	}
}
func main() {
	strs := []string{"a", "demo", "c", "d"}
	ints := []int{1, 2, 3, 4}

	printArray(strs)

	printArray(ints)

	printArray2(strs)
}
