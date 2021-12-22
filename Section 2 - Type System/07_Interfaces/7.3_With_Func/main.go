package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	sideA float64
	sideB float64
}

type Circle struct {
	radius float64
}

func (r *Rectangle) Area() {
	fmt.Printf("Sides: %.2f x %.2f\tArea: %.2f\n", r.sideA, r.sideB, r.sideA*r.sideB)
}

func (r *Rectangle) Perimeter() {
	fmt.Printf("Sides: %.2f x %.2f\tPerimeter: %.2f\n", r.sideA, r.sideB, 2*(r.sideA+r.sideB))
}

func (c *Circle) Area() {
	fmt.Printf("Radius: %.2f\tArea: %.2f\n", c.radius, math.Pi*c.radius*c.radius)
}

func (c *Circle) Perimeter() {
	fmt.Printf("Radius: %.2f\tPerimeter: %.2f\n", c.radius, 2*math.Pi*c.radius)
}

func printAreas(g ...Geometry) {
	for i := 0; i < len(g); i++ {
		g[i].Area()
	}
}

func printPerimeters(g ...Geometry) {
	for i := 0; i < len(g); i++ {
		g[i].Perimeter()
	}
}

func main() {
	r1 := &Rectangle{4, 5}
	r2 := &Rectangle{2, 8}

	c1 := &Circle{13}
	c2 := &Circle{9}

	printAreas(r1, r2, c1, c2)
	fmt.Println()
	printPerimeters(r1, r2, c1, c2)
}
