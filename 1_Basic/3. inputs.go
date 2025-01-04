package main

import (
	"fmt"
)

func main() {
	fmt.Println("Reading from Console")
	var first_name string
	var last_name string
	var age int
	fmt.Println("Enter your first name:")
	fmt.Scanln(&first_name)
	fmt.Println("Enter your last name:")
	fmt.Scanln(&last_name)
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)
	fmt.Printf("First name: %s, Last name: %s, Age: %d", first_name, last_name, age)
}
