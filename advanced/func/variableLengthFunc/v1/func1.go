package main

import "fmt"

// Sum返回所有输入实参的和。
func Sum(values ...int64) (sum int64) {
	// values的类型为[]int64。
	sum = 0
	for _, v := range values {
		sum += v
	}
	return
}

// Concat是一个低效的字符串拼接函数。
// 变长函数声明,最后一个参数的类型前缀为...，表示接受的参数为变长
func Concat(sep string, tokens ...string) string {
	// tokens的类型为[]string。
	r := ""
	for i, t := range tokens {
		if i != 0 {
			r += sep
		}
		r += t
	}
	return r
}

func main() {
	a := "12313"
	b := []string{"132", "465", "798"}

	Concat(a, "123", "456", "789")
	fmt.Println(Concat(a, b...))
}
