package main

import "fmt"

func main() {
    a, b := 0, 1
    for a <= 50 {
        fmt.Println(a)
        a, b = b, a+b
    }
}
