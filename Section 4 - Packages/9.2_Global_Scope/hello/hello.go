package hello

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello,", name)
}

// cantSayHello is private, notice the lowercase first alphabet
// thus, the scope of this function is only local.
// Other files in package hello can still access it
func cantSayHello(name string) {
	fmt.Println("No Hello,", name)
}
