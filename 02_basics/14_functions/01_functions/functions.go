package main

import "fmt"

func main() {

	// Llamado a la función
	sum := add(3, 4)
	fmt.Println("Sum:", sum)
	fmt.Println("Sum:", add(10, 20))

	// Funciones anonimas
	saludo := func() {
		fmt.Println("Hello from an anonymous function!")
	}
	saludo()

	// Funciones como tipos
	operation := add
	result := operation(5, 7)
	fmt.Println("Result from operation:", result)

	// Funciones como argumentos
	result2 := applyOperations(10, 20, add)
	fmt.Println("Result from applyOperations with add:", result2)

	result3 := applyOperations(10, 20, subtract)
	fmt.Println("Result from applyOperations with subtract:", result3)

	// Funciones que retornan funciones
	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Println("Double 5:", double(5))
	fmt.Println("Triple 5:", triple(5))
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func applyOperations(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Func that returns another function
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
