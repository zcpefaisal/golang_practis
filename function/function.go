package main

// function that takes two ints and returns their sum as an int.
func add(a int, b int) int {
	return a + b
}

// When you have multiple consecutive parameters of the same type,
// you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func add_3_number(a, b, c int) int {
	return a + b + c
}

func main() {
	x := add(2, 3)
	println(x)
	y := add_3_number(1, 3, 5)
	println(y)
}
