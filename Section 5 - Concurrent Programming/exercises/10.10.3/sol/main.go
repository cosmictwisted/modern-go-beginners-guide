package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go task(ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func task(ch chan string) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Task: %d", i+1)
	}
}
