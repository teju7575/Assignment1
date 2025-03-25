package main

import "fmt"

func main() {
    var num int = 42
    var ptr *int = &num  // Pointer to num

    fmt.Println("Value of num:", num)       // Prints the value of num
    fmt.Println("Memory address of num:", &num)  // Prints the memory address of num
    fmt.Println("Value of ptr:", ptr)        // Prints the memory address stored in ptr
    fmt.Println("Value at the address pointed by ptr:", *ptr) // Dereferencing the pointer to get the value at that address
}
