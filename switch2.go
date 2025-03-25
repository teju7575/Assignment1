package main

import "fmt"

func main() {
	var choice int

	// Menu
	fmt.Println("Menu:")
	fmt.Println("1. Monday")
	fmt.Println("2. Tuesday")
	fmt.Println("3. Wednesday")
	fmt.Println("4. Thursday")
	fmt.Println("5. Friday")
	fmt.Println("6. Saturday")
	fmt.Println("7. Sunday")
	fmt.Println("Enter your choice (1-7):")
	fmt.Scan(&choice)

	// Switch-case for days
	switch choice {
	case 1:
		fmt.Println("You selected: Monday")
	case 2:
		fmt.Println("You selected: Tuesday")
	case 3:
		fmt.Println("You selected: Wednesday")
	case 4:
		fmt.Println("You selected: Thursday")
	case 5:
		fmt.Println("You selected: Friday")
	case 6:
		fmt.Println("You selected: Saturday")
	case 7:
		fmt.Println("You selected: Sunday")
	default:
		fmt.Println("Invalid choice! Please select a valid day (1-7).")
	}
}
