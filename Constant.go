package main

import "fmt"

func main() {
    // Declaring constants
    const Pi float64 = 3.14159
    const Name = "Go Programming"

    // Using constants
    fmt.Println("Value of Pi:", Pi)
    fmt.Println("Language:", Name)

    // Constants can be used in calculations
    radius := 10
    area := Pi * float64(radius) * float64(radius)
    fmt.Println("Area of a circle with radius", radius, ":", area)
}
