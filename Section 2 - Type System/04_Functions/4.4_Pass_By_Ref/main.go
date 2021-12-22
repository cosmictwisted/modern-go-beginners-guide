package main

import "fmt"

func passByRef(nums []int) {
	// will change the underlying slice
	nums[0] = -2
	// `n` also references the same underlying slice
	// we will understand this behavior when we study `pointers`
	n := nums
	n[1] = -4
}

func passByRefNCopy(nums []int) {
	// create a new slice of equal length
	n := make([]int, len(nums))
	// copy slice
	copy(n, nums)
	n[0] = -1
	fmt.Println("Ref & Copy:", n)
}

func main() {
	n := []int{2, 4, 6, 8}
	fmt.Println("n:", n)

	// arrays and slices are passed by reference
	passByRef(n)
	fmt.Println("n:", n)

	m := []int{1, 3, 5, 7}
	fmt.Println("m:", m)
	passByRefNCopy(m)
}
