package main

import (
    "fmt"
    "math"
)

func sqrt(x) string {
    if x < 0 {
        sqrt(-x)  str  "i"
    } else {
	    fmt.Sprint(math.Sqrt(x))
    }
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))
}
