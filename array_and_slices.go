package main

import "fmt"

func main() {
	// Array
	var a [5]int // Creates an array of type int and length 5 filled with 0.
	a[0] = 10
	fmt.Println(a)
	// [NOTE]: Array are mutable if we add or remove elements int the array it will effect the original array.

	// Slices

	var s = make([]string, 3) // Creates an empty slice of string of length 3.

	s[0] = "hello"
	s[1] = "world"
	s[2] = "123"

	fmt.Println(s)

	// Methods associated with array and slices.

	l := len(s) // Length of the array or slice.

	t := make([]string, 3)

	copy(t, s) // Copy the elements of array s to t.

	s[1] = "wrld"

	fmt.Printf("length: %v \n", l)
	fmt.Printf("t: %v \n", t)

	firstTwo := s[:2] // get first two elements of the array
	all := s[:]       // get all elements in the array.

	// [NOTE]: Now changing the elements of all array will also change the elements of array s.
	// If we want to make the copy of the array then we have to use `copy(destination, source)`.

	fmt.Printf("first tow: %v \n", firstTwo)
	fmt.Printf("all: %v \n", all)
}
