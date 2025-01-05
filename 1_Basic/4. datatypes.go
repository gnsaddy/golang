package main

import (
	"fmt"
)

func main() {
	intTypesDemo()
	uintTypesDemo()
	floatTypesDemo()
	complexTypesDemo()
	byteRuneDemo()
	stringBoolDemo()
	arrayStructDemo()
	sliceMapDemo()
	channelDemo()
	pointerDemo()
	interfaceDemo()
}

// Demonstrates various signed integer types in Go
func intTypesDemo() {
	var intVar int = 42
	var int8Var int8 = 8
	var int16Var int16 = 16
	var int32Var int32 = 32
	var int64Var int64 = 64

	fmt.Println("Integer Data Types:")
	fmt.Printf("int: %d\n", intVar)
	fmt.Printf("int8: %d\n", int8Var)
	fmt.Printf("int16: %d\n", int16Var)
	fmt.Printf("int32: %d\n", int32Var)
	fmt.Printf("int64: %d\n", int64Var)
}

// Demonstrates various unsigned integer types and uintptr in Go
func uintTypesDemo() {
	var uintVar uint = 42
	var uint8Var uint8 = 8
	var uint16Var uint16 = 16
	var uint32Var uint32 = 32
	var uint64Var uint64 = 64
	var uintptrVar uintptr = uintptr(uint64Var)

	fmt.Println("Unsigned Integer Data Types:")
	fmt.Printf("uint: %d\n", uintVar)
	fmt.Printf("uint8: %d\n", uint8Var)
	fmt.Printf("uint16: %d\n", uint16Var)
	fmt.Printf("uint32: %d\n", uint32Var)
	fmt.Printf("uint64: %d\n", uint64Var)
	fmt.Printf("uintptr: %d\n", uintptrVar)
}

// Demonstrates floating-point types in Go
func floatTypesDemo() {
	var float32Var float32 = 3.14
	var float64Var float64 = 2.718

	fmt.Println("Float Data Types:")
	fmt.Printf("float32: %.2f\n", float32Var)
	fmt.Printf("float64: %.2f\n", float64Var)
}

// Demonstrates complex number types in Go
func complexTypesDemo() {
	var complex64Var complex64 = 1 + 2i
	var complex128Var complex128 = 2 + 3i

	fmt.Println("Complex Number Data Types:")
	fmt.Printf("complex64: %v\n", complex64Var)
	fmt.Printf("complex128: %v\n", complex128Var)
}

// Demonstrates byte and rune types in Go
func byteRuneDemo() {
	var byteVar byte = 'A'
	var runeVar rune = 'Ω'

	fmt.Println("Byte and Rune Data Types:")
	fmt.Printf("byte: %c (%d)\n", byteVar, byteVar)
	fmt.Printf("rune: %c (%d)\n", runeVar, runeVar)
}

// Demonstrates string and boolean types in Go
func stringBoolDemo() {
	var stringVar string = "Hello, Go!"
	var boolVar bool = true

	fmt.Println("String and Boolean Data Types:")
	fmt.Printf("string: %s\n", stringVar)
	fmt.Printf("bool: %t\n", boolVar)
}

// Demonstrates array and struct types in Go
func arrayStructDemo() {
	arrayVar := [3]int{1, 2, 3}
	type Person struct {
		Name string
		Age  int
	}
	personVar := Person{Name: "John Doe", Age: 30}

	fmt.Println("Array and Struct Data Types:")
	fmt.Printf("array: %v\n", arrayVar)
	fmt.Printf("struct: %+v\n", personVar)
}

// Demonstrates slice and map types in Go
func sliceMapDemo() {
	sliceVar := []string{"Go", "Python", "Java"}
	mapVar := map[string]int{"one": 1, "two": 2}

	fmt.Println("Slice and Map Data Types:")
	fmt.Printf("slice: %v\n", sliceVar)
	fmt.Printf("map: %v\n", mapVar)
}

// Demonstrates channel type in Go
func channelDemo() {
	ch := make(chan int, 1)
	ch <- 42
	fmt.Println("Channel Data Type:")
	fmt.Printf("channel: %d\n", <-ch)
}

// Demonstrates pointer type in Go
func pointerDemo() {
	var intVar int = 42
	var pointerVar *int = &intVar

	fmt.Println("Pointer Data Type:")
	fmt.Printf("pointer address: %p, value: %d\n", pointerVar, *pointerVar)
}

// Demonstrates interface and type assertion in Go
func interfaceDemo() {
	var anyVar interface{} = "Dynamic Type"
	fmt.Println("Interface Data Type:")
	fmt.Printf("interface holds: %v\n", anyVar)

	str, ok := anyVar.(string)
	if ok {
		fmt.Println("Type assertion to string successful:", str)
	} else {
		fmt.Println("Type assertion failed")
	}
}
