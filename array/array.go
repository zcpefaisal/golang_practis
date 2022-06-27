package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty array:", a)

	a[4] = 100
	fmt.Println("set data:", a)
	fmt.Println("get data:", a[4])

	// The builtin len returns the length of an array.
	fmt.Println("length of array:", len(a))

	// Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	var two [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			two[i][j] = i + j
		}
	}
	fmt.Println(two)
}
