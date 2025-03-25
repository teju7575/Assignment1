package main

import (
	"fmt"
)

func main() {
	var choice int
	var num1, num2 float64

	// Menu
	fmt.Println("Menu:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("Enter your choice (1-4):")
	fmt.Scan(&choice)

	// Input numbers
	fmt.Println("Enter two numbers:")
	fmt.Scan(&num1, &num2)

	// Switch-case for operations
	switch choice {
	case 1:
		fmt.Printf("Result: %.2f\n", num1+num2)
	case 2:
		fmt.Printf("Result: %.2f\n", num1-num2)
	case 3:
		fmt.Printf("Result: %.2f\n", num1*num2)
	case 4:
		if num2 != 0 {
			fmt.Printf("Result: %.2f\n", num1/num2)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Invalid choice! Please select a valid option (1-4).")
	}
}
