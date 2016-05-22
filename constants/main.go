package main

import "fmt"
import "math"

// const declares a constant value that you cannot assign over
const my_drink string = "coffee"

func main() {
    fmt.Println(my_drink)

    // a const statement can appear anywhere a var statement can
    const my_number = 14

    // Constant expressions perform arithmetic with arbitrary precision.
    const a = 14e10 / my_number
    fmt.Println(a)

    fmt.Println(int64(a))

    fmt.Println(math.Sin(a))
}