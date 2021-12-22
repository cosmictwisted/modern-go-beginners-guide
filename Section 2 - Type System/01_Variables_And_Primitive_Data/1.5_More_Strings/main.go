package main

import "fmt"

func main() {
	// strings can contain unicode characters
	status := "I love the rainbow 🌈"
	fmt.Println(status)

	// strings can be indexed
	fmt.Printf("The first character of status is: %q\n", status[0])

	// strings are just read-only slice of bytes
	alphabets := "abcd"
	fmt.Printf("alphabets as code points: % d", []byte(alphabets))
	fmt.Println()

	// There is no concept of characters, instead it is only bytes or rune
	char := 'a'
	fmt.Printf("Char 'a': %v\n", char)
	fmt.Printf("Char 'a': %c\n", char)
	// rune is just an alias for type int32
	fmt.Printf("Char 'b': %c\n", char+1)

	rainbow := "🌈"
	// Even though it looks like 🌈  is a single character,
	// the length is actually 4 bytes.
	fmt.Printf("The length of 🌈 is: %d bytes\n", len(rainbow))
	fmt.Printf("🌈 as list of bytes: %v\n", []byte(rainbow))

	// Converting string to slice of rune
	fmt.Printf("The Unicode of 🌈 is: %U\n", []rune(rainbow))

	// looping over strings
	name := "अनिल कुलकर्णी"
	for index, runeValue := range name {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	fmt.Println()
}
