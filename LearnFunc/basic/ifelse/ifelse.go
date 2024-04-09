package main

//一个文件夹下面package只能是同一个，不然会报错
//一个文件夹下面只能有一个main的入口
import (
	"fmt"
	"math/rand"
	"time"
)

func if1() {
	rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要

	if n := rand.Int(); n%2 == 0 {
		fmt.Println(n, "是一个偶数。")
	} else {
		fmt.Println(n, "是一个奇数。")
	}

}

func if2() {

	n := rand.Int() % 2 // 此n不是上面声明的n
	if n%2 == 0 {
		fmt.Println("一个偶数。")
	}

	if n%2 != 0 {
		fmt.Println("一个奇数。")
	}
	//两种写法一样
	//if ; n % 2 != 0 {
	//	fmt.Println("一个奇数。")
	//}
}

func if3() {
	if h := time.Now().Hour(); h < 12 {
		fmt.Println("现在为上午。")
	} else if h > 19 {
		fmt.Println("现在为晚上。")
	} else {
		fmt.Println("现在为下午。")
		// 左h是一个新声明的变量，右h已经在上面声明了。
		h := h
		// 刚声明的h遮掩了上面声明的h。
		_ = h
	}
	// 上面声明的两个h在此处都不可见。
}

func main() {
	if1()
	if2()
	if3()

}
