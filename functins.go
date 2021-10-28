package main

import "fmt"

func main() {
	hello, world, a := HelloWorld("hello", "world", "1243") // If the function have multiple return values we can receive like this.

	fmt.Println(hello + world + a)

	testFunc(1, 2, 3, 4)
}

// HelloWorld is a function that takes 2 params and have 2 return values.
func HelloWorld(a, b, c string) (string, string, string) {
	defer testDeferFunction() // This line will be always called when the return statement below is executed.
	return a, b, c
}

func testDeferFunction() {
	fmt.Println("I am defer func and i am called.")
}

// Variadic Functions
// In Variadic functions all the passed params are received as array.
// so here num is array of int, even if only one element is passed while calling the function.
func testFunc(num ...int) {
	for _, v := range num {
		fmt.Println(v)
	}
}
