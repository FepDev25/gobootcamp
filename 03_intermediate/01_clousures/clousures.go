package main

import "fmt"

func main() {

	// Clousures are a concept that allows function to capture and
	// manipulate variables that are defined outside their body.

	// A clousure is a function value that references variables from
	// outside its body.

	sequence := adder() // i is initialized to 0 here

	// Each call to sequence() increments and returns the updated value of i

	fmt.Println("First call:", sequence())
	fmt.Println("Second call:", sequence())
	fmt.Println("Third call:", sequence())

}

func adder() func() int {
	i := 0
	fmt.Println("previus i:", i)

	return func() int {
		i++
		fmt.Println("added 1 to i:", i)
		return i
	}

}
