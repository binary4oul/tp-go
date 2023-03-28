package main

import "fmt"

func main() {
	var a int = 100
	// IF statement
	if a > 100 {
		println("a is larger than 100")
	}
	println("a is smaller than 100")

	// Switch statement
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x :%T", i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("don't know the type")
	}

}
