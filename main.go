package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}
	slice = append(slice[:0], slice[1:]...) // delete
	slice = append(slice, 1)                // add
	fmt.Println(slice)
	//
	cat := []string{"sleepy"}
	cat = append(cat, "kiss")
	fmt.Println(cat)
}
