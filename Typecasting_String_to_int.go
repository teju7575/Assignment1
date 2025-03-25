package main

import (
    "fmt"
    "strconv"
)

func main() {
    var str string = "1234"
    num, err := strconv.Atoi(str)  // Converting string to int

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("String value: %s\n", str)
    fmt.Printf("Converted to int: %d\n", num)
}
