package main

import (
	"fmt"
	"testing"
)

var s string
var x = []byte{1023: 'x'}
var y = []byte{1023: 'y'}
var z = []byte{12: 'z'}

func fc() {
	// 下面的四个转换都不需要深复制。
	if string(x) != string(z) {
		s = (" " + string(x) + string(z))[1:]
		fmt.Println(s)
	}
}

func fd() {
	// 两个在比较表达式中的转换不需要深复制，
	// 但两个字符串衔接中的转换仍需要深复制。
	// 请注意此字符串衔接和fc中的衔接的差别。
	if string(x) != string(y) {
		s = string(x) + string(y)
		fmt.Println(s)
	}
}

func main() {
	fmt.Println(x)
	fmt.Println(y)
	//func AllocsPerRun(runs int, f func()) (avg float64)  runs 参数表示要运行的次数，f 是要测试的函数。函数返回的 avg 是每次运行 f 时的平均内存分配次数。
	fmt.Println(testing.AllocsPerRun(1, fc)) // 1
	fmt.Println(testing.AllocsPerRun(1, fd)) // 3
}
