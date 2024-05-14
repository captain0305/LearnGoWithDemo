package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// Store key-value pairs in the sync.Map
	m.Store(123, 2)
	m.Store("key1", "value1")
	m.Store("key2", "value2")

	// Load a value from the sync.Map
	value, ok := m.Load("key1")
	if ok {
		fmt.Println("key1:", value)
	}

	// Delete a key-value pair from the sync.Map
	m.Delete("key1")

	// Range over the sync.Map
	//遍历sync.Map
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	targetValue := "value2"
	found := false

	//查询某个value是否存在
	m.Range(func(key, value interface{}) bool {
		if value == targetValue {
			found = true
			// Break the loop
			return false
		}
		// Continue the loop
		return true
	})

	fmt.Println("Value2 found:", found)
}
