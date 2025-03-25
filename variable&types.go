package main

import "fmt"

func main() {
    // Declaring variables with 'var'
    var age int = 25
    var name string = "Alice"
    var isActive bool = true
    var height float64 = 5.8

    // Declaring multiple variables in a single line
    var x, y int = 5, 10

    // Shorthand variable declaration
    z := 30
    message := "Hello, Go!"

    // Printing the values of the variables
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Active:", isActive)
    fmt.Println("Height:", height)
    fmt.Println("x:", x)
    fmt.Println("y:", y)
    fmt.Println("z:", z)
    fmt.Println("Message:", message)
}
