package main

import "fmt"

func main() {
	s := "gopher"
	// `a` is a pointer receiver which holds or points to
	// the address of `s`.
	var a *string = &s
	fmt.Printf("Value: %s\tType': %T\tAddress: %v\n", *a, a, a)
	fmt.Printf("Address of 's': %v\n", &s)

	// since `a` points to the address of `s`, the value in `s`
	// can be changed by assigning a new value to `a`.
	*a = "beaver"
	// now check the original string
	fmt.Println("s:", s)
}
