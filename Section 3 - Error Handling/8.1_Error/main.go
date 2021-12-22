package main

import (
	"errors"
	"fmt"
	"log"
)

func getError() (string, error) {
	return "", errors.New("Error: Just another error")
}

func main() {
	s, err := getError()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(s)
}
