package main

import "fmt"

// create a `User` struct
type User struct {
	username string
	email    string
	age      int
}

func main() {
	// declare a variable of type `User`
	var user1 User
	user1 = User{
		username: "anilkulkarni",
		email:    "anil@example.com",
		age:      14,
	}
	fmt.Println("user1:", user1)
	// print with details
	fmt.Printf("%#v\n", user1)

	// using the shorthand syntax
	user2 := User{username: "gopher", email: "gopher@go.com", age: 11}
	fmt.Println("user2:", user2)

	// declares a `User` with default zeroth value of all members
	user3 := User{}
	fmt.Println(user3)

	// declare a pointer to the struct
	user4 := &User{}
	// using pointer syntax
	(*user4).username = "tina"
	// convenience syntax for working with struct pointers
	user4.email = "tina@example.com"
	user4.age = 13
	fmt.Println(user4)

	// direct assignment
	user5 := User{"python", "python@example.com", 35}
	fmt.Println(user5)
}
