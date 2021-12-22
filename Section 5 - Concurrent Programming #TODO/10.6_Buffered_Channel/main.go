package main

import "fmt"

func main() {
	// create a buffered channel with capacity 3
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	// fmt.Println("Received:", <-ch)
	// fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
}
