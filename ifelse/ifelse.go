package main

import "fmt"

// Note that you don’t need parentheses around conditions in Go, but that the braces are required.
// There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.

func main() {
	// basic if else example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// can have an if statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// can precede conditionals; any variables declared in this statement are available in all branches.
	if num := -10; num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
