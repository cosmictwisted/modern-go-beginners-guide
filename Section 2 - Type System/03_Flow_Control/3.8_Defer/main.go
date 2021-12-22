package main

import "fmt"

func main() {
	// defer waits execution until other functions in
	// the same block return
	defer fmt.Println("Three")

	fmt.Println("One")
	fmt.Println("Two")
}
