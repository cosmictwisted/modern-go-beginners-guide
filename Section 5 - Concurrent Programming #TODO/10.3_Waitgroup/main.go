package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // add a task
	now := time.Now()

	go func() { // fork and launch in goroutine
		defer wg.Done()
		task()
	}()

	wg.Wait() // wait to join

	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("Done...")
}

func task() {
	fmt.Println("task 1")
}
