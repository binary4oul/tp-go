package main

import "math"

type Circle struct {
	x, y, r float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.r * circle.r
}

func main() {
	circle := Circle{x: 7, y: 7, r: 3}
	println(circle.area())
}
