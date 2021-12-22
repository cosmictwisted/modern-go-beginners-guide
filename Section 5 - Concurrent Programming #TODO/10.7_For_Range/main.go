package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		task(ch)
		close(ch) // we explicitly close channel to intimate the receiver
	}()

	// for-range over a channel is also a blocking operation
	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println("All task done!")
}

func task(ch chan<- string) { // notice the function argument, we can specify a write only channel
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Task: %d", i+1)
	}
}
