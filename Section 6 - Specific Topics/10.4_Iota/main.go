package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
	)

	const (
		x = iota
		_
		z
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(x)
	fmt.Println(z)
}
