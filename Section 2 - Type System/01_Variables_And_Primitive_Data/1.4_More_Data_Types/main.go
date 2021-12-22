package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var (
		// unsigned 8-bit integer
		ip    uint8
		price float32
		num   int32
		// bytes under the hood are unsigned 8-bit integer
		by byte
	)

	ip = 128
	// ip = 256
	// ip = -2

	price = 45.678

	// the default value for float is float64
	latitude := 23.3456

	num = -9000

	by = 'A'
	// by = 1000

	fmt.Printf("Value: %d\tType: %s\tSize: %d bytes\n", ip, reflect.TypeOf(ip), unsafe.Sizeof(ip))
	fmt.Printf("Value: %.2f\tType: %s\tSize: %d bytes\n", price, reflect.TypeOf(price), unsafe.Sizeof(price))
	fmt.Printf("Value: %.2f\tType: %s\tSize: %d bytes\n", latitude, reflect.TypeOf(latitude), unsafe.Sizeof(latitude))
	fmt.Printf("Value: %d\tType: %s\tSize: %d bytes\n", num, reflect.TypeOf(num), unsafe.Sizeof(num))
	fmt.Printf("Value: %c\tType: %s\tSize: %d bytes\n", by, reflect.TypeOf(by), unsafe.Sizeof(by))
	fmt.Printf("Value: %v\tType: %s\tSize: %d bytes\n", by, reflect.TypeOf(by), unsafe.Sizeof(by))
}
