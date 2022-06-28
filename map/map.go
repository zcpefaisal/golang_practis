package main

import "fmt"

func main() {
	// create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax
	m["key1"] = 10
	m["key2"] = 50

	fmt.Println(m)

	// Get a value for a key with name[key1].
	value1 := m["key1"]

	fmt.Println(value1)
	fmt.Println(len(m))

	// builtin delete removes key/value pairs from a map.
	delete(m, "key1")
	fmt.Println(m)
	fmt.Println(len(m))

	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	prs := m["key2"]
	fmt.Println(prs) // 50

	// we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, prs2 := m["key2"]
	fmt.Println(prs2) // true

	_, prs3 := m["key3"]
	fmt.Println(prs3) // false

	// can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 20, "bar": 60}
	fmt.Println(n)

}
