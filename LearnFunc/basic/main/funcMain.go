package main

import (
	_func2 "LearnGoWithDemo/LearnFunc/basic/func"
	"fmt"
	"math/rand"
	_ "math/rand" //匿名引入包，只执行包中的init函数，不会执行包中的其他函数
	. "time"      //句点引入包后，可以直接使用包中的函数，而不用包名.函数名的方式调用
)

// main代码只能在main包中执行作为程序的入口
func init() {
	fmt.Println("funcMain.go init 初始化函数")

}

// 代码断行规则：运算符后换行，点号后换行，逗号后换行，括号后换行
func StatRandomNumbers(numsRands int) (int, int) {
	var a, b int
	for i := 0; i < numsRands; i++ {
		if i >= numsRands/2 {
			break
		}
		if rand.Intn(16) < 8 {
			a = a + 1
		} else {
			b++
		}
	}
	return a, b
}

func main() {
	fmt.Println(Now())
	numbers, i2 := StatRandomNumbers(10)
	fmt.Println(numbers, i2)

	s, d := _func2.SquaresOfSumAndDiff(3, 4)
	fmt.Println(s, d)

	diff2, i := _func2.SquaresOfSumAndDiff2(3, 4)
	fmt.Println(diff2, i)

	_func2.Abc()

	fmt.Println("随机数", rand.Int())
}
