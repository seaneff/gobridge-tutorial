package main

import "fmt"

func main() {

    // create empty array of zeros that is 10 elements long
    var a [10] int
    fmt.Println("initial array:", a)

    // populate this array with the numbers 1-10
    i := 0
    for i <= 9 {
    	a[i] = i+1
    	//fmt.Println("iteration:", i)
    	//fmt.Println("updated array:", a)
    	i = i + 1
    }

    fmt.Println("updated array:", a)

    // define a bolean aray that is ten elements long
    // starts out as all falses
    var is_even [10] bool

    // loop through the first array, and define whether or not each element of that array is even or odd
    j := 0
    for j <= 9 {
    	is_even[j] = (a[j]%2 == 0)
    	j = j + 1
	}
 
    fmt.Println("\nAre the elements of the array even?\n", is_even)


    // Array types are one-dimensional, but you can
    // compose types to build multi-dimensional data
    // structures.

    var matrix [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            matrix[i][j] = i + j
        }
    }
    
    fmt.Println("\ndefine a matrix:", matrix)

}
