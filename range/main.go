package main

import "fmt"

func main() {

    // use use range to sum the numbers in a slice
    nums := []int{5, 10, 15}
    sum := 0
    // _, all elements of num
    for _, num := range nums {
        // sum = sum + number
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    drinks := map[string]string{"beer": "Moose Drool", "drink": "coffee"}
    for k, v := range drinks {
        fmt.Printf("%s -> %s\n", k, v)
    }

   // define a map from integers to booleans
    a := make(map[int]bool)

    i := 1
    for i <= 10 {
        a[i] = i%2 == 0
        // fmt.Println(a)
        i = i + 1
    }

    fmt.Println(a)

    for k,v := range a {
        is_even := "is odd" 

        if v == true {
            is_even = "is even"
        }

        fmt.Println(k, is_even)
    }

}
