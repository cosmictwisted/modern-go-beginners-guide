package main

import "fmt"

func main() {
	name := "Anil"
	length := len(name)

	// Variation 1
	switch length {
	case 1:
	case 2:
	case 3:
	case 4:
		fmt.Println("it's a short name")
	case 7:
		fmt.Println("it's a medium length name")
	default:
		fmt.Println("it's a long name")
	}

	// Variation 2
	switch {
	case length <= 4:
		fmt.Println("it's a short name")
	case length < 7:
		fmt.Println("it's a medium length name")
	default:
		fmt.Println("it's a long name")
	}
}
