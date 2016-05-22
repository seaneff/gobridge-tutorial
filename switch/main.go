package main

import "fmt"  
import "time" 

func main() {

    // example of a basic `switch`.
    my_drink := "coffee"

    fmt.Print("I am currently drinking a " )
    switch my_drink {
    case "whiskey":
        fmt.Println("whiskey drink")
    case "vodka":
        fmt.Println("vodka drink")
    case "cider":
        fmt.Println("cider drink")
    case "coffee":
        fmt.Println("large cup of coffee")
    }

    fmt.Println("The current time is", time.Now().Format(time.RFC850))


    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend! Hooray!")
    default:
        fmt.Println("It's a weekday. You can do it.")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon.")
    default:
        fmt.Println("It's after noon.")
    }

    i := 1
    for i <= 10 {
    	fmt.Print(i, " is")
    	switch {
    	case i%2 != 0:
    		fmt.Println(" odd.")
    	case i%2 == 0:
    		fmt.Println(" even.")
    	}

    	i = i + 1
    }

}

