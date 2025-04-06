package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers in Go")

	// In Go, a pointer is a variable that stores the memory address of another variable.

	// Declaring a pointer
	// The type *T is a pointer to a T value. Its zero value is nil.
	var pointer *int  // pointer to an integer (stores the memory address of an integer)

	// Assigning a value to the pointer
	// The & operator generates a pointer to its operand
	i := 10  // integer variable
	pointer = &i  // storing the memory address of i in pointer

	// Accessing the memory address of i
	fmt.Println("Address of i: ", pointer)

	// Accessing the value of i using the pointer (dereferencing)
	fmt.Println("Value of i: ", *pointer)

	// Changing the value of i using the pointer (indirect assignment)
	*pointer = 20
	fmt.Println("Value of i after modify pointer's value: ", i)
}