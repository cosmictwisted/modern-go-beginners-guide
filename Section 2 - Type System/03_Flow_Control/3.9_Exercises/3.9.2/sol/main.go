package main

import "fmt"

func main() {
	name := "gopher"

	for i := len(name) - 1; i >= 0; i-- {
		fmt.Printf("%c", name[i])
	}
	fmt.Println()
}
