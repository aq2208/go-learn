package main

import (
	"basics/utils"
	"fmt"
	"math/cmplx"
	"math/rand"
)

var globalVal bool
var globalVal2 = 10

func main() {
	// import the packages and uses the there functions
	fmt.Println("=====Imported packages=====")
	fmt.Println("Hello World")
	fmt.Println(utils.ConcatString("Hello", "World"))
	fmt.Println(utils.Add(5, 6))
	fmt.Println(utils.Subtract(10, 5))
	fmt.Println("Random number: ", rand.Int())
	fmt.Println()

	// function calls
	fmt.Println("=====Function calls=====")
	fmt.Println(add(5, 6))
	fmt.Println(multiply(2, 3, 4))
	a, b := "hello", "world"
	fmt.Println(swap(a, b))
	fmt.Println(namedReturn(5, 6))
	fmt.Println()

	// variables
	fmt.Println("=====Variables=====")
	var x, y int = 10, 20
	x = 15
	var typeInference = 30
	shorthand, shortHandBool := 40, false
	var str string = "Hello World"
	fmt.Println(x, y, str, globalVal, globalVal2, typeInference, shorthand, shortHandBool)

	// type inference
	v := 42 // change this value to see the type
	fmt.Printf("v is of type %T\n", v)

	var i int
	j := i // j is an int
	fmt.Printf("j is of type %T\n", j)

	fmt.Println()

	// data types
	fmt.Println("=====Data types=====")
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var byteVal byte = 'a' // byte is an alias for uint8
	var runeVal rune = 10
	fmt.Println(byteVal, string(byteVal), runeVal)

	var intVal int = 10
	var floatVal float64 = float64(intVal)
	fmt.Println(intVal, floatVal)

	str = "Hello"
	byteSlice := []byte(str)    // Convert string to byte slice
	newStr := string(byteSlice) // Convert byte slice back to string

	fmt.Println(byteSlice) // Output: [72 101 108 108 111] (ASCII values)
	fmt.Println(newStr)    // Output: Hello
	fmt.Println()

	// constants
	fmt.Println("=====Constants=====")
	// Constants cannot be declared using the := syntax
	// Constants must be initialized at the time of declaration
	const Pi = 3.14
	const Truth = true
	const Constant string = "Hello World"
	fmt.Println(Pi, Truth, Constant)
	fmt.Println()

	// operators
	fmt.Println("=====Operators=====")
	fmt.Println()

	// loops
	fmt.Println("=====Loops=====")
	// basic for loop
	for i := 0; i <= 5; i++ {
		fmt.Println("Iteration: ", i)
	}

	// the init and post statements are optional
	// for loop with single condition (like while loop)
	sum := 1
	for sum < 5 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println()

	// control structures
	fmt.Println("=====Control structures=====")

	// if else
	x = 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// if with a short statement
	if y := 10; y > 5 {
		fmt.Println("y is greater than 5")
	}

	// switch case
	switch day := "Monday"; day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Weekend is near")
	default:
		fmt.Println("Regular day")
	}

	// switch without a condition
	// switch without a condition is the same as switch true
	// cases can be non-constants
	today := "Monday"
	switch {
	case today == "Monday":
		fmt.Println("Start of the week")
	case today == "Friday":
		fmt.Println("Weekend is near")
	default:
		fmt.Println("Regular day")
	}

	// defer
	// A defer statement defers the execution of a function until the surrounding function returns
	defer fmt.Println("Defer - Last thing to print")
	fmt.Println("Before panic")
	panic("Something went wrong")
	fmt.Println("After panic")

	fmt.Println()
}

// function structure
// func functionName(arg1 type, arg2 type) returnType {
// 	// function body
// }

// functions demo
func add(arg1 int, arg2 int) int {
	fmt.Println("Function called")
	return arg1 + arg2
}

// function with multiple arguments with same type
func multiply(a, b, c int) int {
	return a * b * c
}

// function with multiple return values
func swap(a, b string) (string, string) {
	return b, a
}

// function with named return values
func namedReturn(a, b int) (sum int) {
	sum = a + b
	return
}
