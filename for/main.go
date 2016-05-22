package main

import "fmt"

func main() {

	// The most basic type of for loop in go
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

    // A classic initial/condition/after for loop.
    // this is a shorthand for what we did in the first loop
    // j++ means "add one to j", like saying j = j + 1
    for j := 100; j <= 103; j++ {
        fmt.Println(j)
    }

    // for  without a condition will loop repeatedly
    // until you break out of the loop or return from
    // the enclosing function.
    for {
        fmt.Println("coffee")
        break
    }

    k := 1
    my_beer := "Moose Drool"
    for k <= 3 {
        fmt.Println(my_beer)
        k = k + 1
    }

    z := 2
    for z <= 10000 {
        fmt.Println(z)
        z = z*2
    }
}
