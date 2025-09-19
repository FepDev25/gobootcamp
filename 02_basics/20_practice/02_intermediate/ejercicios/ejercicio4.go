package ejercicios

import "fmt"

/*
EJERCICIO 4: Calculadora de Matriz
=================================
Operaciones básicas con matrices (slices bidimensionales).

Instrucciones:
1. Recibe dos matrices como [][]int
2. Implementa las siguientes operaciones:
  - Suma de matrices (si son del mismo tamaño)
  - Resta de matrices (si son del mismo tamaño)
  - Multiplicación de matrices (si son compatibles)
  - Transpuesta de una matriz

3. Retorna el resultado y un boolean indicando si la operación fue exitosa

Ejemplo para matrices A=[[1,2],[3,4]] y B=[[5,6],[7,8]]:
Suma: [[6,8],[10,12]], exitoso: true
Resta: [[-4,-4],[-4,-4]], exitoso: true
Multiplicación: [[19,22],[43,50]], exitoso: true
Transpuesta de A: [[1,3],[2,4]]
*/
func Ejercicio4(matrizA, matrizB [][]int, operacion string) ([][]int, bool) {
	// TODO: Implementa aquí tu solución
	resultado, realizado := make([][]int, 0), false

	switch operacion {
	case "suma":
		resultado, realizado = sumarMatrices(matrizA, matrizB)
	case "resta":
		resultado, realizado = restarMatrices(matrizA, matrizB)
	case "multi":
		resultado, realizado = multiplicarMatrices(matrizA, matrizB)
	case "trans":
		resultado, realizado = transpuestaMatriz(matrizA)
	default:
		fmt.Println("Operacion no soportada.")
	}

	return resultado, realizado
}

func validacionSumaYResta(matrizA, matrizB [][]int) (int, int, bool) {
	rowsA := len(matrizA)
	columnsA := len(matrizA[0])

	rowsB := len(matrizB)
	columnsB := len(matrizB[0])
	if rowsA != rowsB || columnsA != columnsB {
		fmt.Println("Las matrices deben tener las mismas dimensiones para poder operar.")
		return 0, 0, false
	}
	return rowsA, columnsA, true
}

func sumarMatrices(matrizA, matrizB [][]int) ([][]int, bool) {
	rows, columns, validate := validacionSumaYResta(matrizA, matrizB)

	if !validate {
		return nil, false
	}

	resultado := make([][]int, rows)
	for i := range rows {
		resultado[i] = make([]int, columns)
		for j := range columns {
			resultado[i][j] = matrizA[i][j] + matrizB[i][j]
		}
	}

	return resultado, true
}

func restarMatrices(matrizA, matrizB [][]int) ([][]int, bool) {
	rows, columns, validate := validacionSumaYResta(matrizA, matrizB)

	if !validate {
		return nil, false
	}

	resultado := make([][]int, rows)
	for i := range rows {
		resultado[i] = make([]int, columns)
		for j := range columns {
			resultado[i][j] = matrizA[i][j] - matrizB[i][j]
		}
	}

	return resultado, true
}

func validacionMultiplicacion(matrizA, matrizB [][]int) (int, int, bool) {
	rowsA := len(matrizA)
	columnsA := len(matrizA[0])

	rowsB := len(matrizB)
	columnsB := len(matrizB[0])

	if columnsA == rowsB {
		return rowsA, columnsB, true
	}

	fmt.Println("Para poder multiplicar el numero de columnas de la primera matriz debe ser igual al numero de filas de la segunnda.")
	return 0, 0, false
}

func multiplicarMatrices(matrizA, matrizB [][]int) ([][]int, bool) {
	rowsA, columnsB, validate := validacionMultiplicacion(matrizA, matrizB)

	if !validate {
		return nil, false
	}

	resultado := make([][]int, rowsA)

	for i := range rowsA {
		resultado[i] = make([]int, columnsB)
		for j := range columnsB {
			fila := matrizA[i]
			columna := obtenerColumnaN(matrizB, j)
			valor := 0
			for k := range len(fila) {
				valor += fila[k] * columna[k]
			}
			resultado[i][j] = valor
		}
	}
	return resultado, true
}

func obtenerColumnaN(matriz [][]int, n int) []int {
	resultado := make([]int, 0)
	for i := range len(matriz) {
		for j := range len(matriz[0]) {
			if j == n {
				resultado = append(resultado, matriz[i][j])
			}
		}
	}
	return resultado
}

func transpuestaMatriz(matrizA [][]int) ([][]int, bool) {

	resultado := make([][]int, len(matrizA[0]))

	for i := range len(resultado) {
		resultado[i] = make([]int, len(matrizA))
	}

	for i := range len(resultado[0]) {
		fila := matrizA[i]
		rellenarColumna(resultado, fila, i)
	}
	return resultado, true
}

func rellenarColumna(matriz [][]int, fila []int, n int) {
	count := 0
	for i := range len(matriz) {
		matriz[i][n] = fila[count]
		count++
	}
}
