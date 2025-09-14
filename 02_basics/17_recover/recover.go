package main

import "fmt"

func main() {
	// Recover is a built in function that is used to regain of panicking go rutine
	// It is only useful inside deferred functions.
	// Allow the program to continue its execution after a panic.
	// Returned a value but is only pass if there is a panic.
	process()
	fmt.Println("Program continues after recover")
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Start processing")
	panic("Something went wrong")
}
