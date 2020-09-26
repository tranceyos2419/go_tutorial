package main

import (
	"fmt"
)

func main() {
	square := make(map[string]int)
	square["height"] = 2
	square["width"] = 5
	fmt.Println(square["width"])
}
