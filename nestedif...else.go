package main

import "fmt"

func main() {
    number := 6  // You can change this number to test

    if number > 0 {
        fmt.Println("The number is positive.")
        if number%2 == 0 {
            fmt.Println("The number is even.")
        } else {
            fmt.Println("The number is odd.")
        }
    } else if number < 0 {
        fmt.Println("The number is negative.")
        if number%2 == 0 {
            fmt.Println("The number is even.")
        } else {
            fmt.Println("The number is odd.")
        }
    } else {
        fmt.Println("The number is zero.")
    }
}
