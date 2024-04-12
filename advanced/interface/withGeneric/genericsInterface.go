package main

import "fmt"

// 波浪号代表支持衍生类型
type MyInt interface {
	int | ~int8 | int16 | int32 | int64
}

type Int8AAA int8

func GetMaxNum[T MyInt](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a int = 1
	var b int = 2
	var c int8 = 3
	var d int8 = 4

	var n Int8AAA = 5
	var m Int8AAA = 6
	fmt.Println(GetMaxNum(c, d))
	fmt.Println(GetMaxNum(m, n))
	fmt.Println(GetMaxNum[int](a, b))
}
