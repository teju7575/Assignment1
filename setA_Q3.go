package main

import "fmt"

func main() {
	// Declare two variables
	var a, b int

	// Take input from the user
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	// Display numbers before swap
	fmt.Println("Before Swap: a =", a, "b =", b)

    a = a + b
    b = a - b
    a = a - b

    fmt.Println("After Swap: a =", a, "b =", b)
}