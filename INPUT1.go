package main

import "fmt"

func main() {
    var num int

    // Prompt the user for input
    fmt.Println("Enter an integer:")

    // Read the input from the user
    fmt.Scan(&num)

    // Print the entered value
    fmt.Println("You entered:", num)
}
