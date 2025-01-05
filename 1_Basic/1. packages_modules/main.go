// main.go
package main

import (
	"fmt"
	"go-package-example/math" // Import the math package
)

/*
Package in Go:
A package is a collection of Go source files that are stored in a specific directory.
Every Go file belongs to a package, and you can use packages to organize your code logically.
The most common package is main, which is used to define the entry point for your application.

Module in Go:
A module is a collection of Go packages that are versioned together as a unit.
A module is defined by a go.mod file, which records the module's path and its dependencies.

math package: Provides basic mathematical operations like addition and subtraction.
main package: The entry point of the application that uses the math package functions.

1. package_modules/
│
├── go.mod             # Go module file, manages dependencies
├── main.go            # Main entry point of the application
└── math/
    └── math.go        # Package containing math functions (Add, Subtract)

go.mod: Defines the module path and its dependencies.
main.go: The entry point where the program executes.
math/math.go: Contains the mathematical operations like Add and Subtract.
*/

func main() {
	// Calling functions from the math package
	sum := math.Add(5, 3)             // Adding numbers using Add function
	difference := math.Subtract(5, 3) // Subtracting numbers using Subtract function

	// Output the results
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
}
