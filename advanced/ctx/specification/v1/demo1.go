package main

import (
	"context"
	"fmt"
)

func anotherFunction(ctx context.Context) {
	// Get the value from the context
	value := ctx.Value(myKey{})

	if value != nil {
		// Type assertion to ensure the value is of the expected type
		if strValue, ok := value.(string); ok {
			// Use the value
			fmt.Println("Value:", strValue)
		}
	}
}

type myKey struct{}

func myFunction(ctx context.Context) {
	// Associate the value with the custom key in the context
	key := myKey{}
	fmt.Println(&key)
	ctx = context.WithValue(ctx, key, "myValue")

	// Call other functions, passing the context with the custom key
	anotherFunction(ctx)
}

/*
逆天特性
键的比较：在 Go 语言中，当将一个值与上下文中的键进行比较时，实际上比较的是键的类型和值，而不是键本身的内存地址。因此，即使两个自定义类型的实例具有不同的内存地址，它们仍然可以被视为相同的键，因为它们具有相同的类型和值。
类型唯一性：在 Go 语言中，每个类型都是唯一的。即使您创建了两个具有相同结构的自定义类型，它们也被视为不同的类型。因此，即使两个自定义类型的实例具有相同的内存地址，它们也被视为不同的键，因为它们具有不同的类型。
*/
func main() {
	ctx := context.Background()
	fmt.Println(&myKey{})
	myFunction(ctx)

}
