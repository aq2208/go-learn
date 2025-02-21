package main

import "fmt"

func main() {
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(m)

	// Retrieve an element
	key1 := m["key1"]
	fmt.Println(key1)
}