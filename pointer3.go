package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}
    fmt.Println("Array:", arr)

    ptr := &arr  // Pointer to the array

    // Accessing elements through the pointer
    fmt.Println("First element through pointer:", (*ptr)[0]) // Dereferencing the pointer to get the first element
}
