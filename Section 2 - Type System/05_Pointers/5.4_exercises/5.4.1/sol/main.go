package main

import "fmt"

func multiplyOrDivide(n *int) {
	if *n < 10 {
		*n *= 2
	} else {
		*n /= 2
	}
}

func main() {
	n := 8
	multiplyOrDivide(&n)
	fmt.Println("n:", n)

	n = 12
	multiplyOrDivide(&n)
	fmt.Println("n:", n)
}
