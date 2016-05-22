// Interfaces are named collections of method signatures.

package main

import "fmt"

// define an entity interface
type entity interface {
  name() string
  kind() string
}

// define a function that operates on entity

func identify(e entity) {
  fmt.Printf("entity %s is a %s\n", e.name(), e.kind())
}


// define two structures, person and favorite_drink, that each implement the entity interface
type person struct {
    my_name, my_kind string
}

func (p *person) name() string {
    return p.my_name
}

func (p *person) kind() string {
    return p.my_kind
}

type favorite_drink struct {
    my_name, my_kind string
}

func (d *favorite_drink) name() string {
    return d.my_name
}

func (d *favorite_drink) kind() string {
    return d.my_kind
}

func main() {

    // define a person and their favorite drink
    p := person{my_name: "Zach", my_kind:"person"}
    a := person{my_name: "Steph", my_kind:"person"}
    d := favorite_drink{my_name: "Founders Breakfast Stout", my_kind:"drink"}
    b := favorite_drink{my_name: "coffee", my_kind:"drink"}

    identify(&p)
    identify(&d)
    identify(&a)
    identify(&b)
    identify(&favorite_drink{"negroni", "drink"})

}