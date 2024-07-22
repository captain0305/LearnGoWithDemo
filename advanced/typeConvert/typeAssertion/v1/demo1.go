package main

import (
	"errors"
	"fmt"
)

func a(checkArgs ...interface{}) error {
	if len(checkArgs) != 2 {
		fmt.Println("参数有误")
		return errors.New("入参格式有误，参数一为请求模式,参数二为")
	}
	method, _ := checkArgs[0].(string)
	path, _ := checkArgs[1].(string)

	if method != "GET" && method != "POST" {
		fmt.Println("参数有误2")
		return errors.New("method must be either GET or POST")
	}
	fmt.Println(path)
	return nil

}

func main() {
	a("123", 123)

}
