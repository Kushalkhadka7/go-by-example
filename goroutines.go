package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting to execute....")

	// This go routine will be executed in background.
	go firstGoRoutine()

	fmt.Println("Starting to execute1....")

	// This go routine will be executed in background.
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Hello from second go routine.")
	}()

	fmt.Println("Completed main function execution.")

	// This loop prevents from the main go routine to exit.
	// In another example, in place of this we will be using wait groups.
	for {

	}
}

func firstGoRoutine() {
	time.Sleep(time.Second * 5)
	fmt.Println("Hello from first go routine.")
}
