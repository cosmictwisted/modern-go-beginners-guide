package main

import "fmt"

// function with multiple return values
func smartDivision(numerator int, denominator int) (int, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("Zero Division Error: received 0 as denominator.")
	}
	return numerator / denominator, nil
}

func main() {
	n, err := smartDivision(4, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Division:", n)
	}

	if n, err := smartDivision(3, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Division:", n)
	}

	// ignoring a return value
	if _, err := smartDivision(3, 2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Division works fine!")
	}
}
