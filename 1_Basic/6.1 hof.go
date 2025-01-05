package main

import "fmt"

// Example:
// 1. Passing Functions as Arguments
// You can pass functions to other functions to customize the behavior dynamically.
// Higher-order function
func apply(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Simple functions to pass as arguments
func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

// 2. Returning Functions (Function as Result)
// A higher-order function can return a function that can be used later.
// Function returning another function
func adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// 3. Combining Both – Passing and Returning Functions
// You can create highly flexible functions that take functions as arguments and return functions as results.
// Higher-order function that returns a function
func createOperation(op string) func(int, int) int {
	switch op {
	case "add":
		return func(a, b int) int { return a + b }
	case "multiply":
		return func(a, b int) int { return a * b }
	default:
		return func(a, b int) int { return 0 }
	}
}

// 4. Practical Use Case – Filtering Data
// Higher-order functions are often used for filtering, mapping, and reducing data in slices.
// Example – Filtering Slice of Integers:
// Higher-order function to filter slice
func filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, n := range nums {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}

// Predicate functions
func isEven(n int) bool {
	return n%2 == 0
}

func isGreaterThanFive(n int) bool {
	return n > 5
}

func main() {
	fmt.Println("Sum:", apply(5, 3, add))          // 5 + 3 = 8
	fmt.Println("Product:", apply(5, 3, multiply)) // 5 * 3 = 15

	/*
			Explanation:
		The apply function takes two integers and another function operation that performs some calculation.
		The add and multiply functions are passed as arguments, making apply reusable for different operations.
	*/

	addFive := adder(5)      // Returns a function that adds 5
	fmt.Println(addFive(10)) // 5 + 10 = 15
	fmt.Println(addFive(20)) // 5 + 20 = 25

	/*
			Explanation:
		The adder function returns another function that adds a specific value (x) to a provided argument (y).
		addFive is a closure that remembers x = 5 and adds it to different numbers.
	*/

	adder := createOperation("add")
	multiplier := createOperation("multiply")

	fmt.Println(adder(4, 6))      // 4 + 6 = 10
	fmt.Println(multiplier(4, 6)) // 4 * 6 = 24

	/*
			Explanation:
		The createOperation function returns a different function based on the provided string (add or multiply).
		This allows dynamic function generation at runtime, making code highly adaptable.
	*/

	numbers := []int{1, 2, 3, 6, 7, 8, 9}

	evenNumbers := filter(numbers, isEven)
	fmt.Println("Even Numbers:", evenNumbers) // [2, 6, 8]

	greaterThanFive := filter(numbers, isGreaterThanFive)
	fmt.Println("Greater than 5:", greaterThanFive) // [6, 7, 8, 9]

	/*
			Explanation:
		The filter function takes a slice and a predicate function.
		It returns a slice of elements that satisfy the condition defined by the predicate.
		isEven and isGreaterThanFive are passed as predicates.
	*/
}

/*
Higher-Order Functions in Go
A higher-order function is a function that does one or both of the following:

Takes one or more functions as arguments.
Returns a function as its result.
Higher-order functions enable functional programming techniques, allowing for flexible, reusable, and concise code.

Why Use Higher-Order Functions?
Code Reusability – Perform the same operation with different logic by passing different functions.
Abstraction – Encapsulate behavior and pass it around as parameters.
Cleaner Code – Avoid repetitive loops and conditionals by abstracting common patterns

Benefits of Higher-Order Functions:
Flexibility – Easily swap out behavior by passing different functions.
Code Simplification – Replace loops and conditionals with concise functional expressions.
Reusability – Generic functions like filter can handle any logic by passing different predicates.
Functional Composition – Combine multiple functions for more complex behavior.

Common Use Cases:
Map, Filter, Reduce – Data transformations and aggregations.
Middleware in Web Frameworks – Custom request processing functions.
Event Handlers – Pass callback functions to handle events dynamically.
Functional Pipelines – Chain multiple functions to perform complex transformations.
*/
