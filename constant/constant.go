package main

import (
	"fmt"
	"math"
)

const constantValue = "this is constant value"
const x = 10
const y = 50

func main() {
	fmt.Println(constantValue)
	const a = 50
	const b = 3e20 / a // A const statement can appear anywhere a var statement can.
	fmt.Println(b)
	fmt.Println(int64(b)) // A numeric constant has no type until it’s given one, such as by an explicit conversion.
	fmt.Println(x)
	fmt.Println(y)
	// x = 100 // constant variable cannot reassign
	// y = 200 // constant variable cannot reassign
	fmt.Println(math.Sin(a)) // A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
}
