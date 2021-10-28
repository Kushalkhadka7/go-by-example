package main

import "fmt"

func main() {
	var empty interface{} // Creates an empty interface that can holds any value.

	empty = 10

	fmt.Println(empty)
}
