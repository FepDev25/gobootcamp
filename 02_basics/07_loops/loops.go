package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Simple loop
	for i := 0; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Iterate over collections
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Break and continue
	secretNumber := 17
	for i := range 21 {
		if i == secretNumber {
			fmt.Println("Found the secret number:", secretNumber)
			break // Stop the loop if the secret number is found
		}

		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println("Odd number:", i)
	}

	// Outer loop
	rows := 8
	for i := range rows {
		// Inner loop for spaces before stars
		for j := 1; j < rows-i; j++ {
			fmt.Print(" ")
		}
		// Inner loop for printing stars
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println() // Move to the next line
	}

	// For loop as while loop
	count := 1
	for count <= 10 {
		fmt.Println("Count:", count)
		count++
	}

	game()
}

func game() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	target := random.Intn(100) + 1

	var userNum int
	for userNum != target {
		fmt.Print("Enter your guess (1-100): ")
		fmt.Scanln(&userNum)

		if userNum < target {
			fmt.Println("Too low!")
		} else if userNum > target {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Congratulations! You've found the secret number:", target)
		}
	}
}
