package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8, 10}
	fmt.Println("nums:", nums)
	fmt.Printf("Capacity: %d\tLength: %d\n", cap(nums), len(nums))

	// Adding element to the slice at last
	nums = append(nums, 12)
	fmt.Println("nums:", nums)
	fmt.Printf("Capacity: %d\tLength: %d\n", cap(nums), len(nums))

	moreNums := []int{14, 16, 18, 20, 22}
	// Append a slice
	nums = append(nums, moreNums...)
	fmt.Println("nums:", nums)
	fmt.Printf("Capacity: %d\tLength: %d\n", cap(nums), len(nums))

	// copy slice
	a := make([]int, len(nums))
	copy(a, nums)
	fmt.Println("a:", a)

	// delete an element from slice at the desired index
	// let's delete 12 which is at index 5 in `a`
	a = append(a[:5], a[6:]...)
	fmt.Println("a after deleting 12:", a)

	// delete the first element
	a = a[1:]
	fmt.Println("a after deleting first element:", a)

	// delete the last element
	a = a[:len(a)-1]
	fmt.Println("a after deleting last element:", a)

	// insert back 12 into slice at index 4
	// inserting takes a little more work
	// make space by appending one element of the zeroth type
	a = append(a, 0)
	// create slice with over lapping element
	copy(a[4:], a[3:])
	fmt.Println("After copying:", a)
	// replace element at the desired index
	a[4] = 12
	fmt.Println("After inserting 12:", a)

	// insert an element at the start
	a = append([]int{2}, a...)
	fmt.Println("After inserting 2:", a)

	// pop
	x, a := a[0], a[1:]
	fmt.Println("x:", x)
	fmt.Println("a:", a)
}
