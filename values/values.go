package main

import (
	"fmt"
)

// varius value types (string, integer, float, boolean, etc)

func main() {
	// string
	fmt.Println("go " + "lang")
	// integer and float
	fmt.Println("2+3 =", 2+3)
	fmt.Println("5.0 / 2 = ", 5.0/2.0)
	// boolean
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(5 + 2)
	fmt.Print(5 + 2)
	// fmt.Printf(5.2 + 2.2)
}
