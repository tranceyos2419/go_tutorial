package main

import "fmt"

func main() {
	//sum
	r := sum(2, 3)
	println(r)

	// div
	fmt.Println(div(6, 3))

	// swap
	a, b := swap("cat", "dog")
	fmt.Println(a, b)
}

func sum(x int, y int) int {
	return x + y
}

// named return
func div(x, y int) (z int) {
	z = x / y
	return
}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}
