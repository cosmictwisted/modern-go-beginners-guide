package main

import "fmt"

func square(n []int) {
	for i := 0; i < len(n); i++ {
		n[i] = n[i] * n[i]
	}
}

func main() {
	nums := []int{1, 3, 5, 7, 9}
	fmt.Println("nums:", nums)

	// since nums is mutable it is passed by ref
	square(nums)
	fmt.Println("nums:", nums)
}
