package main

import "fmt"

// Go has built-in support for multiple return values.
// This feature is used often in idiomatic Go, for example to return both result and error values from a function.

// The (int, int) in this function signature shows that the function returns 2 ints.
func test() (int, int) {
	return 5, 6
}
func main() {
	// Here we use the 2 different return values from the call with multiple assignment.
	a, b := test()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values, use the blank identifier _.
	_, c := test()
	fmt.Println(c)

	// If you only want a subset of the returned values, use the blank identifier _.
	d, _ := test()
	fmt.Println(d)
}
