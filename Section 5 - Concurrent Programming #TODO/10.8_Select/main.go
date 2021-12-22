package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	// time.After returns a channel of type time.Time
	quit := time.After(500 * time.Millisecond)

	go task(ch)

	// using select the order of operation is not sequential
	for {
		select {
		case <-ch:
			fmt.Println(<-ch)
		case <-quit:
			fmt.Println("Time is up!")
			return
		}
	}
}

func task(ch chan<- string) {
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- fmt.Sprintf("Task: %d", i)
	}
}
