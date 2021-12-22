package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	go task()                          // create a fork
	time.Sleep(100 * time.Millisecond) // wait for task to finish

	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("Done...")
}

func task() {
	fmt.Println("task 1")
}
