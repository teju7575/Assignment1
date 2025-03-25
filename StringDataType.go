package main

import "fmt"

func main() {
    var greeting string = "Hello, Go!"
    var name string = "Alice"

    // Concatenate strings
    message := greeting + " My name is " + name + "."

    fmt.Println(message)
}
