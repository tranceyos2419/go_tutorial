package main

import "fmt"

type Shape interface {
	Area()
	Volume()
}

type Cube struct {
	Height float64
}

func (c Cube) Area() float64 {
	return 6 * (c.Height * c.Height)
}

func (c Cube) Volume() float64 {
	return c.Height * c.Height * c.Height
}
func main() {

	c := Cube{Height: 165}
	fmt.Println("Area: ", c.Area())
	fmt.Println("Volume: ", c.Volume())
}
