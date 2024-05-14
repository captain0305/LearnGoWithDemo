package main

import (
	"fmt"
	"strconv"
)

func bool2string(bool2 bool) (str string) {
	str = strconv.FormatBool(bool2)
	return

}

func string2bool(str string) (bool2 bool) {
	bool2, _ = strconv.ParseBool(str)
	return
}

func main() {

	boolstr := bool2string(true)
	bool2 := string2bool("true")
	fmt.Println(bool2)

	//fmt.Println("found some value?"+true) // 编译错误：cannot convert true (type bool) to type string
	fmt.Println("found some value:" + boolstr) // 输出：true
}
