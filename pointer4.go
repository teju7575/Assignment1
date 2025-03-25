package main

import "fmt"

// Function that takes a pointer as an argument
func incrementValue(num *int) {
    *num = *num + 1
}

func main() {
    num := 5
    fmt.Println("Before increment:", num)

    incrementValue(&num)  // Passing the address of num to the function

    fmt.Println("After increment:", num)  // The value of num has been modified
}
