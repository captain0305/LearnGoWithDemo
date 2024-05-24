package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	Name string `json:"name,string"`
	age  int    `json:"age"`
}

func main() {
	var t Test
	str := `{"name":"\"test\""}`
	str2 := `{"name":"test"}`
	fmt.Println(json.Unmarshal([]byte(str2), &t))
	fmt.Println(json.Unmarshal([]byte(str), &t))
	fmt.Println(t.Name)
	return
}
