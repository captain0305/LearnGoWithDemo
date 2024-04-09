package main

import "fmt"

/**
 * goto语句实现循环
 */
func goto1() {
	i := 0

Next: // 跳转标签声明
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next // 跳转
	}
}

// label的层级关系
func goto2() {
	goto Label1
Label1:
	{
	Label2:
		{
			fmt.Println(111)
		}
		goto Label2
	}

	//错误
	//goto Label1 // error
	//{
	//Label1:
	//	goto Label2 // error
	//}
	//{
	//Label2:
	//}

}

func goto3() {
	i := 0
Next:
	if i >= 5 {
		// error: goto Exit jumps over declaration of k
		goto Exit
	}
	{
		k := i + i
		fmt.Println(k)
	}

	i++
	goto Next
Exit: // 此标签声明在k的作用域内，但
	// 它的使用在k的作用域之外。

}

func main() {
	goto3()
}
