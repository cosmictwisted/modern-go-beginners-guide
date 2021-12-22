package main

import "fmt"

func main() {
	// printf is used for formatting your output
	price := 13.44456
	// print float as it is
	fmt.Printf("Price as original: %f\n", price)
	// round up to 2 decimal places
	fmt.Printf("Price with 2 decimal places: %.2f\n", price)
	// print in exponent notation
	fmt.Printf("Price with exponent notation %e\n", price)

	// convert price to int and then print
	fmt.Printf("Price as integer: %d\n", int(price))

	name := "Anil Kulkarni"
	// printing a string
	fmt.Printf("Hi, %s\n", name)

	// create a custom string
	greeting := fmt.Sprintf("Hi with Sprintf, %s", name)
	fmt.Println(greeting)

	// string concatenation
	fmt.Println("Go" + "pher")

	// print just a single character
	fmt.Printf("Character: %c\n", name[0])
	fmt.Printf("Quoted Character: %q\n", name[0])
	fmt.Printf("Quoted String: %q\n", name)

	num := 200
	// print as binary
	fmt.Printf("Binary: %b\n", num)

	// print with the default style
	fmt.Printf("Default: %v\n", num)
	fmt.Printf("Default: %v\n", price)
	fmt.Printf("Default: %v\n", name)

	// Escaping characters
	fmt.Println("My best friends name is \"Robby\"")
	fmt.Println("Escaping backslash: 4\\2")
}
