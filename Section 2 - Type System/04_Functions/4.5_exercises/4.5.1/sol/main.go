package main

import "fmt"

func factorial(n int) int {
	fact := 1
	if n == 0 {
		return 1
	} else {
		// simple implementation of factorial
		for i := 1; i <= n; i++ {
			fact *= i
		}
		return fact
	}
}

func main() {
	fmt.Println(factorial(5))
}
