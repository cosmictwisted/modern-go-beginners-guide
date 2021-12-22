package main

import (
	"fmt"
	"strings"
)

func changeCaseByRef(s *string) {
	// will change the original string
	*s = strings.Title(*s)
}

func changeCaseByValue(s string) {
	// no change to the original string
	// `s` is a copy of the original string
	s = strings.Title(s)
}

func greet() *string {
	s := "Hello Gopher..."
	return &s
}

func main() {
	s := "the big blue sky"
	// pass by value as `s` is not mutable
	changeCaseByValue(s)
	fmt.Println("With Simple:", s)

	// pass by reference
	changeCaseByRef(&s)
	fmt.Println("With Pointer:", s)

	g := greet()
	fmt.Printf("Greeting: %s\n", *g)
}
