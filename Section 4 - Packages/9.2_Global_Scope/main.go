package main

import "github.com/classroomcamp/hello/hello"

func main() {
	hello.SayHello("Anil")
	// hello.cantSayHello("Anil") will throw an error as the function is not exported
	hello.Greet("Teja")
}
