package main

import (
	"fmt"
	"strconv"
)

/*
在 Go 语言中，string(a) 并不会将整数 a 转换为其对应的字符串表示。
相反，它会将 a 视为 Unicode 码点，并返回对应的字符。
例如，string(65) 会返回 "A"，因为 65 是字符 "A" 的 Unicode 码点
*/
func errInt2string(a int) (str string) {
	str = string(a)
	return
}

// itoa是integer to ascii的缩写
func int2string(a int) (str string) {
	str = strconv.Itoa(a)
	return
}

// atoi是ascii to integer的缩写
func string2int(str string) (a int) {
	a, _ = strconv.Atoi(str)
	return

}

func main() {
	a := 1
	b := int2string(a)
	fmt.Println(b) // 输出：1

	c := "123"
	d := string2int(c)
	fmt.Println(d) // 输出：1

}
