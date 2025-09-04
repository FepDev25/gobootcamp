package main

import (
	"fmt"
	"slices"
)

func main() {
	delclaraInicializarSlices()
	sliceFunctions()
	compareSlice()
	twoDimSlices()
	slicesPackage()
}

func delclaraInicializarSlices() {
	var numbers []int            // Declarar un slice vacío
	var numers2 = []int{1, 2, 3} // Declarar un slice con valores
	numbers3 := []int{1, 2, 3}   // Declarar y inicializar un slice
	numbers4 := make([]int, 10)  // Crear un slice con longitud 10

	fmt.Println((numbers))
	fmt.Println((numers2))
	fmt.Println((numbers3))
	fmt.Println((numbers4))

	array := [4]int{10, 20, 30, 40}
	slice := array[1:3] // Crear un slice a partir de un array
	fmt.Println(array)
	fmt.Println(slice)
}

func sliceFunctions() {
	slice := []int{10, 20, 30}
	fmt.Println(slice)

	slice = append(slice, 40)         // Agregar un elemento al slice
	slice = append(slice, 50, 60, 70) // Agregar múltiples elementos
	fmt.Println(slice)

	// Copiar un slice
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	fmt.Println(sliceCopy)

	// Nil slice
	var nilSlice []int
	fmt.Println(nilSlice)

	// Recorrer un slice
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// Acceder a elementos
	fmt.Println("First element:", slice[0])
	fmt.Println("Third element:", slice[2])
	fmt.Println("Last element:", slice[len(slice)-1])
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice))

	// Modificar un elemento
	slice[1] = 25
	fmt.Println("Modified slice:", slice)
}

func compareSlice() {
	// Compare slices
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6}

	fmt.Println("S1: ", s1)
	fmt.Println("S2: ", s2)
	fmt.Println("S3: ", s3)
	fmt.Println("s1 == s2:", slices.Equal(s1, s2))
	fmt.Println("s1 == s3:", slices.Equal(s1, s3))
}

func twoDimSlices() {
	// Crear un slice bidimensional usando un bucle for
	two := make([][]int, 3)
	count := 2
	for i := range two {
		two[i] = make([]int, 4)
		for j := range two[i] {
			two[i][j] = count
			count = count * 2
		}
	}
	fmt.Println("2D Slice:")
	fmt.Println(two)

	// Crear un slice bidimensional con valores iniciales
	two2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D Slice:")
	fmt.Println(two2)

	// Modificar slices internos
	two2[0] = append(two2[0], 10)
	two2[1] = append(two2[1], 15)
	two2[2] = append(two2[2], 20)
	fmt.Println(two2)

	// Modificar slice
	two2 = append(two2, []int{90, 100, 120, 500})
	fmt.Println(two2)
}

func slicesPackage() {
	numbers := []int{1, 10, 20, 60, 65, 72}
	names := []string{"Alice", "Bob", "Charlie"}

	fmt.Println("Original numbers:", numbers)
	fmt.Println("Original names:", names)

	// Usar funciones del paquete slices
	fmt.Println("Contains 20:", slices.Contains(numbers, 20))
	fmt.Println("Index of 60:", slices.Index(numbers, 60))
	fmt.Println("Min:", slices.Min(numbers))
	fmt.Println("Max:", slices.Max(numbers))

	slices.Sort(numbers)
	fmt.Println("Sorted:", numbers)
	fmt.Println("Is sorted:", slices.IsSorted(numbers))

	numbers = slices.Delete(numbers, 2, 4) // Eliminar elementos desde el índice 2 hasta el 4 (exclusivo)
	fmt.Println("After deletion:", numbers)

	numbers = slices.Insert(numbers, 2, 15) // Insertar elementos en el índice 2
	fmt.Println("After insertion:", numbers)

	names = slices.Replace(names, 1, 2, "Felipe", "Juan") // Reemplazar "Bob" con "Felipe" y "Charlie" con "Juan"
	fmt.Println("After replacement:", names)

	slices.Reverse(names) // Revertir el slice
	fmt.Println("After reversal:", names)
}
