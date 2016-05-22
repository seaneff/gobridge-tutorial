// structs are typed collections of fields.
// They're useful for grouping data together to form
// records

package main

import "fmt"

// This vegetable struct type has name, color, and goodness fields.
type vegetable struct {
    name string
    color  string
    goodness int
}

func main() {

    // This syntax creates a new struct.
    fmt.Println(vegetable{"eggplant", "purple", 5})

    // You can name the fields when initializing a struct
    fmt.Println(vegetable{name: "carrot", color: "orange", goodness: 9})

    // Omitted fields will be zero-valued.
    fmt.Println(vegetable{name: "broccoli", color: "green"})

    // An & prefix yields a pointer to the struct.
    fmt.Println(&vegetable{name: "spinich", color: "green", goodness: 7})

    // Access struct fields with a dot.
    s := vegetable{name: "kohlrabi", color: "green", goodness:6}

    fmt.Println(s.name)

    // You can also use dots with struct pointers - the
    // pointers are automatically dereferenced.
    sp := &s
    fmt.Println(sp)    
    fmt.Println(sp.goodness)

    // Structs are mutable.
    sp.goodness = 7
    fmt.Println(sp)
}