package main

import (
	"fmt"
	"reflect"
)

func useReflectKind() {
	reflectKind(1)
	reflectKind(0.01)
	reflectKind([]int{1, 2, 3, 4})
	reflectKind(map[string]struct{}{})
}

func reflectKind(v interface{}) {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Int:
		fmt.Printf("int value is: %d\n", value.Int())
	case reflect.Float64:
		fmt.Printf("float64 value is: %f\n", value.Float())
	case reflect.Slice:
		fmt.Printf("slice value is: %v1\n", value.Slice(0, value.Len()))
	default:
		fmt.Printf("defaule type is: %v1\n", value)
	}
}

func main() {
	useReflectKind()
}
