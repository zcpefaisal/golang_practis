package main

import "fmt"

func main() {
	// basic loop with a single condition.
	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	// clasic loop with iteration, condition, after for loop
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		println("loop")
		break
	}

	// can also continue to the next iteration of the loop.
	for k := 0; k < 10; k++ {
		if k%2 > 0 {
			continue
		} else {
			fmt.Println(k)
		}

	}

}
