package main

func main() {
	defer func() {
		if err := recover(); err != nil {
			println("recover success")
		}
	}()
	panic("panic")
}
