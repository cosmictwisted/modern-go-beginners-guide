package main

import "fmt"

type User struct {
	firstName string
	lastName  string
}

// create a method on `User`
func (user User) greet() {
	fmt.Println("Hello,", user.firstName)
}

// create a method on `User` which returns a string
func (user User) getFullName() string {
	return fmt.Sprintf("%s %s", user.firstName, user.lastName)
}

// methods can also accept arguments
func (user User) getFullNameWithAge(age int) string {
	return fmt.Sprintf("%s %s is %d years old", user.firstName, user.lastName, age)
}

// methods can also return multiple values just as regular functions
func (user User) getFullNameWithAgeAndStatus(age int) (string, bool) {
	isAdult := false
	if age >= 18 {
		isAdult = true
	}
	return fmt.Sprintf("%s %s is %d years old", user.firstName, user.lastName, age), isAdult
}

func main() {
	user1 := User{
		firstName: "Anil",
		lastName:  "Kulkarni",
	}

	user1.greet()
	fmt.Println(user1.getFullName())
	fmt.Println(user1.getFullNameWithAge(14))

	_, isAdult := user1.getFullNameWithAgeAndStatus(14)
	fmt.Println("Adult:", isAdult)
}
