package main

import "fmt"

func main() {
	x := testClosureFunction()

	fmt.Println(x(10))
	fmt.Println(x(100))
	fmt.Println(x(1000))

	y := testClosureFunction1("Hello")
	fmt.Println(y("good morning"))
	fmt.Println(y("good evening"))
}

// Closure
func testClosureFunction() func(a int) int {
	i := 10

	return func(a int) int {
		return i + a
	}

}

// Closure
func testClosureFunction1(data string) func(a string) string {
	greeting := data

	return func(a string) string {
		return greeting + " " + a
	}
}
