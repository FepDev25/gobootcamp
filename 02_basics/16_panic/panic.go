package main

import "fmt"

func main() {
	// In go, a panic is a built-in function that stops the normal execution of a program.
	// It stops the execution of the current function and begins panicking, and then executes any deferred functions.
	//process(10)
	process(-5) // This will cause a panic
	fmt.Println("This line will not be executed due to panic")
}

func process(input int) {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")

	if input < 0 {
		panic("Negative input not allowed")
	}
	fmt.Println("Processing:", input)
}
