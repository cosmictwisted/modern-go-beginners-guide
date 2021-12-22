package main

import "fmt"

func main() {
	go task() // creates a fork
	fmt.Println("Done without task 1...")
	// program exits without waiting for task() to complete
}

func task() {
	fmt.Println("task 1")
}
