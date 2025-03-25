package main

import "fmt"

func main() {
    var num1 float64 = 12.345
    var num2 int = int(num1)  // Typecasting from float64 to int

    fmt.Printf("Float value: %f\n", num1)
    fmt.Printf("Converted to int: %d\n", num2)
}
