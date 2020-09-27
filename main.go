package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 2
	s := strconv.FormatInt(int64(i), 10)
	fmt.Println(s)
}
