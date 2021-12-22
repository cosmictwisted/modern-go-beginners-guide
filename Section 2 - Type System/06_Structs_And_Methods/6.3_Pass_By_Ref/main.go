package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p Point) changeX() {
	p.x = 2
}

// Pass by ref is a commonly used pattern
func (p *Point) changeXByRef() {
	p.x = 2
}

func main() {
	p := Point{7, 5}
	// pass by value
	p.changeX()
	fmt.Println("p:", p)

	// pass by reference
	// compiler does the conversion to (*p).changeXByRef()
	p.changeXByRef()
	fmt.Println("p:", p)

	p2 := &Point{8, 9}
	p2.changeX()
	// no change in original value
	fmt.Println("p2:", p2)

	p2.changeXByRef()
	fmt.Println("p2:", p2)
}
