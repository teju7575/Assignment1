package main

import "fmt"

func main() {
    // Initialize an array
    arr := [5]int{1, 2, 3, 4, 5}
    sum := 0

    // Loop through the array to sum its elements
    for _, value := range arr {
        sum += value
    }

    // Print the sum
    fmt.Println("Sum of array elements:", sum)
}
