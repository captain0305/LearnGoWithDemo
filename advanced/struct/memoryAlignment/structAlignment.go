package main

import (
	"fmt"
	"unsafe"
)

type demo1 struct {
	a int8
	b int16
	c int32
}

type demo2 struct {
	a int8
	c int32
	b int16
}

type demo3 struct {
	a int8
	b int64
	c int32
	d int16
}

type demo4 struct {
	c int8
	d int16
	a int32
	b int64
}

func main() {
	fmt.Println(unsafe.Sizeof(demo1{})) // 8
	fmt.Println(unsafe.Sizeof(demo2{})) // 12

	fmt.Println(unsafe.Sizeof(demo3{})) // 24
	fmt.Println(unsafe.Sizeof(demo4{})) // 16
}
