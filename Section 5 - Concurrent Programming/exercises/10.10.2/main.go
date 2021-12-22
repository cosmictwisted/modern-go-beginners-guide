// Run the program and debug from the error messages
package main

import "fmt"

func main() {
	ch := make(chan int)
	go even(ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func even(ch <-chan int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			ch <- i
		}
	}
}
