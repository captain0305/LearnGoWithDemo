package main

import (
	"fmt"
	"strings"
)

func main() {
	s := " 123sad \n\naa$\n     \n"
	for len(s) > 0 && (s[len(s)-1] == ' ' || s[len(s)-1] == '\n') {
		s = strings.TrimRight(s, " \n") // 去除末尾的空格或换行符
	}
	res := []int{}
	tmp := []int{1, 2, 3}
	res = append(res, tmp...)
	tmp = []int{4, 5, 6}
	res = append(res, tmp...)
	for _, re := range res {
		fmt.Println(re)
	}
	fmt.Println()
	fmt.Println(s)
	fmt.Println(len(s))

}
