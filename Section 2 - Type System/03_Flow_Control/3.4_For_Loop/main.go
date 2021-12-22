package main

import "fmt"

func main() {
	// Go provides only `for` syntax in loops

	fmt.Println("For Loop Variation 1:")
	// display first 5 numbers
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\n\nFor Loop Variation 2:")
	i := 10
	for ; i < 15; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\n\nFor Loop Variation 3:")
	i = 20
	for i < 25 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	// for-range syntax
	nums := []int{1, 2, 3, 4, 5}
	for i, v := range nums {
		fmt.Printf("Index:%d\tValue:%d\n", i, v)
	}
}
