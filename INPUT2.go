package main

import "fmt"

func main() {
	var name string

	// Prompt the user for input
	fmt.Println("Enter your name:")

	// Read the input from the user
	fmt.Scanln(&name)

	// Print the entered value
	fmt.Println("Hello,", name)
}
