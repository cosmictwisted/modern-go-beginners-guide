package main

import (
	"log"
)

func zeroDivision(a int, b int) int {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	// panic here
	return a / b
}

func main() {
	zeroDivision(3, 0)
}
