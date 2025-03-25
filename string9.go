package main

import "fmt"
import "strings"

func main() {
    str := "Hello, World!"
    newStr := strings.Replace(str, "World", "Go", -1)
    fmt.Println(newStr)
}
