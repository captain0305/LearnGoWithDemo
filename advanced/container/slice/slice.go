package main

import "fmt"

func slice() {
	// 切片类型
	var a []bool
	var b []int64
	var c []map[int]bool // 元素类型为一个映射类型：map[int]bool
	var d []*int         // 元素类型为一个指针类型：*int
	_ = a
	_ = b
	_ = c
	_ = d
}

// 切片基本操作
func slice2() {
	// 下面这些切片字面量都是等价的。
	_ = []string{"break", "continue", "fallthrough"}
	_ = []string{0: "break", 1: "continue", 2: "fallthrough"}
	var b = []string{2: "fallthrough", 1: "continue", 0: "break"}
	var a = []string{2: "fallthrough", 0: "break", "continue"}

	fmt.Println(cap(a), len(a), a, a[0], a[1], a[2])
	fmt.Println(cap(b), len(b), b, b[0], b[1], b[2])
	a = append(a, "d")

	fmt.Println(cap(a), len(a), a, a[0], a[1], a[2], a[3])
	a = a[:len(a)-1]
	fmt.Println(cap(a), len(a), a, a[0], a[1], a[2])
}

func slice3() {
	var a = []int{1, 2, 3, 4, 5}

	fmt.Println(a[0:3]) // [1 2 3]
	ints := a[1:2:4]
	//子切片在原切片上操作在max的范围内，如果超过的话，复制到新的切片上再append，原切片不操作
	i := append(ints, 100, 2, 3)
	fmt.Println(a)
	fmt.Println(i) // [2 3]
}

func sliceToArray() {
	var s = []int{0, 1, 2, 3}
	var a = [3]int(s[1:])
	s[2] = 9
	fmt.Println(s) // [0 1 9 3]
	fmt.Println(a) // [1 2 3]

	//数组长度大雨被转化切片的长度
	_ = [3]int(s[:2]) // panic
}

func copySliceAndArray() {
	type Ta []int
	type Tb []int
	dest := Ta{1, 2, 3}
	src := Tb{5, 6, 7, 8, 9}
	//src复制到dest里
	n := copy(dest, src)
	fmt.Println(n, dest) // 3 [5 6 7]

	n = copy(dest[1:], dest)
	fmt.Println(n, dest) // 2 [5 5 6]

	a := [4]int{} // 一个数组
	n = copy(a[:], src)
	fmt.Println(n, a) // 4 [5 6 7 8]

	n = copy(a[:], a[2:])
	fmt.Println(n, a) // 2 [7 8 7 8]
}

func main() {

}
