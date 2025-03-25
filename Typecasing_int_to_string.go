package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num1 int = 100
    var str1 string = strconv.Itoa(num1)  // Converting int to string

    fmt.Printf("Integer value: %d\n", num1)
    fmt.Printf("Converted to string: %s\n", str1)
}
