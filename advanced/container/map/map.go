package main

import "fmt"

func map1() {
	// 映射类型
	var a map[string]int
	var b map[int]bool
	var c map[int16][6]string       // 元素类型为一个数组类型：[6]string
	var d map[bool][]string         // 元素类型为一个切片类型：[]string
	var e map[struct{ x int }]*int8 // 元素类型为一个指针类型：*int8；
	// 键值类型为一个结构体类型。

	_ = a
	_ = b
	_ = c
	_ = d
	_ = e
}

func map2() {
	var a uint = 1
	var _ = map[uint]int{a: 123} // 没问题
	var _ = []int{0: 100}        // error: 下标必须为常量
	var b = [5]int{1: 100}       // error: 下标必须为常量

	fmt.Println(b[0], b[1])
}

type User struct {
	Name string
	Age  int
}

func map3() {
	m := map[string]int{"Go": 2007}
	m2 := map[User]int{{"Alice", 12}: 1}
	fmt.Println(m2)
	m["C"] = 1972     // 添加
	m["Java"] = 1995  // 添加
	fmt.Println(m)    // map[C:1972 Go:2007 Java:1995]
	m["Go"] = 2009    // 修改
	delete(m, "Java") // 删除
	fmt.Println(m)    // map[C:1972 Go:2009]

}
func main() {
	map3()
}
