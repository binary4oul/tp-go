package main

import "fmt"

func main() {
	var x float64 = 20.0
	y := "I am a child"
	fmt.Println(x, y)
	fmt.Printf("\a %T, %T\n", x, y)

	// Mixed Variables.
	var a, b, c = 2, "fds", 32.2
	fmt.Println(a, b, c)
}
