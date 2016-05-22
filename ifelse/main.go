package main
import "fmt"

func main() {

    my_number := 14
    fmt.Print("My favorite number is ", my_number, ".\n", "")

    if my_number%2 == 0 {
        fmt.Println(my_number, "is even.")
    } else {
        fmt.Println(my_number, "is odd.")
    }

    // You can have an if statement without an else.
    if my_number%7 == 0 {
        fmt.Println(my_number, "is divisible by 7.")
    }

    if num := my_number; num < 0 {
        fmt.Println(num, "is negative.")
    } else if num <= 9 {
        fmt.Println(num, "has 1 digit.")
    } else {
        fmt.Println(num, "has multiple digits.")
    }


    i := 1
    for i <= 10 {
    	if i%2 != 0 {
    		fmt.Println(i, "is odd.")
    	} else {
    		fmt.Println(i, "is even.")
    	}
    	i = i + 1
    }
}
