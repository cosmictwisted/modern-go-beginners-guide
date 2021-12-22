package main

import "fmt"

func main() {
	num := 12
	// true and true -> true
	if num > 10 && num%2 == 0 {
		fmt.Println("num is divisible by 2")
	} else if num > 10 || num%3 != 0 {
		fmt.Println("num is not divisible by 3")
	}

	num = 6
	// false and true -> false
	if num > 10 && num%2 == 0 {
		fmt.Println("num is divisible by 2")
	} else if num > 10 || num%3 == 0 {
		// false or true -> true
		fmt.Println("num is less than 10 and divisible by 3")
	}

	num = 18
	// true and true and true -> true
	if num > 10 && num%2 == 0 && num%3 == 0 {
		fmt.Println("num is divisible by 2 and 3")
	} else if num > 10 && num%2 == 0 {
		fmt.Println("num is divisible only by 2")
	} else if num > 10 || num%3 == 0 {
		fmt.Println("num is divisible only by3")
	}

	num = 13
	// false or false -> false
	if num%2 == 0 || num%3 == 0 {
		fmt.Println("This won't print")
	}
}
