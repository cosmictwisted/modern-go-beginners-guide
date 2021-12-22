package main

import "fmt"

func greet() string {
	return "Hello, World!"
}

func noGreet() string {
	return "Sorry!"
}

func main() {
	fmt.Println(greet())
}
