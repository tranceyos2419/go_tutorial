package main

import "fmt"

type person struct {
	name     string
	age      int
	passport Passport
}

type Passport struct {
	nationatlity string
}

func main() {
	yoshi := person{name: "yoshi", age: 22, passport: Passport{nationatlity: "japan"}}
	fmt.Println("Hi, my name is ", yoshi.name, "I am from", yoshi.passport.nationatlity)
	fmt.Println(yoshi)
}
