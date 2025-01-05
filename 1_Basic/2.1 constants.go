package main

import "fmt"

func main() {
	fmt.Println("Constant variables")
	const c1 string = "circle"
	const c2 = "circle"
	const (
		c3 = "circle"
		s1 = "square"
	)

	// Typed Constant
	// A const declared specifying the type in the declaration is a typed constant.
	// For example below we are declaring a const of type int32
	const a int32 = 8

	// Untyped Constant
	// An untyped constant is a constant whose type has not been specified.
	// A untyped constant in GO can be either named or unnamed.
	// In both cases it doesn’t have any type associated with it.

	const a1 = 123     //Default hidden type is int
	const b = "circle" //Default hidden type is string
	const c = 5.6      //Default hidden type is float64
	const d = true     //Default hidden type is bool
	const e = 'a'      //Default hidden type is rune
	const f = 3 + 5i   //Default hidden type is complex128
}
