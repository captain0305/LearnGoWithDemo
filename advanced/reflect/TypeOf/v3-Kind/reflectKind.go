package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)

}
func kind1() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

}

func kind2() {
	createQuery(1)
}

func main() {
	kind1()
	kind2()
	createQuery(map[string]int{"key": 1})
}
