package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Starting to execute....")

	wg.Add(2) // This will tell wait group that there will be 2 go routines running.

	// This go routine will be executed in background.
	go firstGoRoutine()

	fmt.Println("Starting to execute1....")

	// This go routine will be executed in background.
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Hello from second go routine.")
		wg.Done() // Once the go routine is done it will notify the wait group that it has been completed so no need to wait for it now.
	}()

	fmt.Println("Completed main function execution.")

	// This loop prevents from the main go routine to exit.
	// In another example, in place of this we will be using wait groups.
	wg.Wait() // It will prevnet the main go routine to exit, make it to wait until all the go routines notifies done.
}

func firstGoRoutine() {
	time.Sleep(time.Second * 5)
	fmt.Println("Hello from first go routine.")
	wg.Done() // Once the go routine is done it will notify the wait group that it has been completed so no need to wait for it now.
}
