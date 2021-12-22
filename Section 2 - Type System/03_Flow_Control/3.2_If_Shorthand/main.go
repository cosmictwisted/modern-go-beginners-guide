package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// setting the random seed
	rand.Seed(time.Now().UnixNano())

	// short hand with if-else chaining
	if n := rand.Intn(10); n < 5 {
		fmt.Println("n is less than 5")
	} else if n < 7 {
		fmt.Println("n is greater than 5 but less than 7")
	} else {
		fmt.Println("n is greater than 7")
	}
	// variable `n` is not accessible here

	n := 20
	num := 30
	if num < n {
		fmt.Println("num is greater than n")
		return
	}
	fmt.Println("n is greater than num")
}
