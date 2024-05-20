package main

import "fmt"

func containsValue(m map[string]int, target int) bool {
	for _, v := range m {
		if v == target {
			return true
		}
	}
	return false
}

func main() {
	a := map[string]int{"a": 1, "demo": 2}

	containValue := containsValue(a, 2)

	fmt.Println(containValue)
}
