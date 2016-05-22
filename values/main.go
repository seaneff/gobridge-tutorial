package main

import "fmt"

func main() {
	// strings
	fmt.Println("Hi, I'm", "Steph")
	// strings
    fmt.Println("I like to drink", "coffee")
    // strings
    fmt.Println("My favorite color is", "purple")
    // strings
    fmt.Println("My favorite beer is", "Moose Drool from Big Sky Brewing Company")
    // string, integer
    fmt.Println("My favorite number is", 7*2*1)
    // string, float
    fmt.Println("My favorite fraction is one third:", 1.0/3.0)
    // string, float, string
    fmt.Println("Today I have drank", float64(7)/float64(2), "cups of coffee")
    // string, integer, string
    fmt.Println("Today I have drank", 7/2, "full cups of coffee")
    // string, boolean
    fmt.Println("Do I like dance parties?", true)
    // string, boolean
    fmt.Println("Do I like jello?", true && false)
    // string, boolean
    fmt.Println("Do I like statistics?", true || false)
}
