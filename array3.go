package main

import "fmt"

func main() {
    // Initialize an array
    arr := [5]int{10, 20, 30, 25, 15}
    max := arr[0] // Start with the first element as the initial maximum

    // Loop through the array to find the maximum value
    for _, value := range arr {
        if value > max {
            max = value
        }
    }

    // Print the maximum value
    fmt.Println("Maximum value in the array:", max)
}
