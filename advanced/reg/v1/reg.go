package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 定义一个正则表达式模式
	pattern := `^[0-9]+(\+[0-9]+)*$`

	// 编译正则表达式
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("正则表达式编译错误:", err)
		return
	}

	// 要判断的字符串
	testStr := "123"

	// 使用正则表达式匹配字符串
	match := re.MatchString(testStr)

	// 输出匹配结果
	if match {
		fmt.Printf("%s 符合正则表达式\n", testStr)
	} else {
		fmt.Printf("%s 不符合正则表达式\n", testStr)
	}
}
