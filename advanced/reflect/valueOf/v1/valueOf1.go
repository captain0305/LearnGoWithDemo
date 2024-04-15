package main

import (
	"fmt"
	"reflect"
)

/*
3.要修改反射对象，其值必须可设置；
*/

/*
调用 reflect.ValueOf 获取变量指针；
调用 reflect.Value.Elem 获取指针指向的变量；
调用 reflect.Value.SetInt 更新变量的值：
*/
func set1() {
	i := 1
	v := reflect.ValueOf(&i)
	fmt.Println(reflect.TypeOf(v))
	v.Elem().SetInt(10)
	fmt.Println(i)
}

// panic由于v的值不可设置
func set2() {
	i := 1
	v := reflect.ValueOf(i)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(&i))
	v.SetInt(10)
	fmt.Println(i)
}

func main() {
	set1()
}
