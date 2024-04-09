package _func

import (
	"fmt"
)

// go引入代码包时不支持循环依赖，即A包引入B包，B包不能引入A包
// 一个函数名首字母是小写，这个函数只能在同一个包内被访问，他是私有的
func abc() {
	fmt.Println("func包里的私有abc函数执行")
	s, d := SquaresOfSumAndDiff(3, 4)
	fmt.Println(s, d)

	// SquaresOfSumAndDiff2(3, 4)
}

func Abc() {
	fmt.Println("func包里的公有Abc函数执行")
	abc()
}
