package main

import (
	"fmt"
)

func main() {

	var x int = 2
	y := 3
	var s1, s2 string
	s1 = "foo"
	s2 = "too"

	var s3, s4 = "yoshi", "takafumi"
	sum := x + y

	// Constants cannot be declared using the := syntax.
	const constant = "can not change later"

	fmt.Println(sum)
	fmt.Println(s1, s2)
	fmt.Println(s3, s4)
	fmt.Println(constant)
}

// Variables declared without an explicit initial value are given their zero value.

// The zero value is:

// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.
