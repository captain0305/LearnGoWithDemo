package main

import "fmt"

func make1() {
	// 创建映射。
	fmt.Println(make(map[string]int)) // map[]
	//映射
	m := make(map[string]int, 3)
	fmt.Println(m, len(m)) // map[] 0
	m["C"] = 1972
	m["Go"] = 2009
	fmt.Println(m, len(m)) // map[C:1972 Go:2009] 2

	// 创建切片。
	s := make([]int, 3, 5)
	s[0] = 2
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 5
	s = append(s, 1)

	fmt.Println(s, len(s), cap(s)) // [0 0 0 1] 4 5
	s = make([]int, 2)
	fmt.Println(s, len(s), cap(s)) // [0 0] 2 2
}

func main() {
	make1()
}
