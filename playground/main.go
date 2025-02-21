package main

import (
	"fmt"
)

func main() {
	sum1 := (add(5,6))
	fmt.Println(sum1)
	sum2 := sum1
	sum2 = 100
	fmt.Println(sum2)
	fmt.Println(sum1)
}

func add(a, b int) (sum int) {
	sum = a + b
	return
}