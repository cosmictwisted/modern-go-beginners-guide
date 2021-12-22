package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10)

	// switch without a condition
	// switch evaluates from top to bottom
	switch {
	case r < 3:
		fmt.Println("r is less than 3")
	case r < 5:
		fmt.Println("r is greater than 3 but less than 5")
	case r < 7:
		fmt.Println("r is greater than 5 but less than 7")
	default:
		fmt.Println("r is greater than 7")
	}
}
