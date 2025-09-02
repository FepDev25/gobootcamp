package main

import "fmt"

type Employee struct {
	ID        int
	firstName string
	lastName  string
}

func main() {
	// PascalCase
	// Structs, interfaces, enum
	// Ej: CalculateArea, UserInfo

	// snake_case
	// Ej: user_id, first_name

	// UPPERCASE
	// Constantes
	// Ej: PI, MAX_SIZE
	const MAXRETRIES = 5

	// mixedCase
	// Variables, identificadores
	// Ej: userID, firstName
	var employeeId = 123
	fmt.Println(employeeId)
}
