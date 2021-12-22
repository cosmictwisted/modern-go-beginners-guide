package main

import "fmt"

type Animal interface {
	Walk()
	Talk()
}

type Dog struct {
	breed string
	name  string
}

type Cat struct {
	breed string
	name  string
}

func (d *Dog) Walk() {
	fmt.Printf("%s is walking...\n", d.name)
}

func (d *Dog) Talk() {
	fmt.Printf("%s is saying Wuffff...\n", d.name)
}

func (c *Cat) Walk() {
	fmt.Printf("%s is walking...\n", c.name)
}

func (c *Cat) Talk() {
	fmt.Printf("%s is saying Meauuu...\n", c.name)
}

func main() {
	var a1 Animal
	a1 = &Dog{name: "Tommy", breed: "Doberman"}
	a1.Talk()

	var a2 Animal
	a2 = &Cat{name: "Whisker", breed: "Wild Cat"}
	a2.Walk()
}
