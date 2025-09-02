package main

import "fmt"

func main() {
	ifElseStatement()
	switchStatement()
}

func ifElseStatement() {
	// If else
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	examScore := 71
	if examScore >= 90 && examScore <= 100 {
		fmt.Println("Excellent")
	} else if examScore >= 70 && examScore < 90 {
		fmt.Println("Good")
	} else if examScore < 70 && examScore >= 0 {
		fmt.Println("Needs Improvement")
	} else {
		fmt.Println("Invalid Score")
	}

	num := 15
	if num%2 == 0 {
		if num%3 == 0 {
			fmt.Println("Divisible by 2 and 3")
		} else {
			fmt.Println("Divisible by 2 but not by 3")
		}
	} else {
		fmt.Println("Not divisible by 2")
	}

	if num%2 == 0 && num%3 == 0 {
		fmt.Println("Divisible by 2 and 3")
	} else if num%2 == 0 || num%3 == 0 {
		fmt.Println("Divisible by 2 or 3")
	} else {
		fmt.Println("Not divisible by 2 or 3")
	}
}

func switchStatement() {
	// Switch case

	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's a banana")
	default:
		fmt.Println("Unknown fruit")
	}

	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Tuesday":
		fmt.Println("Second day of the work week")
	case "Wednesday":
		fmt.Println("Midweek day")
	case "Thursday":
		fmt.Println("Fourth day of the work week")
	case "Friday":
		fmt.Println("End of the work week")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Unknown day")
	}

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Unknown day")
	}

	number := 15
	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10:
		fmt.Println("Number is greater or equal to 10")
	default:
		fmt.Println("Unknown number")
	}

	num := 2
	switch {
	case num > 1:
		fmt.Println("Number is greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is equal to 2")
	default:
		fmt.Println("Number is not two")
	}

	checkType(42)
	checkType("Hello")
	checkType(3.14)
	checkType(true)
}

func checkType(x any) {
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is a string")
	case float64:
		fmt.Println("x is a float64")
	default:
		fmt.Println("Unknown type")
	}
}
