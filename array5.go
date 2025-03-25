package main

import "fmt"

func main() {
    // Initialize an array
    arr := [5]int{1, 2, 3, 2, 4}
    element := 2
    count := 0

    // Loop through the array and count occurrences of 'element'
    for _, value := range arr {
        if value == element {
            count++
        }
    }

    // Print the count of the element
    fmt.Printf("Element %d occurs %d times in the array.\n", element, count)
}
