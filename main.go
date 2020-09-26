package main

import "fmt"

func main() {
	println("nubmer")
	for i := 0; i < 5; i++ {
		println(i)
	}

	println("array")
	a := []string{"Foo", "Bar"}
	for i, s := range a {
		fmt.Println(i, s)
	}

	println("map")
	m := make(map[string]string)
	m["me"] = "yoshi"
	m["sister"] = "yuki"
	for k, v := range m {
		println(k, v)
	}
}
