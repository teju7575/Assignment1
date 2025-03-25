package main

import "fmt"

func main() {
    number := 5
    factorial := 1
    i := 1
    for i <= number {
        factorial *= i
        i++
    }
    fmt.Println("Factorial of", number, "is", factorial)
}
