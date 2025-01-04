package main

import "fmt"

func main() {
	fmt.Println("Learning Variables")
	variables()
}

func variables() {
	var first_name string = "Aditya"
	var last_name string = "Raj"
	age := 27
	const vote_age = 18

	fmt.Println("First name is: ", first_name)
	fmt.Println("Last name is: ", last_name)
	fmt.Println("Age is: ", age)
	fmt.Println("Vote age is: ", vote_age)
}
