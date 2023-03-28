package main

import "math"

type Shape interface {
	area() float64
}

type CircleStruct struct {
	r float64
}

type RectangleStruct struct {
	w, h float64
}

func (circle CircleStruct) area() float64 {
	return circle.r * circle.r * math.Pi
}

func (rect RectangleStruct) area() float64 {
	return rect.w * rect.h
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := CircleStruct{r: 5}
	rect := RectangleStruct{w: 3, h: 4}
	println("area of circle is ", circle.area())
	println("area of rectangle is ", rect.area())
}
