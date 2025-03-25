package main

import "fmt"
import "strings"

func main() {
    str := "Hello, World!"
    substring := "World"
    if strings.Contains(str, substring) {
        fmt.Println("Substring found!")
    } else {
        fmt.Println("Substring not found.")
    }
}
