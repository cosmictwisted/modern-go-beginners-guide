package main

import "fmt"

func main() {
	// 1. Create a slice of alphabets from `a` to `d`
	a := []string{"a", "b", "c", "d"}
	fmt.Println(a)

	// 2. add `e` at the beginning
	a = append([]string{"e"}, a...)
	fmt.Println(a)

	// 3. add `f` at the end and delete `e` from the beginning
	a = append(a, "f")
	a = a[1:]
	fmt.Println(a)

	// 4. insert `e` at the right location
	a = append(a, "")
	copy(a[4:], a[3:])
	a[4] = "e"
	fmt.Println(a)

	// 5. add alphabets from `g` to `j`
	b := []string{"g", "h", "i", "j"}
	a = append(a, b...)
	fmt.Println(a)

	// 6. replace `d` with `k`
	a[3] = "k"
	fmt.Println(a)

	// 7. delete `k`
	a = append(a[:3], a[4:]...)
	fmt.Println(a)

	// 8. insert `d` at the right location
	a = append(a, "")
	copy(a[3:], a[2:])
	a[3] = "d"
	fmt.Println(a)

	// 9. swap `b` and `h`
	a[1], a[7] = a[7], a[1]
	fmt.Println(a)
}
