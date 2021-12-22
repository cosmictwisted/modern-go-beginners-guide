package main

import "fmt"

func main() {
	name := "Anil"

	if len(name) <= 4 {
		fmt.Println("it's a short name")
	} else if len(name) < 7 {
		fmt.Println("it's a medium length name")
	} else {
		fmt.Println("it's a long name")
	}
}
