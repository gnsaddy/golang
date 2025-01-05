package main

import "fmt"

// BasicForLoop demonstrates the classic "for" loop with initialization, condition, and post statement.
func BasicForLoop() {
	fmt.Println("Basic For Loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i) // Print numbers from 1 to 5
	}
}

// InfiniteLoop demonstrates an infinite loop with a break condition.
func InfiniteLoop() {
	fmt.Println("\nInfinite Loop with Break:")
	count := 0
	for {
		count++
		fmt.Println(count)
		if count == 5 {
			break // Exit the loop when count reaches 5
		}
	}
}

// ConditionBasedLoop demonstrates a loop based on a condition (like while loop in other languages).
func ConditionBasedLoop() {
	fmt.Println("\nCondition Based Loop (like While):")
	count := 1
	for count <= 5 {
		fmt.Println(count)
		count++
	}
}

// MultipleConditionsLoop demonstrates a loop with multiple conditions in the initialization and post statement.
func MultipleConditionsLoop() {
	fmt.Println("\nMultiple Conditions Loop:")
	for i, j := 0, 10; i <= 5; i, j = i+1, j-1 {
		fmt.Println("i =", i, "j =", j) // Prints the values of i and j
	}
}

// LoopOverArray demonstrates a loop iterating over an array.
func LoopOverArray() {
	fmt.Println("\nLoop Over Array:")
	numbers := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index", i, ":", numbers[i]) // Prints each element of the array
	}
}

// LoopOverSlice demonstrates a loop iterating over a slice.
func LoopOverSlice() {
	fmt.Println("\nLoop Over Slice:")
	fruits := []string{"Apple", "Banana", "Cherry"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Fruit", i, ":", fruits[i]) // Prints each element of the slice
	}
}

// LoopOverMap demonstrates a loop iterating over a map.
func LoopOverMap() {
	fmt.Println("\nLoop Over Map:")
	scores := map[string]int{
		"John": 90,
		"Emma": 85,
		"Alex": 88,
	}
	for name, score := range scores {
		fmt.Println(name, "scored", score) // Prints each key-value pair of the map
	}
}

// RangeBasedLoop demonstrates a loop using the `range` keyword.
func RangeBasedLoop() {
	fmt.Println("\nRange-Based Loop:")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Println("Index", index, "Value", value) // Prints each index and value of the slice
	}
}

func RangeBasedLoopWithMap() {
	scores := map[string]int{
		"John": 90,
		"Emma": 85,
		"Alex": 88,
	}
	for name, score := range scores {
		fmt.Println(name, "scored", score) // Prints each key-value pair in the map
	}
}

func main() {
	// Calling all functions to demonstrate different loop types
	BasicForLoop()
	InfiniteLoop()
	ConditionBasedLoop()
	MultipleConditionsLoop()
	LoopOverArray()
	LoopOverSlice()
	LoopOverMap()
	RangeBasedLoop()
	RangeBasedLoopWithMap()
}
