package main

import (
	"fmt"
	"reflect"
)

/*
go语言三大反射法则
1.从 interface{} 变量可以反射出反射对象；
2.从反射对象可以获取 interface{} 变量；
3.要修改反射对象，其值必须可设置；
*/

func main() {
	author := "draven"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))
}
