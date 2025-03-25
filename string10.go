package main

import "fmt"
import "strings"

func main() {
    str := "Hello World Go Language"
    words := strings.Split(str, " ")
    fmt.Println(words)
}
