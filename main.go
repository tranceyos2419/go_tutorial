package main

import "fmt"

func main() {
	r := sum(2, 3)
	println(r)
	fmt.Println(div(6, 3))

}

func sum(x int, y int) int {
	return x + y
}

// named return
func div(x, y int) (z int) {
	z = x / y
	return
}
