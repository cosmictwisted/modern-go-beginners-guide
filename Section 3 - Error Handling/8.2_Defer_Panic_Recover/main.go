package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Beginning panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovering from error:", err)
		}
	}()
	panic("Boom 💣")
}
