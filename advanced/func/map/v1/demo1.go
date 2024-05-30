package main

import (
	"context"
	"fmt"
)

func mapping(searchMap map[string]interface{}) {
	value := "a"
	i := searchMap[value]
	fmt.Println(i)
	searchMap["b"] = myType{name: "secondName", age: 123}
}

type myType struct {
	name string
	age  int32
}

func main() {
	m := map[string]interface{}{
		"a": myType{age: 33, name: "myname"},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "mykey", m)

	value := ctx.Value("mykey").(map[string]interface{})
	fmt.Println(value)
	fmt.Println(value["a"])
	mapping(m)
	fmt.Println(m["b"])
}
