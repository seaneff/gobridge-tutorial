
package main

import "fmt"

// create a function called drinkVal that takes a string value argument
// and assigns the string "Moose Drool" to it.
func drinkVal(val string) {
    val = "Moose Drool"
}

// create a function called drinkPtr that takes a pointer to a string argument
// and assigns the string "coffee" to it

func drinkPtr(ptr *string) {
    *ptr = "coffee"
}


func main() {

    s := "negroni"

    fmt.Println("initial string:", s)

    drinkVal(s)

    fmt.Println("string update 1 (no pointer):", s)

    // The &i syntax gives the memory address of i,
    // we call this a pointer to i

    drinkPtr(&s)

    fmt.Println("string update 2 (with pointer):", s)

    // Pointers can be printed too
    fmt.Println("name of the pointer location:", &s)

}