// Tour Example
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// This will throw a runtime error
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
