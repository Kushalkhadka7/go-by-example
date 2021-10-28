package main

import "fmt"

type Rectangle struct {
	length int
	breath int
}

func (r *Rectangle) Reset() {
	r.breath = 0
	r.length = 0
}

func main() {
	r := Rectangle{
		length: 10,
		breath: 10,
	}

	fmt.Println(r.length, r.breath)
	r.Reset()

	fmt.Println(r.length, r.breath)

	x := 20
	y := &x // Y is now the memory address of x.

	fmt.Println(x)
	fmt.Println(*y)
}
