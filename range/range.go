package main

import "fmt"

func main() {

	nums := []int{2, 3, 4, 5}
	sum := 0

	// Here we use range to sum the numbers in a slice. Arrays work like this too.
	for _, val := range nums {
		sum += val
		// println(_)   // here _ is blank key
		println(val) // num
	}
	println(sum)

	// range on arrays and slices provides both the index and value for each entry.
	// Above we didn’t need the index, so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.
	for i, val := range nums {
		sum += val
		println(i)   // here i is index
		println(val) // val is value
		if val == 4 {
			println("index:", i)
		}
	}
	println(sum)

	kvp := map[string]string{"a": "apple", "b": "banana"}
	// range on map iterates over key/value pairs.
	for k, v := range kvp {
		fmt.Printf("%s -> %s\n", k, v)
		println(k, v)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "faisal" {
		println(i, c)
	}
}
