package main

import "fmt"

// function with single return value
func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Add:", add(1, 2))
}
