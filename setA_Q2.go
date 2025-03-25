package main

import "fmt"

func main() {
    // Declare a variable to store the number
    var num int

    // Take input from the user
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    // Check if the number is even or odd
    if num%2 == 0 {
        fmt.Println(num, "is even.")
    } else {
        fmt.Println(num, "is odd.")
    }
}
