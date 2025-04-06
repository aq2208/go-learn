package main

import (
	"fmt"
)

// A struct is a composite data type that groups together zero or more named values
// of arbitrary types as a single entity.
// Each value is called a field -> A struct is a collection of fields

// Declaring a struct
// A struct type is defined by a keyword struct followed by a list of fields inside curly braces.
type person struct {
	firstName string
	lastName  string
	age       int
}

type school struct {
	name       string
	address    string
	tuitionFee int
	persons    []person // struct inside a struct
}

// A struct can also contain a field that is another struct (Embedded struct)
type InternaltionalSchool struct {
	school
	Languages []string
}

// Constructor function
// A constructor function is a function that returns a new instance of a struct.
func newPerson(firstName string, lastName string, age int) person {
	return person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

// Constructor function with pointer
func newPersonPointer(firstName string, lastName string, age int) *person {
	return &person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func main() {
	fmt.Println("Structs in Go")

	// Creating a struct
	// A struct value is created by specifying the field values in curly braces.
	// The order of the field values must match the order of the fields in the struct type.
	person1 := person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	person2 := person{"Anh Quan", "Nguyen", 22}

	// If a field is not specified, it is initialized with the zero value of its type.
	// When initializing a struct value, the field names can be omitted if the values are specified in the order of the fields.
	person3 := person { firstName: "Ngoc" }

	school1 := school{
		name:       "FPT University",
		address:    "Hoa Lac, Hanoi",
		tuitionFee: 30000000,
		persons:    []person{person1, person2, person3},
	}

	internationalSchool1 := InternaltionalSchool { 
		school: school {
			name:       "FPT International School",
			address:    "Hoa Lac, Hanoi",
			tuitionFee: 50000000,
			persons:    []person{person1, person2, person3},
		},
		Languages: []string{"English", "French", "Japanese"},
	}

	// Accessing the fields of a struct
	// The fields of a struct value are accessed using a dot (.) followed by the field name.
	fmt.Println("Person1 First Name: ", person1.firstName)
	fmt.Println("Person1 Last Name: ", person1.lastName)
	fmt.Println("Person1 Age: ", person1.age)

	fmt.Println("Person2 First Name: ", person2.firstName)
	fmt.Println("Person2 Last Name: ", person2.lastName)
	fmt.Println("Person2 Age: ", person2.age)

	fmt.Println("Person3 First Name: ", person3.firstName)
	fmt.Println("Person3 Last Name: ", person3.lastName)
	fmt.Println("Person3 Age: ", person3.age)

	fmt.Println("School Name: ", school1.name)
	fmt.Println("School Address: ", school1.address)
	fmt.Println("School Tuition Fee: ", school1.tuitionFee)
	fmt.Println("School Persons: ", school1.persons)

	fmt.Println("International School Name: ", internationalSchool1.name)
	fmt.Println("International School Address: ", internationalSchool1.address)
	fmt.Println("International School Tuition Fee: ", internationalSchool1.tuitionFee)
	fmt.Println("International School Persons: ", internationalSchool1.persons)
	fmt.Println("International School Languages: ", internationalSchool1.Languages)

	// Modifying the fields of a struct
	person1.age = 40
	fmt.Println("Person 1 Age after modify: ", person1.age)

	// Struct with pointer
	// A pointer to a struct can be created using the address-of operator (&) 
		// and dereferenced using the dereference operator (*).
	// The pointer to a struct is used to modify the fields of the struct.
	person4 := person {
		firstName: "Jane",
		lastName:  "Doe",
		age:       25,
	}
	pointer := &person4
	fmt.Println("Pointer to person4: ", pointer)
	fmt.Println("Person 4 Age through pointer: ", pointer.age)
	pointer.age = 30
	fmt.Println("Person 4 Age after modify through pointer: ", pointer.age)

}
