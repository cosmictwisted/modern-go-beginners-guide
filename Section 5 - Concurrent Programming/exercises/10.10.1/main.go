package main

// Debug and clear the deadlock
import "fmt"

func main() {
	ch := make(chan string)

	ch <- "hello"

	fmt.Println(<-ch)
}
