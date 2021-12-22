package main

import "fmt"

func main() {
	// array's are fixed size data structures
	var words [2]string
	words[0] = "Gopher"
	words[1] = "Booting"
	fmt.Println(words[0], words[1])

	nums := [3]int{1, 2, 3}
	fmt.Println(nums)
	fmt.Printf("Capacity: %d\tLength: %d\n", cap(nums), len(nums))
}
