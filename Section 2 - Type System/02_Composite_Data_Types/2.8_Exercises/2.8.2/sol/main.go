package main

import "fmt"

func main() {
	emojis := make(map[string]string)

	emojis["Like"] = "👍"
	emojis["Dog"] = "🐶"
	emojis["Fire"] = "🔥"
	emojis["Sparkles"] = "✨"

	// check for key
	if val, ok := emojis["Dog"]; ok == true {
		fmt.Println("Key found for: Dog", val)
	} else {
		fmt.Println("Key not found")
	}

	// delete key
	delete(emojis, "Like")

	// print emoji's
	for key, val := range emojis {
		fmt.Printf("%s: %s\n", key, val)
	}
}
