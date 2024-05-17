package main

func init() {
	println("init a")
}

func init() {
	println("init b")
}

func init() {
	println("init c")
}

// 类似java的静态代码块，在main函数之前执行
func main() {
	println("main")
}
