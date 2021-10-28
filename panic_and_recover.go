package main

import "fmt"

func main() {
	testPanicAndRecover()
}

func testPanicAndRecover() {
	// This function will run at end fo the function testPanicAndRecover,
	// so once panic is occurred this function will catch the paniced value and will deal with it,
	// not giving the program to crash.
	// We can also take panic and recover as try and catch in other languages.
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovering from panic in printAllOperations error is: %v \n", r)
		}
	}()

	isActive := true

	// This will create the program to panic.
	if isActive {
		panic("Paniced")
	}
}
