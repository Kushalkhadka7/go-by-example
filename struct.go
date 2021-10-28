package main

import "fmt"

// Rectangle is a struct.
type Rectangle struct {
	length int
	breath int
}

// Area is a method associated with Rectangle.
func (r *Rectangle) Area() int {
	return r.breath * r.length
}

func main() {
	r := Rectangle{
		length: 10,
		breath: 10,
	}

	fmt.Println(r.Area())

	// Inline struct
	s := struct {
		length int
		breath int
	}{

		length: 20,
		breath: 20,
	}

	fmt.Println(s.length, s.breath)
}
