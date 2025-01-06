package main

import "fmt"

/*
for index, value := range array/slice {
    //Do something with index and value
}

 for-range loop works in case of array/slice.
 It iterates over the given array/slice starting from index zero and the body of the for range loop
 is executed for every value present at the index.
 Both index and value are optional in for-range when using with array/slice.
*/

func rangeFunction() {
	fmt.Println("Range Loop in go")
	var letters = []string{"a", "e", "i", "o", "u"}

	// with index and value
	fmt.Println("Print Index and value")
	for i, letter := range letters {
		fmt.Printf("Index: %d, value: %s\n", i, letter)
	}

	//Only value
	fmt.Println("\nOnly value")
	for _, letter := range letters {
		fmt.Printf("Value: %s\n", letter)
	}

	//Only index
	fmt.Println("\nOnly Index")
	for i := range letters {
		fmt.Printf("Index: %d\n", i)
	}

	//Without index and value. Just print array values
	fmt.Println("\nWithout Index and Value")
	i := 0
	for range letters {
		fmt.Printf("Index: %d Value: %s\n", i, letters[i])
		i++
	}
}

/*
for-range loop with a map
In case of map for-range iterates over key and values of a map. Below is the format for for-range when using with a map.
for key, value := range map {
    //Do something with key and value
}
	A point to be noted that both key and value are optional to be used while using for-range with maps.
*/

func rangeMap() {
	sample := map[string]string{
		"a": "x",
		"b": "y",
	}

	//Iterating over all keys and values
	fmt.Println("Both Key and Value")
	for k, v := range sample {
		fmt.Printf("key :%s value: %s\n", k, v)
	}

	//Iterating over only keys
	fmt.Println("\nOnly keys")
	for k := range sample {
		fmt.Printf("key :%s\n", k)
	}

	//Iterating over only values
	fmt.Println("\nOnly values")
	for _, v := range sample {
		fmt.Printf("value :%s\n", v)
	}
}

/*
for-range loop with a channel
for-range loop works differently too for a channel.
For a channel, an index doesn't make any sense as the channel is similar to a pipeline
where values enter from one and exit from the other end. So in case of channel, the for-range loop will
iterate over values currently present in the channel.
After it has iterated over all the values currently present (if any),
the for-range loop will not exit but instead wait for next value that might be
pushed to the channel and it will exit only when the channel is closed

for value := range channel {
    //Do something value
}
*/

func rangeWithMake() {
	ch := make(chan int)

	go func() {
		ch <- 100
		ch <- 200
		close(ch)
	}()

	for value := range ch {
		fmt.Println("Channel Values: ", value)
	}
}

func main() {
	rangeFunction()
	rangeMap()
	rangeWithMake()

	var fruits map[string]string
	fruits = map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Cherry",
	}

	for key, value := range fruits {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
