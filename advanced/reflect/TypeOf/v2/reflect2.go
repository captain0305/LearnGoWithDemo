package main

import (
	"fmt"
	"reflect"
)

type T []interface {
	m()
}

func (T) m() {}

func main() {
	tp := reflect.TypeOf(new(interface{}))
	tt := reflect.TypeOf(T{})
	fmt.Println(tp.Kind(), tt.Kind()) // ptr slice

	// 使用间接的方法得到表示两个接口类型的reflect.Type值。
	ti, tim := tp.Elem(), tt.Elem()
	fmt.Println(ti.Kind(), tim.Kind()) // interface interface

	fmt.Println(tt.Implements(tim))  // true
	fmt.Println(tp.Implements(tim))  // false
	fmt.Println(tim.Implements(tim)) // true

	// 所有的类型都实现了任何空接口类型。
	fmt.Println(tp.Implements(ti))  // true
	fmt.Println(tt.Implements(ti))  // true
	fmt.Println(tim.Implements(ti)) // true
	fmt.Println(ti.Implements(ti))  // true
}
