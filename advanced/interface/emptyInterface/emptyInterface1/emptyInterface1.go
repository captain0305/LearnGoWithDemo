package main

import "fmt"

func emptyInterface() {
	var data map[string]interface{}
	//空接口可以接受任意类型参数
	data = make(map[string]interface{})
	data["name"] = "John Doe"
	data["age"] = 30
	data["isMarried"] = true
	fmt.Println(data)

	//空接口可以接受任意类型的参数
	var slice []interface{}
	//var slice2 []int
	slice = append(slice, "Hello", 123, true)
	fmt.Println(slice)
	//slice2 = append(slice2, 1, 2, true)

}

func emptyInterface2() {
	var i interface{}
	i = []int{1, 2, 3}
	fmt.Println(i) // [1 2 3]
	i = map[string]int{"Go": 2012}
	fmt.Println(i) // map[Go:2012]
	i = true
	fmt.Println(i) // true
	i = 1
	fmt.Println(i) // 1
	i = "abc"
	fmt.Println(i) // abc

	// 将接口值i中包裹的值清除掉。
	i = nil
	fmt.Println(i) // <nil>
}

// 类型断言
func doSomething(v interface{}) {
	switch u := v.(type) {
	case int:
		fmt.Println("It's an int!", u)
	case string:
		fmt.Println("It's a string!", u)
	default:
		fmt.Println("I don't know what this is.")
	}
}

func main() {

	doSomething(10)
	doSomething("hello")
	doSomething(true)

	emptyInterface2()
}
