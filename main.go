package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	yoshi := person{name: "yoshi", age: 22}
	fmt.Println("Hi, my name is ", yoshi.name)
	fmt.Println(yoshi)
}
