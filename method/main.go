package main

import ("fmt")

func main() {
	p := Person {Name: "Alice", Age: 30}
	p.Greet()
	p.UpdateProfile("Anh Quan", 22)

	f := MyFloat(2.2)
	fmt.Println(f.Add(3.3)) // This will call the Add() method on the MyFloat type
}

type Person struct {
	Name string
	Age  int
}

// The Greet() method is a receiver function that takes a Person object as a receiver.
// That Greet() method can be called on any Person object.
// The receiver is the first parameter of the method, and it is used to access the fields of the struct.
// The receiver is a value receiver, which means that the method operates on a copy of the struct.
// We can only declare a method with a receiver whose type is defined in the same package as the method.
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)

	p.Name = "Bob" // This will not change the original struct
	// because p is a copy of the original struct.
	// If you want to modify the original struct, you can use a pointer receiver.
}

// Pointer receiver
func (p *Person) UpdateProfile(newName string, newAge int) {
	fmt.Println("Hello, my name is - before update", p.Name)

	p.Name = newName // This will change the original struct
	p.Age = newAge  // This will change the original struct
	// because p is a pointer to the original struct.
	// If you want to modify the original struct, you can use a pointer receiver.
	// This is useful when you want to modify the original struct.
	// It is also useful when the struct is large and you want to avoid copying it.

	fmt.Println("Hello, my name is - after update", p.Name)
}

// This function is equivalent to the above Greet() method.
func Greet(p Person) {
	fmt.Println("Hello, my name is", p.Name)
}
// The Greet() function is a standalone function that takes a Person object as an argument.


// We can declare a method with any type of receiver (not just struct type), including built-in types, user-defined types, and interfaces.
type MyFloat float64
func (f MyFloat) Add(x MyFloat) MyFloat {
	return f + x
}