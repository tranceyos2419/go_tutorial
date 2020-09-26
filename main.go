package main

import (
	"fmt"
)

func main() {

	z := 20

	if z < 10 {
		fmt.Printf("%d", z)
	} else if z < 50 {
		fmt.Println("more than 10 but less than 50")
	} else {
		fmt.Println("z is too big")
	}

}
