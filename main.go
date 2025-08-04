package main

import "fmt"

func getNilPointer() *int {
	return nil
}

func main() {
	fmt.Println("Getting nil pointer from function...")
	ptr := getNilPointer()
	fmt.Println("About to dereference the nil pointer...")
	fmt.Println(*ptr) // This will cause a nil pointer dereference panic
}
