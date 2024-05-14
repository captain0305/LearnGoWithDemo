package main

import (
	"fmt"
	"strconv"
)

func float2string(a float64) (str string) {
	//prec是默认精度
	//bitsize是指float32或float64
	str = strconv.FormatFloat(a, 'f', -1, 64)
	return
}

func string2float(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func main() {
	//这个a用32进行编码转换，会有精度问题
	a := 1.12123123123123
	b := float2string(a)
	println(b) // 输出：1.1

	floatVal, err := string2float("1.23")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The float value is", floatVal)
	}
}
