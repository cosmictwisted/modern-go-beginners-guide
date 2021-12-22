package main

import (
	"fmt"
	"strings"
)

func getFormattedString(s string) (string, error) {
	if len(s) >= 5 {
		return strings.ToUpper(s), nil
	}
	return "", fmt.Errorf("Length Validation Error: length should be 5 or more")
}

func main() {
	if s, err := getFormattedString("Anil"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(s)
	}

	if s, err := getFormattedString("watermelon"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(s)
	}
}
