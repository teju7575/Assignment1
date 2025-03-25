package main

import "fmt"

func main() {
    num := 10
    fmt.Println("Before modification, num:", num)

    ptr := &num  // Pointer to num
    *ptr = 20    // Modifying the value of num using the pointer

    fmt.Println("After modification, num:", num) // Prints the updated value of num
}
