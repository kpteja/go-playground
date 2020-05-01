package main

import (
	"fmt"
)

// Add adds two numbers.
func Add(x, y int) int {
	return x + y
}

// Subtract subtracts two numbers.
func Subtract(x, y int) int {
	return x - y
}

func main() {
	x, y := 8, 2
	sum := Add(x, y)
	diff := Subtract(x, y)
	fmt.Printf("Sum of %d and %d is %d\n", x, y, sum)
	fmt.Printf("Difference of %d and %d is %d\n", x, y, diff)
}
