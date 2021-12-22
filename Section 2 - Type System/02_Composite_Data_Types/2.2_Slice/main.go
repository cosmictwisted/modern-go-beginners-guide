package main

import "fmt"

func main() {
	// declare a slice of type int
	var nums []int
	fmt.Println(nums)
	fmt.Printf("Capacity: %d\tLength: %d\n", cap(nums), len(nums))
	// add element to slice
	nums = append(nums, 2)
	fmt.Println(nums)
	// check the new capacity and length
	fmt.Printf("Capacity: %d\tLength: %d\n", cap(nums), len(nums))

	// slice literal
	names := []bool{true, true, true, false}
	fmt.Println(names)

	// declaring a array
	cities := [5]string{"Mumbai", "Delhi", "Toronto", "Perth", "Oslo"}

	// create a slice and allocate cities,
	// more in `slice operations`
	cityChoices := cities[0:3]
	fmt.Println(cityChoices)

	// slice does not create new memory, instead points to the underlying array,
	// we will study more about slice in `slice internals`
	cities[1] = "Dublin"
	fmt.Println(cityChoices)
}
