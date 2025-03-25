package main

import "fmt"

func main() {
    number := 25 // You can change this to test different values

    // Check multiple conditions using if statements
    if number < 10 {
        fmt.Println("The number is less than 10.")
    } else if number >= 10 && number <= 20 {
        fmt.Println("The number is between 10 and 20.")
    } else if number > 20 && number <= 30 {
        fmt.Println("The number is between 21 and 30.")
    } else if number > 30 && number <= 40 {
        fmt.Println("The number is between 31 and 40.")
    } else {
        fmt.Println("The number is greater than 40.")
    }
}
