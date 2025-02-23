package main

import "fmt"

type Student struct {
	ID int
	Name string
	Address string
}

func main() {
	// Declare and initialize the map
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(m)
	
	mapStudent := map[string]Student{
		"Quan": {ID: 1, Name: "Nguyen Anh Quan"},
	}
	fmt.Println(mapStudent)
	
	// Adding and Updating elements in the map
	m["key4"] = "value4"
	m["key1"] = "NAQ"
	fmt.Println(m)
	
	// Delete elements in the map
	delete(m, "key4")
	fmt.Println(m)

	// Retrieve an element from the map
	key1 := m["key1"]
	fmt.Println(key1)

	// Retrieve an element from the map, and also check exist
	value, exist := m["key3"]
	if exist {
		fmt.Println("Value: ", value)
	} else {
		fmt.Println("Key3 not exist in the map")
	}

	// Accessing length and capacity of the map
	fmt.Printf("Map length: %d\n", len(m))

	// Iterate over the map
	for key, value := range m {
		fmt.Println(key, ": ", value)
	}
}