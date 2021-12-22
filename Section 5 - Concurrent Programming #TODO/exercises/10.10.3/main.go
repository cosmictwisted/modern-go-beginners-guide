package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go task(ch)
	// read and print the values from ch
}

func task(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Task: %d", i+1)
	}
}
