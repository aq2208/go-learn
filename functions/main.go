package main

import (
	"fmt"
	"math"
)

func add(a, b int) int {
	return a + b
}

func operate(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// multiplier returns a function that multiplies its input by a preset factor.
func multiplier(factor int) func(int) int {
	return func (n int) int {
		return n * factor
	}
}

func main() {
	// Assign function to a variable
	sum := add
	fmt.Println(sum(5, 6))

	// Passing function as an argument to other function
	fmt.Println(operate(7, 8, add))

	// Return a function from a function
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5))
	fmt.Println(triple(5))

	// Assign function to a variable
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// Passing function as an argument to other function
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Function closures

	// Define a closure that captures the variable 'x'
	x := 10
	addX := func (y int) int {
		return x + y  // 'x' is captured from the surrounding scope
	}
	fmt.Println(addX(10001))

}