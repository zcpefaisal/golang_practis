// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
package main

import "fmt"

// function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Println(nums)

	total := 0

	// Within the function, the type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 5, 6, 8, 7)

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum(numbers...)
}
