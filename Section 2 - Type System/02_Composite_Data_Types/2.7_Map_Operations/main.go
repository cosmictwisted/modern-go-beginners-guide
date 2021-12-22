package main

import "fmt"

func main() {
	marks := make(map[string]int)

	marks["Math"] = 60
	marks["Science"] = 65
	marks["Geography"] = 65

	// iterating over maps
	for key, value := range marks {
		fmt.Printf("Subject: %s\tMarks: %d out of 100\n", key, value)
	}

	// check for any key
	if val, ok := marks["Math"]; ok == true {
		fmt.Println("Math:", val)
	} else {
		fmt.Println("Key not found")
	}

	// getting value
	mathMarks := marks["Math"]
	fmt.Println("Math marks:", mathMarks)

	// delete key
	delete(marks, "Science")
	if _, ok := marks["Science"]; ok != true {
		fmt.Println("Key not found")
	}
}
