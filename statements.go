package main

import "fmt"

func main() {
	isActive := false

	// If else
	if isActive {
		fmt.Println("active")
	} else {
		fmt.Println("inactive")
	}

	// If statement with inline declaration.
	if num := 10; num%2 == 0 {
		fmt.Println("even")
	}

	var x int = 100

	// If else if.
	if x < 10 {
		fmt.Println("smaller")
	} else if x > 10 {
		fmt.Println("greater")
	} else {
		fmt.Println("equal")
	}

	// switch statement.
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Print("three")
	default:
		fmt.Println("ten")
	}

	// Another switch.
	switch {
	case 1 < 10:
		fmt.Println("smaller")
	case 1 > 10:
		fmt.Println("Greater")
	default:
		fmt.Println("equal")
	}

	// [NOTE]: We can run switch statement based on type also.
}
