package main

import "fmt"

type Rectangle struct {
	sideA float64
	sideB float64
}

func (r Rectangle) area() float64 {
	return r.sideA * r.sideB
}

func (r *Rectangle) inc() {
	r.sideA += 5
	r.sideB += 5
}

func (src *Rectangle) isSmaller(cmp *Rectangle) bool {
	return cmp.area() < src.area()
}

func main() {
	r1 := Rectangle{
		sideA: 4,
		sideB: 5,
	}
	fmt.Println("r1 area:", r1.area())

	r2 := &Rectangle{
		sideA: 2,
		sideB: 3,
	}
	r2.inc()
	fmt.Println("inc r2:", r2)

	fmt.Println("Is r1 smaller than r2:", r2.isSmaller(&r1))
	fmt.Println("Is r2 smaller than r1:", r1.isSmaller(r2))
}
