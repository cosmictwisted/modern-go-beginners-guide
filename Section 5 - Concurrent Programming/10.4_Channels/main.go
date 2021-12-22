package main

import (
	"fmt"
	"time"
)

func main() {
	// declare a channel of type bool
	ch := make(chan string)

	/* Channel Directions -
	Send: ch <- data
	Receive: <-ch
	*/

	now := time.Now()
	// start a goroutine that executes a task
	go func() {
		// execute task
		task(ch)
		// close(ch) we can explicitly close channel but it is not required
		// channels not in use are garbage collected
	}()

	// block till we receive a message from done channel
	msg := <-ch // notice the direction of arrow for `receiving` from the channel

	fmt.Println(msg)
	fmt.Println("Elapsed time:", time.Since(now))
	// fmt.Println(<-ch)
}

func task(ch chan string) {
	// notice the direction of arrow for `writing` message to a channel
	ch <- fmt.Sprintln("Task 1...")
}
