package main

import "fmt"

type Direction int

func main() {
	const (
		a = iota
		b
		c
		d
	)

	const (
		North Direction = iota
		East
		South
		West
	)

	// Here north will be 0, east=1,south=2,west=3
	fmt.Println(North, East, South, West)

	// I will assign sequential number to the variables starting form a=0.

	fmt.Println(a, b, c, d)
}
