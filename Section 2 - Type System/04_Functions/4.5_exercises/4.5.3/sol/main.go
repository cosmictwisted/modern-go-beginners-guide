package main

import "fmt"

func getSlice(nums []float64) ([]float64, bool, error) {
	isBelow10 := true
	newNums := make([]float64, len(nums))
	copy(newNums, nums)

	for i := 0; i < len(newNums); i++ {
		if newNums[i] < 0 {
			return []float64{}, false, fmt.Errorf("Validation Error: only positive numbers are allowed")
		} else if newNums[i] > 10 {
			isBelow10 = false
		}
		newNums[i] = newNums[i] * 2
	}
	return newNums, isBelow10, nil
}

func main() {
	n := []float64{1.2, 3.4, -4, 12.98, 45.0}
	if s, b, err := getSlice(n); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All below 10.0:", b)
		fmt.Println(s)
	}
	fmt.Println("n:", n)

	m := []float64{1.2, 3.4, 4.67, 12.98, 23.18}
	if s, _, err := getSlice(m); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
	fmt.Println("m:", m)
}
