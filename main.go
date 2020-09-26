package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 1)
	// a := [2][2]int{}
	fmt.Println(slice)
}
