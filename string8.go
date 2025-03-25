package main

import "fmt"
import "strings"

func main() {
    str := "Hello, World!"
    suffix := "World!"
    if strings.HasSuffix(str, suffix) {
        fmt.Println("String ends with suffix!")
    } else {
        fmt.Println("String does not end with suffix.")
    }
}
