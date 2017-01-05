package main
import "fmt"

func main() { 
	//builtin `make` creates empty map
	m := make(map[string]int)

	// set key/vaul pairs 
	// `name[key] = val`
	m["k1"] = 7
	m["k2"] = 13

	// printing a map with e.g. `Println` will show
	// all of it's key/value pairs.
	fmt.Println("map:", m)

	// get a value for a key with `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// builtin `len` returns the number of key/value
	// pairs when called on a map.
	fmt.Println("len:", len(m))

	// builtin `delete` removes key/value pairs from
	// a map.
	delete(m, "k2")
	fmt.Println("map:", m)
	/* The optional second return value when getting a
	 * value from a map indicates if the key was present
	 * in the map. This can be used to disambiguate
	 * between missing keys and keys with zero values
	 * like `0` or `""`. here we didn't need the value
	 * itself, so we ignored it with the _ blank identifier _
	 * `_`. */ 
	 
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
