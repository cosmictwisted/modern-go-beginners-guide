package main

import "fmt"

func main() {
	// Adding first 10 natural numbers
	sum := 0
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Skipping 5...")
			// Skip adding 5
			continue
		}
		if i == 8 {
			fmt.Println("8!, breaking out of loop...")
			// Break the loop when we reach 8
			break
		}
		sum += i
		fmt.Println("Sum:", sum)
	}
	fmt.Println("Sum:", sum)
}
