package main

import "fmt"

func main() {
    var num1 int = 10
    var num2 float64 = float64(num1)  // Typecasting from int to float

    fmt.Printf("Integer value: %d\n", num1)
    fmt.Printf("Converted to float: %f\n", num2)
}
