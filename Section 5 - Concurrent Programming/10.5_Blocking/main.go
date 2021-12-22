package main

import "fmt"

func main() {
	ch := make(chan int)
	go task(ch)

	// receiving from channel is a blocking operation
	<-ch
	<-ch
	<-ch
	fmt.Println("All task done!")
}

func task(ch chan int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Completed Task %d...\n", i)
		ch <- i
	}
}
