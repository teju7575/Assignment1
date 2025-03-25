//This program demonstrates the usage of float32 and float64 data types to store decimal values.
package main

import "fmt"

func main() {
    var num1 float32 = 3.14
    var num2 float64 = 2.718

    fmt.Println("Float32 value:", num1)
    fmt.Println("Float64 value:", num2)

    // Performing calculations with different float types
    result := float64(num1) + num2
    fmt.Println("Sum of float32 and float64:", result)
}
