package main

import "fmt"
import "strings"

func main() {
    str := "Hello, World!"
    index := strings.Index(str, "World")
    fmt.Println("Index of 'World':", index)
}
