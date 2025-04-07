package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write to the file
	_, err = io.WriteString(file, "Hello, World!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Read from the file
	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	content := make([]byte, 64)
	n, err := file.Read(content)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, content[:n])
	// Close the file
	if err := file.Close(); err != nil {
		fmt.Println("Error closing file:", err)
		return
	}

	fmt.Println("File created and written successfully.")
}