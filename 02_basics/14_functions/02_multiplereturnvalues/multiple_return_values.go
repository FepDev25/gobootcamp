package main

import (
	"errors"
	"fmt"
)

func main() {
	// Multiple return values 1
	quotient, remainder := divide(10, 3)
	fmt.Println("10 divided by 3 is", quotient, "with a remainder of", remainder)

	quotient, remainder = divide(20, 5)
	fmt.Println("20 divided by 5 is", quotient, "with a remainder of", remainder)

	// Multiple return values 2
	result, err := compare(5, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	result, err = compare(2, 4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	result, err = compare(7, 7)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// Named return values
	res, err := divide2(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 divided by 2 is", res)
	}

	res, err = divide2(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 divided by 0 is", res)
	}
}

func divide(a, b int) (int, int) {
	cociente := a / b
	residuo := a % b
	return cociente, residuo
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if a < b {
		return "a is less than b", nil
	} else {
		return "", errors.New("Unable to compare, values are equal")
	}
}

func divide2(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return
	}
	result = a / b
	return
}
