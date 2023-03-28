package main

import (
	"errors"
	"math"
)

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math::: negative number passed to sqrt")
	}
	return math.Sqrt(value), nil
}

func main() {
	result, err := Sqrt(-1)
	println(result, err.Error())
	result, err = Sqrt(3)
	println(result)

}
