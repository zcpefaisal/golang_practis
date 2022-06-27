package main

import (
	"fmt"
	"time"
)

func main() {

	// basic switch
	i := -2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("this is getter then three or less then one")
	}

	// can use commas to separate multiple expressions in the same case statement
	// use the optional default case in this
	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("this is weekend day")
	default:
		fmt.Println("this is week day")
	}

	// switch without an expression is an alternate way to express if/else logic.
	// also show how the case expressions can be non-constants
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("this is before noon")
	default:
		fmt.Println("this is after noon")
	}

	// A type switch compares types instead of values. You can use this to discover the type of an interface value.
	// In this example, the variable t will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Dont now type %T", t)
		}
	}

	whatAmI(10)
	whatAmI(true)
	whatAmI("hello faisal")

}
