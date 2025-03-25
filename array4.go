package main

import "fmt"

func main() {
    // Initialize an array
    arr := [5]int{1, 2, 3, 4, 5}

    // Reverse the array
    for i := 0; i < len(arr)/2; i++ {
        arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
    }

    // Print the reversed array
    fmt.Println("Reversed array:", arr)
}
