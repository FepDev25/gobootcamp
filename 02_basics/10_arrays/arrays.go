package main

import "fmt"

func main() {
	arraysIntroduction()
	iterateArrays()
	compareArrays()
	multiDimensionalArrays()
	copyArraysWithPointers()
}

func arraysIntroduction() {
	// Arrays
	var numbers [6]int
	fmt.Println(numbers)

	numbers[0] = 42
	numbers[1] = 27
	numbers[2] = 33
	numbers[3] = 19
	numbers[4] = 55
	numbers[5] = 78
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Grapes", "Orange"}
	fmt.Println(fruits)

	// Acceder a elementos
	fmt.Println(fruits[2]) // Grapes

	// Modificar elementos
	fruits[1] = "Banana (Updated)"
	fmt.Println(fruits)
	fmt.Println(fruits[1])

	// Copy arrays
	twoM := [5]int{2, 4, 6, 8}
	twoMCopy := twoM
	twoMCopy[4] = 10
	fmt.Println(twoM)     // [2 4 6 8 0]
	fmt.Println(twoMCopy) // [2 4 6 8 10]

	// Longitud del array
	fmt.Println("Length of fruits array:", len(fruits))
}

func iterateArrays() {
	fruits := [4]string{"Apple", "Banana", "Grapes", "Orange"}

	// Recorrer arreglos
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	for _, fruit := range fruits { // Guion bajo significa "ignorar el valor"
		fmt.Println(fruit)
	}

	for index := range fruits {
		fmt.Println(fruits[index])
	}
}

func compareArrays() {
	// Compare arrays
	array := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	array3 := [3]int{4, 5, 6}

	fmt.Println("Array 1 == Array 2:", array == array2)
	fmt.Println("Array 1 == Array 3:", array == array3)
}

func multiDimensionalArrays() {
	// Multi-dimensional arrays
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:", matrix)
	fmt.Println("Matrix length:", len(matrix))
	fmt.Println("First row length:", len(matrix[0]))
	firstRow := matrix[0]
	fmt.Println("First row:", firstRow)
	fmt.Println("Element at (2,3):", matrix[1][2]) // 6
}

func copyArraysWithPointers() {
	// Copy arrays and references the original array
	original := [3]int{1, 5, 10} // Arreglo con 3 enteros

	var copiedArray *[3]int // Direccion de memoria de un arreglo con 3 enteros
	copiedArray = &original // Copiando la direccion de memoria del arreglo original

	fmt.Println("Original array:", original)   // [1 5 10]
	fmt.Println("Copied array:", *copiedArray) // [1 5 10]

	copiedArray[0] = 100 // Modificando el primer elemento del arreglo a traves de la referencia

	fmt.Println("Original array:", original)   // [100 5 10]
	fmt.Println("Copied array:", *copiedArray) // [100 5 10]
}
