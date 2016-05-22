package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		printParity(i)
	}
}
func even(x int) bool {
	return x%2 == 0
}

func printParity(x int) {
	if even(x) {
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "is odd")
	}
}
