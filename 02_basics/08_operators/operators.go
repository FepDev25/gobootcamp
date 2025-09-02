package main

import "fmt"

func main() {
	// Logical operators
	// AND (&&), OR (||), NOT (!)

	a := true
	b := false
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// AND
	if a && b {
		fmt.Println("Both are true")
	} else {
		fmt.Println("At least one is false")
	}

	// OR
	if a || b {
		fmt.Println("At least one is true")
	} else {
		fmt.Println("Both are false")
	}

	// NOT
	if !a {
		fmt.Println("a is false")
	} else {
		fmt.Println("a is true")
	}

	// Bitwise operators
	// AND (&), OR (|), XOR (^), NOT (^), bit clear (&^), left shift (<<), right shift (>>)
	x := 5  // 0101
	y := 3  // 0011

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// AND
	fmt.Println("AND:", x&y) // 1

	// OR
	fmt.Println("OR:", x|y) // 7

	// XOR
	fmt.Println("XOR:", x^y) // 6

	// NOT
	fmt.Println("NOT:", ^x) // -6

	// Bit clear
	fmt.Println("BIT CLEAR:", x&^y) // 4

	// Left shift
	fmt.Println("LEFT SHIFT:", x<<1) // 10

	// Right shift
	fmt.Println("RIGHT SHIFT:", x>>1) // 2

	// Comparison operators
	// ==, !=, >, <, >=, <=
	fmt.Println("x == y:", x == y) // false
	fmt.Println("x != y:", x != y) // true
	fmt.Println("x > y:", x > y)   // true
	fmt.Println("x < y:", x < y)   // false
	fmt.Println("x >= y:", x >= y) // true
	fmt.Println("x <= y:", x <= y) // false
}
