package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		// Deferred function calls are executed in Last In First Out
		defer fmt.Println(i)
	}
}
