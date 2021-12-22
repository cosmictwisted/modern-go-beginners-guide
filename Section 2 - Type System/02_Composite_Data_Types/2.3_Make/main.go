package main

import "fmt"

func main() {
	// create a zero size slice
	var nilNums []int
	fmt.Println(nilNums)
	fmt.Printf("Capacity: %d\tLength: %d\n", cap(nilNums), len(nilNums))

	// make creates & allocates
	// make initializes to zeroth value of type
	nums := make([]int, 5, 10)
	fmt.Println(nums)
	fmt.Printf("Capacity: %d\tLength: %d\n", cap(nums), len(nums))

	// perform regular operations on a slice
	newNums := nums[1:5]
	fmt.Println(newNums)
}
