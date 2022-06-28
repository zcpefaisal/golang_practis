package main

import (
	"fmt"
)

func main() {
	// create an empty slice with non-zero length, use the builtin make.
	// Here we make a slice of strings of length 3
	s := make([]string, 3)
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])
	// can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])

	// returns the length of the slice as expected.
	fmt.Println(len(s))

	// slices support several more that make them richer than arrays. One is the builtin append,
	// which returns a slice containing one or more new values.
	// Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)
	fmt.Println(len(s))

	// create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	fmt.Println(len(c))

	// Slices support a “slice” operator with the syntax slice[low:high]
	l := s[2:5]
	fmt.Println(l)
	l = s[:5]
	fmt.Println(l)
	l = s[3:]
	fmt.Println(l)

	// can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 4)
	for i := 0; i < 4; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD) // [[0] [1 2] [2 3 4] [3 4 5 6]]

	twoD2 := make([][]int, 4)
	fmt.Println(twoD2)
	for i := 0; i < 4; i++ {
		twoD2[i] = make([]int, 3)
		for j := 0; j < 1; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println(twoD2) // [[0 0 0] [1 0 0] [2 0 0] [3 0 0]]
}
