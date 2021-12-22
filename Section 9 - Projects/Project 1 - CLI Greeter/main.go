package main

import (
	"fmt"
	"os"
)

func main() {
	// read all the runtime arguments
	names := os.Args

	// print names
	switch len(names) {
	case 1:
		fmt.Println("Hello, Octallium!")
	case 2:
		fmt.Println("Hello,", names[1])
	default:
		printNames(names)
	}
}

func printNames(names []string) {
	for i := 1; i < len(names); i++ {
		fmt.Println("Hello,", names[i])
	}
}
