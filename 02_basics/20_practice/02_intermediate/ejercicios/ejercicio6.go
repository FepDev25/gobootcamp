package ejercicios

import (
	"fmt"
)

/*
EJERCICIO 6: Algoritmos de Búsqueda y Ordenamiento
==================================================
Implementa algoritmos clásicos de búsqueda y ordenamiento
para repasar a fondo estos conceptos.

Instrucciones:
1. Recibe un slice de enteros.
2. Implementa los siguientes algoritmos de búsqueda:
   - Búsqueda lineal: retorna índice del elemento o -1.
   - Búsqueda binaria: retorna índice del elemento o -1 (requiere slice ordenado).
   - (Opcional) Búsqueda por interpolación.
   - (Opcional) Búsqueda exponencial.

3. Implementa los siguientes algoritmos de ordenamiento:
   - Burbuja (Bubble Sort): retorna slice ordenado y número de intercambios.
   - Selección (Selection Sort): retorna slice ordenado y número de comparaciones.
   - Inserción (Insertion Sort).
   - Merge Sort (divide y vencerás).
   - Quick Sort (divide y vencerás).
   - Heap Sort (usa un heap máximo).
   - Counting Sort (no comparativo, útil para enteros pequeños).
   - (Opcional) Radix Sort y Bucket Sort.

4. Cada función debe retornar múltiples valores según se especifica
   (índice, booleano, intercambios, comparaciones, etc.).

5. Demuestra cada algoritmo con un ejemplo práctico.

Ejemplo de uso con slice = [64, 34, 25, 12, 22, 11, 90]:

- Búsqueda lineal de 22: índice = 4, encontrado = true
- Ordenamiento burbuja: [11 12 22 25 34 64 90], intercambios = 16
- Búsqueda binaria de 25 en slice ordenado: índice = 3, encontrado = true
- QuickSort: [11 12 22 25 34 64 90]
- MergeSort: [11 12 22 25 34 64 90]
- HeapSort: [11 12 22 25 34 64 90]
*/

func Ejercicio6(numeros []int, elemento int) {
	fmt.Println("Slice original:", numeros)

	// --- BÚSQUEDA ---
	idx, found := busquedaLineal(numeros, elemento)
	fmt.Printf("Búsqueda lineal de %d: índice=%d, encontrado=%v\n", elemento, idx, found)

	ordenado, _ := ordenamientoBurbuja(append([]int(nil), numeros...))
	idx, found = busquedaBinaria(ordenado, elemento)
	fmt.Printf("Búsqueda binaria de %d: índice=%d, encontrado=%v\n", elemento, idx, found)

	// --- ORDENAMIENTOS ---
	ordenB, swaps := ordenamientoBurbuja(append([]int(nil), numeros...))
	fmt.Println("Burbuja:", ordenB, "swaps:", swaps)

	ordenS, comps := ordenamientoSeleccion(append([]int(nil), numeros...))
	fmt.Println("Selección:", ordenS, "comparaciones:", comps)

	ordenI := ordenamientoInsercion(append([]int(nil), numeros...))
	fmt.Println("Inserción:", ordenI)

	ordenM := mergeSort(append([]int(nil), numeros...))
	fmt.Println("MergeSort:", ordenM)

	ordenQ := quickSort(append([]int(nil), numeros...))
	fmt.Println("QuickSort:", ordenQ)

	ordenH := heapSort(append([]int(nil), numeros...))
	fmt.Println("HeapSort:", ordenH)

	ordenC := countingSort(append([]int(nil), numeros...))
	fmt.Println("CountingSort:", ordenC)
}

// ==================== BÚSQUEDAS ====================

func busquedaLineal(numeros []int, elemento int) (int, bool) {
	for i, v := range numeros {
		if v == elemento {
			return i, true
		}
	}
	return -1, false
}

func busquedaBinaria(numeros []int, elemento int) (int, bool) {
	low, high := 0, len(numeros)-1
	for low <= high {
		mid := (low + high) / 2
		if numeros[mid] == elemento {
			return mid, true
		} else if numeros[mid] < elemento {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, false
}

// ==================== ORDENAMIENTOS ====================

func ordenamientoBurbuja(numeros []int) ([]int, int) {
	swaps := 0
	n := len(numeros)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numeros[j] > numeros[j+1] {
				numeros[j], numeros[j+1] = numeros[j+1], numeros[j]
				swaps++
			}
		}
	}
	return numeros, swaps
}

func ordenamientoSeleccion(numeros []int) ([]int, int) {
	comparaciones := 0
	n := len(numeros)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			comparaciones++
			if numeros[j] < numeros[minIdx] {
				minIdx = j
			}
		}
		numeros[i], numeros[minIdx] = numeros[minIdx], numeros[i]
	}
	return numeros, comparaciones
}

func ordenamientoInsercion(numeros []int) []int {
	for i := 1; i < len(numeros); i++ {
		key := numeros[i]
		j := i - 1
		for j >= 0 && numeros[j] > key {
			numeros[j+1] = numeros[j]
			j--
		}
		numeros[j+1] = key
	}
	return numeros
}

func mergeSort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}
	mid := len(numeros) / 2
	left := mergeSort(numeros[:mid])
	right := mergeSort(numeros[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func quickSort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}
	pivot := numeros[len(numeros)/2]
	less, equal, greater := []int{}, []int{}, []int{}
	for _, v := range numeros {
		if v < pivot {
			less = append(less, v)
		} else if v == pivot {
			equal = append(equal, v)
		} else {
			greater = append(greater, v)
		}
	}
	less = quickSort(less)
	greater = quickSort(greater)
	return append(append(less, equal...), greater...)
}

func heapSort(numeros []int) []int {
	n := len(numeros)
	// construir heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(numeros, n, i)
	}
	for i := n - 1; i > 0; i-- {
		numeros[0], numeros[i] = numeros[i], numeros[0]
		heapify(numeros, i, 0)
	}
	return numeros
}

func heapify(arr []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func countingSort(numeros []int) []int {
	if len(numeros) == 0 {
		return numeros
	}
	maxVal := numeros[0]
	for _, v := range numeros {
		if v > maxVal {
			maxVal = v
		}
	}
	count := make([]int, maxVal+1)
	for _, v := range numeros {
		count[v]++
	}
	result := []int{}
	for i, c := range count {
		for j := 0; j < c; j++ {
			result = append(result, i)
		}
	}
	return result
}
