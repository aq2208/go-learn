package main

import "fmt"

func main() {
	// The type [n]T is an array of n values of type T.

	// Declaring an array
	// An array's length is part of its type, so arrays cannot be resized.
	var a [3]int  // array of 3 integers
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println("Array a: ", a)

	// Declaring and initializing an array
	b := [3]int{1, 2, 3}
	fmt.Println("Array b: ", b)

	// Array length
	fmt.Println("Length of array b: ", len(b))

	// Accessing elements of an array
	fmt.Println("Element at index 0: ", b[0])

	// Modifying elements of an array
	b[0] = 10
	fmt.Println("Array b after modify element at index 0: ", b)
}