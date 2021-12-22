package main

import (
	"fmt"
)

func main() {
	// create and initialize a map
	// keys can be any type that are comparable
	contacts := make(map[string]string)
	// inserting values
	contacts["Anil"] = "anil@example.com"
	contacts["Vani"] = "vani@example.com"
	fmt.Println(contacts)
	fmt.Println("Length:", len(contacts))

	// Maps Literal
	var contacts2 = map[string]string{
		"Teja": "teja@example.com",
		"Sam":  "sam@example.com",
	}
	fmt.Println(contacts2)

	// Map with different data types
	contactNums := make(map[string][]int)
	contactNums["Anil"] = []int{99999, 22222}
	contactNums["Teja"] = []int{33333, 11111}
	fmt.Println(contactNums)
}
