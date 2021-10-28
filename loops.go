package main

import "fmt"

func main() {
	var data []string = []string{"Hello", "world", "123", "456"}

	// For range.
	for i, v := range data { // Here i is the index of the element and v is the actual element.
		fmt.Printf("I:%v, V:%v \n", i, v)
	}

	// Normal for loop.
	for i := 0; i < len(data); i++ {
		fmt.Printf("I:%v, V:%v \n", i, data[i])
	}
}
