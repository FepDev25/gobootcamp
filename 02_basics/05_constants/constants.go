package main

import "fmt"

const PI = 3.14        // Untyped
const GRAVITY = 9.81   // Untyped
const E float64 = 2.71 // Typed

func main() {
	fmt.Println("Pi:", PI)
	fmt.Println("Gravity:", GRAVITY)
	fmt.Println("E:", E)

	const (
		MONDAY    = 1
		TUESDAY   = 2
		WEDNESDAY = 3
		THURSDAY  = 4
		FRIDAY    = 5
	)

	fmt.Println("Monday:", MONDAY)
	fmt.Println("Tuesday:", TUESDAY)
	fmt.Println("Wednesday:", WEDNESDAY)
	fmt.Println("Thursday:", THURSDAY)
	fmt.Println("Friday:", FRIDAY)

}
