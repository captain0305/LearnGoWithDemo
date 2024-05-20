package main

import (
	"LearnGoWithDemo/advanced/typeEmbedding/usage/selector/occlusionAndCollision/v4/foo"
	"fmt"
)

type B struct {
	n bool
}

func (b B) m() {
	fmt.Println("B", b.n)
}

type C struct {
	foo.A
	B
}

// m()改成M()就会出问题。为什么小写不碰撞，大写就碰撞了？
func main() {
	var c C
	c.m() // B false\
	//c.A.m()    // A false
	foo.Bar(c) // A 0
}
