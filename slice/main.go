package main

import "fmt"

func main() {
	// The type []T is a slice with elements of type T.
	// A slice is a dynamically-sized, flexible view into the elements of an array.
	// Slices are much more common than arrays in practice.

	// Declaring a slice
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	var s []int // slice of integers
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4)
	fmt.Println("Slice s: ", s)

	// Declaring and initializing a slice
	t := []int{1, 2, 3}
	fmt.Println("Slice t: ", t)

	// Slice length and capacity
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	fmt.Println("Length of slice t: ", len(t))
	fmt.Println("Capacity of slice t: ", cap(t))

	// Accessing elements of a slice
	fmt.Println("Element at index 0: ", t[0])

	// Modifying elements of a slice
	t[0] = 10
	fmt.Println("Slice t after modify element at index 0: ", t)

	// Slicing a slice
	// Slicing a slice creates a new slice that points to the same array.
	// The result of slicing a slice s is a new slice that starts at some element of s and ends at or after the last element of s.
	// The result of slicing s[i:j] is a slice of the elements from s[i] to s[j-1], inclusive.
	// The result of slicing s[:j] is a slice of the elements from the beginning of s to s[j-1], inclusive.
	// The result of slicing s[i:] is a slice of the elements from s[i] to the end of s.
	// The result of slicing s[:] is a slice of all of the elements of s.
	u := t[1:3]
	fmt.Println("Slice u: ", u)
	v := t[:2]
	fmt.Println("Slice v: ", v)

	// Appending to a slice
	// The built-in append function appends elements to the end of a slice.
	// If the backing array of a slice is too small to fit all the given values a bigger array will be allocated.
	// The returned slice will point to the newly allocated array.
	t = append(t, 4)
	fmt.Println("Slice t after append element 4: ", t)
	t = append(t, 5, 6)
	fmt.Println("Slice t after append elements 5 and 6: ", t)
	t = append(t, u...)  // Append two slices together

	// Copying a slice
	// The built-in copy function copies elements from a source slice into a destination slice.
	// The destination slice will be modified to have the same length as the source slice.
	// The copy function returns the number of elements copied.
	w := make([]int, len(t))
	copy(w, t)
	fmt.Println("Slice w after copy from slice t: ", w)

	// Iterating over a slice
	// The range form of the for loop iterates over a slice or map.
	// When ranging over a slice, two values are returned for each iteration.
	// The first value is the index, and the second value is a copy of the element at that index.
	for i, v := range t {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// Slice of slices
	matrix := [][]string{
		{"A", "_", "_"},
		{"_", "B", "_"},
		{"_", "_", "C"},
	}
	fmt.Println(matrix)

	for _, value := range matrix {
		fmt.Println("Matrix value:", value)
	}
	for index := range matrix {
		fmt.Println("Index in matrix: ", index)
	}

	// Slice resizing
	slice1 := []int{1,2,3}
	fmt.Printf("Slice1: %d with length: %d with capacity: %d\n", slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 4)
	fmt.Printf("Slice1: %d with length: %d with capacity: %d\n", slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 5,6,7,8,9)
	fmt.Printf("Slice1: %d with length: %d with capacity: %d\n", slice1, len(slice1), cap(slice1))
}