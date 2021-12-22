package main

import "fmt"

type Transport interface {
	Describe(name string)
	Run() bool
}

type Car struct {
	mfg   string
	year  uint
	doors uint
}

type Motorcycle struct {
	wheels uint
	model  string
	price  float64
}

func (c *Car) Describe(name string) {
	fmt.Printf("Name: %s\tManufacturer: %s\tYear: %d\tDoors: %d\n", name, c.mfg, c.year, c.doors)
}

func (c *Car) Run() bool {
	fmt.Println("Car is running...")
	return true
}

func (m *Motorcycle) Describe(name string) {
	fmt.Printf("Name: %s\tModel: %s\tWheels: %d\tPrice: $%.2f\n", name, m.model, m.wheels, m.price)
}

func (m *Motorcycle) Run() bool {
	fmt.Println("Motorcycle is running...")
	return true
}

func main() {
	var t1 Transport
	t1 = &Car{mfg: "Tata Motors", year: 2021, doors: 5}
	t1.Describe("Thunder")
	t1.Run()

	var t2 Transport
	t2 = &Motorcycle{model: "Bajaj Pulsar", wheels: 2, price: 1500}
	t2.Describe("Chetak")
	t2.Run()
}
