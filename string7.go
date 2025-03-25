package main

import "fmt"
import "strings"

func main() {
    str := "Hello, World!"
    prefix := "Hello"
    if strings.HasPrefix(str, prefix) {
        fmt.Println("String starts with prefix!")
    } else {
        fmt.Println("String does not start with prefix.")
    }
}
