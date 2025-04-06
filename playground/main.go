package main

import (
	"fmt"
)

func main() {
	// sum1 := (add(5,6))
	// fmt.Println(sum1)
	// sum2 := sum1
	// sum2 = 100
	// fmt.Println(sum2)
	// fmt.Println(sum1)

	a := 10
	fmt.Println(a)
	
	modifyParameter(a)
	fmt.Println(a)

	modifyPointer(&a)
	fmt.Println(a)

	// v := Vertex{X: 1, Y: 2}
	// fmt.Println(v.X)
	// fmt.Println(v.Y)

	// // p := &v
	// fmt.Println(&v)
	// fmt.Println(&v.X)
	// fmt.Println(&v.Y)
}

type Vertex struct {
	X int
	Y int
}

func add(a, b int) (sum int) {
	sum = a + b
	return
}

func modifyParameter(a int) {
	a = 100
	fmt.Println(a)
}

func modifyPointer(a *int) {
	*a = 100
	fmt.Println(*a)
}