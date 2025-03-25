package main

import "fmt"

// Defining a struct
type Person struct {
    name string
    age  int
}

func main() {
    // Initializing a struct
    p := Person{name: "Alice", age: 25}
    fmt.Println("Before modification:", p)

    // Creating a pointer to the struct
    ptr := &p

    // Modifying struct values using the pointer
    ptr.age = 30
    ptr.name = "Bob"

    fmt.Println("After modification:", p)  // Struct fields are modified using the pointer
}
