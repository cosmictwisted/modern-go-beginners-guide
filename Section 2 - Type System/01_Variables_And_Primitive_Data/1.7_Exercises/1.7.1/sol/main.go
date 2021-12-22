package main

import "fmt"

func main() {
	name := "Anil Kulkarni"
	city := "Mumbai"
	var pincode int = 400077
	time := 10.20

	fmt.Printf("Hi, my name is %s and I stay in %s (%d). Its currently %.2f a.m here\n", name, city, pincode, time)
}
