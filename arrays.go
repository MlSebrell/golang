package main

import "fmt"

func main() {

	var a [5]int // create an array that will hold exactly 5 integers
	fmt.Println("emp:", a)

	// We set a vaule at an index using the
	// `array[index] = value` syntax
	// and get a value with
	// `array[index]`. 
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	
	// `len` returns the length of an array
	fmt.Println("len:", len(a))
	
	// declares and intialize an array in one line.
	b := [5]int{1, 2 ,3 ,4 ,5}
	fmt.Println("dcl:", b)
	
	// Array tpes are one dimensional
	// but you can compose types to build
	// multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
