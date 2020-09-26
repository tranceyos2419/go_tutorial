package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	res, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}
