package main

import "fmt"

// Basic Function - Function without parameters and return
// This type of function performs an action but doesn't take any input or return any value.
func greet() {
	fmt.Println("Hello! Welcome to Golang.")
}

// Function with Parameters (No Return)
// This function accepts parameters but doesn't return any value.
// Parameters allow the function to work with different inputs.
func greetUser(name string) {
	fmt.Println("Hello,", name)
}

// Function returning a string - Function with Return (No Parameters)
// A function can return a value without accepting any parameters.
func getGreeting() string {
	return "Hello from Go!"
}

// Multiple Return Values
// Function returning multiple values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Named return values
// Return values can be named, which improves readability and
// allows you to return values without explicitly specifying them.
func rectangleDimensions(width, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return // Implicit return of named values
}

// Generic Function Example
// The Print[T any](value T) function can accept any data type due to the use of generics (T any).
func Print[T any](value T) {
	fmt.Println("Generic function called with value:", value)
}

// Function as Type
// Operation is a type that represents a function signature accepting two int parameters and returning an int.
type Operation func(int, int) int

// Function as Value - Function with Parameters and Return
// This is the most common type. The function takes input parameters and returns a value (or multiple values).
func add(a, b int) int {
	return a + b
}

// or func add(a int, b int) int {}

// Function Closure Example
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Higher-Order Function Example
func applyOperation(a, b int, op Operation) int {
	return op(a, b)
}

// Immediately Invoked Function (IIF)
func main() {
	// greet
	greet()
	greetUser("Aditya")
	getGreeting_obj := getGreeting()
	fmt.Println("Get Greeting: ", getGreeting_obj)
	// Generic function usage
	Print(42)
	Print("Hello, Generics!")
	Print(3.14)

	// Performing calculations or operations that require input and return a result
	addition := add(5, 7)
	fmt.Println("Sum:", addition)

	// Error handling, returning results along with metadata or status codes
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	// When functions are long, named returns make it easier to understand the purpose of return values
	a, p := rectangleDimensions(5, 4)
	fmt.Println("Area:", a, "Perimeter:", p)

	// Function as type and function as value
	var operation Operation = add
	fmt.Println("Addition using function as value:", operation(3, 4))

	// Closure usage
	multiplyByTwo := makeMultiplier(2)
	fmt.Println("Multiply by 2:", multiplyByTwo(5))

	// Higher-Order function usage
	result := applyOperation(5, 3, func(a, b int) int { return a * b })
	fmt.Println("Multiplication using higher-order function:", result)

	// Immediately Invoked Function (IIF)
	func() {
		fmt.Println("IIF: This function is immediately invoked!")
	}()

	// Variadic function example - Aggregation, concatenation, or processing of multiple inputs.
	printNumbers(1, 2, 3, 4, 5)

	// Method demonstration
	point := Point{X: 3, Y: 4}
	fmt.Println("Point Distance from Origin:", point.DistanceFromOrigin())

	// Anonymous Functions (Function Without a Name)
	// Anonymous functions are defined inline and can be assigned to variables or immediately invoked.
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("Product:", multiply(3, 4))

	// Immediately Invoked
	// Short-lived operations or closures that don't need to be reused elsewhere.
	immediate := func(x int) int {
		return x * x
	}(5) // Immediate invocation
	fmt.Println("Square:", immediate)
}

// Variadic Functions (Accepts Multiple Arguments)
// Variadic functions can accept any number of arguments of a specified type.
func printNumbers(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

// Method example
type Point struct {
	X, Y int
}

// Method to calculate distance from origin
func (p Point) DistanceFromOrigin() float64 {
	return float64(p.X*p.X + p.Y*p.Y)
}
