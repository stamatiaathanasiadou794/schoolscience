package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(numbers)
}
