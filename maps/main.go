package main
import "fmt"

func main() {

    // define and print an empty map
    preferences := make(map[string]string)
    fmt.Println("empty map:", preferences)

    // set key/value pairs using typical name[key] = value syntax.
    preferences["beer"] = "Moose Drool"
    preferences["coffee"] = "medium roast"
    preferences["wine"] = "Cabernet Sauvignon"
    preferences["meat"] = "lamb"
    preferences["pasta"] = "linguine"

    fmt.Println("updated map:", preferences)

   // define a map from integers to booleans
    a := make(map[int]bool)

    i := 1
    for i <= 10 {
        a[i] = i%2 == 0
        // fmt.Println(a)
        i = i + 1
    }

    fmt.Println("The number 5 is even:", a[3])
}