package main

// import "fmt"

// Interfaces in Go are a way to define a contract that types can implement.
// An interface defines a set of methods that a type must implement.
// This allows for polymorphism, where different types can be treated as the same type
//// if they implement the same interface.
// This is useful for writing generic code that can work with different types.
// In Go, an interface is defined using the `interface` keyword.
// An interface can contain any number of methods, and a type implements an interface
//// by providing implementations for all the methods in the interface.
// A type can implement multiple interfaces, and an interface can be implemented by multiple types.
type Speaker interface {
	Speak() string
}

type Dog struct {}

// The Dog struct implements the Speaker interface by providing a Speak method.
// Implicitly implementing an interface in Go is done by defining the methods of the interface.
// There is no need to explicitly declare that a type implements an interface.
// This is different from some other languages, where you have to explicitly declare that a type implements an interface.
// In Go, if a type has all the methods of an interface, it implements that interface.
// This is done at compile time, so if a type does not implement an interface,
//// the code will not compile.
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {}
func (c Cat) Speak() string {
	return "Meow!"
}

type Cow struct {}
func (c Cow) Speak() string {
	return "Moo!"
}

func main() {
	dog := Dog{}
	cat := Cat{}

	var speaker Speaker

	speaker = dog  // Assigning a Dog to the Speaker interface
	println(speaker.Speak())

	speaker = cat  // Assigning a Cat to the Speaker interface
	println(speaker.Speak())

	// Using the interface as a parameter
	println(MakeSomeNoise(dog))
	println(MakeSomeNoise(cat))

	var speaker2 Speaker = Cow{}
	println(speaker2.Speak())

	// Type assertions
	accessDog, ok := speaker.(Dog)
	if ok {
		println("Dog's sound:", accessDog.Speak())
	} else {
		println("Not a dog")
	}

	// Type switches
	switch s := speaker.(type) {
	case Dog:
		println("It's a dog:", s.Speak())
	case Cat:
		println("It's a cat:", s.Speak())
	case Cow:
		println("It's a cow:", s.Speak())
	default:
		println("Unknown type")
	}
}

func MakeSomeNoise(s Speaker) string {
	return s.Speak()
}