package main

import "fmt"

/*
make is a built-in function in Go that allocates and initializes slices, maps, and channels.
It prepares these data structures for use by creating the underlying data structure and returning a reference to it.

Why Use make?
Prevents nil errors – When you declare slices, maps, or channels without make, they are nil by default.
Using them without initialization results in runtime errors.
Efficient memory allocation – make allocates memory efficiently, allowing you to control the size and capacity upfront.
Avoids repetitive initialization logic – With a single call to make,
you can allocate and initialize slices, maps, and channels without additional boilerplate code.

var m map[string]int
m["A"] = 1  // Causes a runtime panic (map is nil)

m = make(map[string]int)  // Now it's initialized
m["A"] = 1  // Works fine

Without make – The map is nil and assigning values to it results in a panic.
With make – The map is initialized and ready to accept key-value pairs
*/


// Use make to create and initialize a slice
func createSlice() []int {
	slice := make([]int, 5, 10) // Slice of length 5, capacity 10
	for i := 0; i < len(slice); i++ {
		slice[i] = i * 2
	}
	return slice
}

// Use make to create and initialize a map
func createMap() map[string]int {
	m := make(map[string]int) // Initialize empty map
	m["Alice"] = 85
	m["Bob"] = 92
	m["Charlie"] = 78
	return m
}

// Use make to create and initialize a buffered channel
func createChannel() chan string {
	ch := make(chan string, 2) // Buffered channel with capacity 2
	ch <- "Hello"
	ch <- "World"
	return ch
}

// Demonstrate working with unbuffered channel
func unbufferedChannelExample() chan int {
	ch := make(chan int) // Unbuffered channel
	go func() {
		ch <- 42 // Sending value in a goroutine
	}()
	return ch
}

// Create a 2D slice (slice of slices)
func create2DSlice(rows, cols int) [][]int {
	grid := make([][]int, rows) // Slice of slices (outer slice)

	for i := range grid {
		grid[i] = make([]int, cols) // Initialize each inner slice
		for j := 0; j < cols; j++ {
			grid[i][j] = i + j // Fill the 2D slice with values
		}
	}
	return grid
}

// Main function to call each use case
func main() {
	// Slice Example
	slice := createSlice()
	fmt.Println("Slice:", slice)

	// Map Example
	m := createMap()
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Buffered Channel Example
	ch := createChannel()
	fmt.Println(<-ch) // Receive first value
	fmt.Println(<-ch) // Receive second value

	// Unbuffered Channel Example
	unbufferedCh := unbufferedChannelExample()
	fmt.Println("Received from unbuffered channel:", <-unbufferedCh)

	// 2D Slice Example
	grid := create2DSlice(3, 4)
	fmt.Println("2D Slice:")
	for _, row := range grid {
		fmt.Println(row)
	}
}
