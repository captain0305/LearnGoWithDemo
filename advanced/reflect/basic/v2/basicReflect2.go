package main

import (
	"fmt"
	"reflect"
)

/*
2.从反射对象可以获取 interface{} 变量；
*/

func main() {
	v := reflect.ValueOf(1)
	fmt.Println(v)
	fmt.Println(reflect.TypeOf(v))

	i := v.Interface().(int)
	fmt.Println(i)
	fmt.Println(v)
}
