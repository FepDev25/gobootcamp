package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 5
	var result int

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	const p float64 = 222.0 / 78.0
	fmt.Println("P:", p)

	result = a % b
	fmt.Println("Modulus:", result)

	// Exponentiation using math.Pow
	expResult := math.Pow(float64(a), float64(b))
	fmt.Println("Exponentiation:", expResult)

	// Overflow with signed integers
	var maxInt int64 = 1<<63 - 1
	fmt.Println("Max Int:", maxInt)
	maxInt++
	fmt.Println("Overflowed Max Int:", maxInt)

	// Overflow with unsigned integers
	var maxUint uint64 = 1<<64 - 1
	fmt.Println("Max Uint:", maxUint)
	maxUint++
	fmt.Println("Overflowed Max Uint:", maxUint)
}
