package main

import "fmt"

func main() {

    // var can declare 1 or more variable(s)
    var my_name string = "Steph"
    fmt.Println("My name is", my_name)

    my_drink := "coffee"
    fmt.Println("I like to drink", my_drink)

    my_beer := "Moose Drool"
    my_brewery := "Big Sky Brewing Company"
    fmt.Println("My favorite beer is", my_beer, "from", my_brewery)

    my_number := 14
    fmt.Println("My favorite number is", my_number)

    // You can declare multiple variables at once.
    var hawkins, steph, kelly, kodi int = 7, 14, 17, 23
    fmt.Println("Frisbee numbers:", kelly, kodi, hawkins, steph)

    var a = true
    fmt.Println("Do I like cheese?", a)

    a = false
    fmt.Println("Do I like jello?", a)

    // variables declared without a corresponding
    // initialization are _zero-valued_. for example, the
    // zero value for an inteter is 0, and the 
    // default value for a boolean is FALSE
    var b int
    fmt.Println(b)

}