package main

import "fmt"

func main() {
	numbers := []int{123, 234, 345, 456, 567, 678}
	for i, val := range numbers {
		fmt.Printf("Slice Item %d is %d", i, val)
	}

}
