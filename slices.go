package main

import "fmt"

func main() {

	// unlike arrays slices are typed only by
	// the elements they contain 
	// not the number of elements.
	// to create an empty slice with non-zero
	// length, use the builtin `make`
	// here we make a slice of `strings`s of length
	// `3` (intially zero-values)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like the arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as 
	// expected.
	fmt.Println("len:", len(s))

	/* In addition to these basic operations,
	slices support several more that make them
	richer than arrays. One is a builtin `append`,
	which returns a slice containing one or more
	new values

	Note that we need to accept a return value
	from appen as we may get a new slice value. */ 

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	
