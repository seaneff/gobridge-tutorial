// _Slices_ are a key data type in Go, giving a more
// powerful interface to sequences than arrays.

package main

import "fmt"

func main() {

    // create an empty slice of two integers and print it out
    s := make([]int, 0)
    fmt.Println("initial empty slice:", s)

    // loop over the numbers 1-10 and append them to slice
    i := 1
    for i <= 10  {
        s = append(s, i)
        i = i + 1
    }

    fmt.Println("updated slice:", s)

    // print out the first five numbers of the slice
    fmt.Println("first five elements:", s[:5])

    fmt.Println("last five elements:", s[5:])

    fmt.Println("middle elements:", s[4:6])

}
