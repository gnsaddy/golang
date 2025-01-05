package main

import "fmt"

func main() {
	fmt.Println("Learning Variables")
	variables()
	other_vars()
	VariableDeclarationDemo()
}

func variables() {
	var first_name string = "Aditya"
	var last_name string = "Raj"
	age := 27
	// constant variable
	const vote_age = 18

	fmt.Println("First name is: ", first_name)
	fmt.Println("Last name is: ", last_name)
	fmt.Println("Age is: ", age)
	fmt.Println("Vote age is: ", vote_age)
}

func other_vars() {
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

func VariableDeclarationDemo() {
	// Explicit declaration with type
	var name string
	var age int
	var isDeveloper bool

	// Assign values
	name = "John Doe"
	age = 30
	isDeveloper = true

	// Short declaration (type inferred)
	city := "Bangalore"
	salary := 80000
	isMarried := false

	// Multiple variable declaration
	var x, y, z int = 10, 20, 30
	a, b, c := 1.1, 2.2, 3.3

	// Print values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Developer:", isDeveloper)
	fmt.Println("City:", city)
	fmt.Println("Salary:", salary)
	fmt.Println("Is Married:", isMarried)
	fmt.Println("x, y, z:", x, y, z)
	fmt.Println("a, b, c:", a, b, c)
}
