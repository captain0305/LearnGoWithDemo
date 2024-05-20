package main

import "net/http"

func main() {
	type P = *bool
	type M = map[int]int
	var x struct {
		string // 一个具名非指针类型
		error  // 一个具名接口类型
		*int   // 一个无名指针类型
		P      // 一个无名指针类型的别名
		M      // 一个无名类型的别名

		http.Header // 一个具名映射类型
	}
	x.string = "Go"
	x.error = nil
	x.int = new(int)
	x.P = new(bool)
	x.M = make(M)
	x.Header = http.Header{}
}
