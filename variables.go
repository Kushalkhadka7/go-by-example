package main

import "fmt"

func main() {
	var x int32 = 10 // x is of type int32
	var y int64 = 10 // x is of type int64
	a := 10          // declaration and initialization

	const s string = "constant string" // constant variable

	var b int

	b = 20

	// To hold the unicode character go has a type called rune.
	var r rune = '~'

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(r)
}
