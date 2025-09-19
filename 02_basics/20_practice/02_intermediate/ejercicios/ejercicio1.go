package ejercicios

import "slices"

/*
EJERCICIO 1: Manipulación Avanzada de Slices
===========================================
Implementa operaciones comunes en slices.

Instrucciones:
1. Recibe un slice de enteros
2. Implementa las siguientes operaciones (retorna múltiples valores):
  - Duplicados: slice con elementos duplicados removidos
  - Invertido: slice con elementos en orden inverso
  - Ordenado: slice ordenado de menor a mayor (implementa tu propio algoritmo)
  - Estadísticas: suma, promedio, mediana, moda

Ejemplo para slice = [3, 1, 4, 1, 5, 9, 2, 6, 5, 3]:
Duplicados removidos: [3 1 4 5 9 2 6]
Invertido: [3 5 6 2 9 5 1 4 1 3]
Ordenado: [1 1 2 3 3 4 5 5 6 9]
Suma: 39, Promedio: 3.90, Mediana: 3.50, Moda: [1 3 5]
*/

func Ejercicio1(numeros []int) ([]int, []int, []int, int, float64, float64, []int) {
	// TODO: Implementa aquí tu solución
	// Retorna: duplicados, invertido, ordenado, suma, promedio, mediana, moda

	// --- Eliminar duplicados ---
	set := make(map[int]struct{})
	var noDuplicates []int
	for _, v := range numeros {
		set[v] = struct{}{}
	}
	for k := range set {
		noDuplicates = append(noDuplicates, k)
	}

	// --- Invertido ---
	var invertido = make([]int, len(numeros))
	copy(invertido, numeros)
	slices.Reverse(invertido)

	// --- Ordenado ---
	var ordenado = bubbleSort(numeros)

	// --- Suma ---
	suma := 0
	for _, v := range numeros {
		suma += v
	}

	// --- Promedio ---
	promedio := float64(suma / len(numeros))

	// --- Mediana ---
	longitud := len(ordenado)
	var mediana float64
	if longitud%2 == 0 {
		medio1 := ordenado[longitud/2-1]
		medio2 := ordenado[longitud/2]
		mediana = float64(medio1+medio2) / 2.0
	} else {
		mediana = float64(ordenado[longitud/2])
	}

	// --- Moda ---
	frecuencias := make(map[int]int)
	maxFrecuencia := 0
	for _, v := range numeros {
		frecuencias[v]++
		if frecuencias[v] > maxFrecuencia {
			maxFrecuencia = frecuencias[v]
		}
	}
	var moda []int
	for num, freq := range frecuencias {
		if freq == maxFrecuencia {
			moda = append(moda, num)
		}
	}

	return noDuplicates, invertido, ordenado, suma, promedio, mediana, moda
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	sorted := make([]int, n)
	copy(sorted, arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	return sorted
}
