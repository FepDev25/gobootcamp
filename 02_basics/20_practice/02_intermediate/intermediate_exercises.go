package main

import "fmt"

/*
EJERCICIOS NIVEL INTERMEDIO
===========================

Estos ejercicios cubren conceptos intermedios de Go:
- Arrays, slices y maps avanzados
- Funciones con múltiples valores de retorno
- Range y manipulación de datos complejos
- Control de flujo más sofisticado
- Manejo básico de errores

Instrucciones: Implementa cada función según la descripción en los comentarios.
Puedes probar tu código ejecutando: go run intermediate_exercises.go
*/

func main() {
	fmt.Println("=== EJERCICIOS NIVEL INTERMEDIO ===")

	// Descomenta las líneas siguientes cuando hayas implementado las funciones
	// testExercise1()
	// testExercise2()
	// testExercise3()
	// testExercise4()
	// testExercise5()
	// testExercise6()
	// testExercise7()
	// testExercise8()
	// testExercise9()
	// testExercise10()

	fmt.Println("¡Todos los ejercicios completados!")
}

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
func ejercicio1(numeros []int) ([]int, []int, []int, int, float64, float64, []int) {
	// TODO: Implementa aquí tu solución
	// Retorna: duplicados, invertido, ordenado, suma, promedio, mediana, moda

	return nil, nil, nil, 0, 0, 0, nil
}

/*
EJERCICIO 2: Sistema de Inventario con Maps
==========================================
Gestiona un inventario de productos usando maps.

Instrucciones:
1. Define un map donde la clave es el nombre del producto (string) y el valor es la cantidad (int)
2. Implementa las siguientes operaciones que retornen bool indicando éxito:
  - Agregar producto: si ya existe, suma la cantidad
  - Remover producto: elimina del inventario
  - Actualizar cantidad: cambia la cantidad de un producto existente
  - Buscar producto: verifica si existe y retorna la cantidad

3. Implementa función para mostrar reporte del inventario

Ejemplo de uso:
Inventario inicial: map[manzanas:50 naranjas:30 plátanos:25]
Agregar manzanas (20): true → manzanas: 70
Remover naranjas: true
Buscar plátanos: true, cantidad: 25
Inventario final: map[manzanas:70 plátanos:25]
*/
func ejercicio2() {
	// TODO: Implementa aquí tu solución
	// Crea el inventario inicial y ejecuta las operaciones

}

/*
EJERCICIO 3: Análisis de Texto Avanzado
======================================
Procesa texto y genera estadísticas detalladas.

Instrucciones:
1. Recibe un texto (string)
2. Retorna múltiples valores con estadísticas:
  - Mapa de frecuencia de palabras
  - Mapa de frecuencia de caracteres (excluyendo espacios)
  - Palabra más común
  - Carácter más común
  - Número de oraciones (terminadas en ., !, ?)
  - Número de párrafos (separados por \n\n)

Ejemplo para texto="Hola mundo. ¿Cómo estás? ¡Muy bien!":
Frecuencia de palabras: map[Hola:1 mundo:1 Cómo:1 estás:1 Muy:1 bien:1]
Frecuencia de caracteres: map[H:1 o:4 l:2 a:2 ...]
Palabra más común: "mundo" (todas empatan)
Carácter más común: "o"
Oraciones: 3
Párrafos: 1
*/
func ejercicio3(texto string) (map[string]int, map[rune]int, string, rune, int, int) {
	// TODO: Implementa aquí tu solución
	// Pista: usa strings.Fields(), strings.Count(), range con string

	return nil, nil, "", 0, 0, 0
}

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
  - Determinante (solo para matrices 2x2 y 3x3)

3. Retorna el resultado y un boolean indicando si la operación fue exitosa

Ejemplo para matrices A=[[1,2],[3,4]] y B=[[5,6],[7,8]]:
Suma: [[6,8],[10,12]], exitoso: true
Resta: [[-4,-4],[-4,-4]], exitoso: true
Multiplicación: [[19,22],[43,50]], exitoso: true
Transpuesta de A: [[1,3],[2,4]]
Determinante de A: -2, exitoso: true
*/
func ejercicio4(matrizA, matrizB [][]int, operacion string) ([][]int, bool) {
	// TODO: Implementa aquí tu solución

	return nil, false
}

/*
EJERCICIO 5: Sistema de Calificaciones de Estudiantes
====================================================
Gestiona calificaciones de múltiples estudiantes y materias.

Instrucciones:
1. Define estructuras de datos usando maps anidados
2. Almacena: map[estudiante]map[materia][]calificaciones
3. Implementa funciones que retornen múltiples valores:
  - Agregar calificación para estudiante y materia
  - Calcular promedio de un estudiante en una materia
  - Calcular promedio general de un estudiante
  - Obtener el mejor estudiante por materia
  - Generar ranking general de estudiantes

Ejemplo:
Estudiantes: ["Ana", "Luis", "Maria"]
Materias: ["Matemáticas", "Historia", "Ciencias"]
Ana - Matemáticas: [85, 90, 92] → Promedio: 89.0
Luis - Matemáticas: [78, 82, 85] → Promedio: 81.7
Mejor en Matemáticas: Ana (89.0)
Ranking general: [Ana: 87.5, Luis: 83.2, Maria: 91.1]
*/
func ejercicio5() {
	// TODO: Implementa aquí tu solución
	// Crea la estructura de datos y demuestra todas las operaciones

}

/*
EJERCICIO 6: Algoritmo de Búsqueda y Ordenamiento
================================================
Implementa algoritmos de búsqueda y ordenamiento.

Instrucciones:
1. Recibe un slice de enteros
2. Implementa los siguientes algoritmos:
  - Búsqueda lineal: retorna índice del elemento o -1
  - Búsqueda binaria: retorna índice del elemento o -1 (requiere slice ordenado)
  - Ordenamiento burbuja: retorna slice ordenado y número de intercambios
  - Ordenamiento por selección: retorna slice ordenado y número de comparaciones

3. Cada función debe retornar múltiples valores según se especifica

Ejemplo para slice=[64, 34, 25, 12, 22, 11, 90]:
Búsqueda lineal de 22: índice 4, encontrado: true
Ordenamiento burbuja: [11 12 22 25 34 64 90], intercambios: 16
Búsqueda binaria de 25 en slice ordenado: índice 3, encontrado: true
*/
func ejercicio6(numeros []int, elemento int) {
	// TODO: Implementa aquí tu solución
	// Demuestra todos los algoritmos

}

/*
EJERCICIO 7: Generador de Secuencias Matemáticas
===============================================
Genera diferentes secuencias matemáticas usando range y slices.

Instrucciones:
1. Implementa funciones para generar secuencias:
  - Fibonacci hasta n términos
  - Números primos hasta un límite
  - Factoriales de 0 a n
  - Triángulo de Pascal de altura n
  - Serie armónica hasta n términos

2. Retorna slices con las secuencias generadas

Ejemplo para n=6:
Fibonacci: [0, 1, 1, 2, 3, 5]
Primos hasta 20: [2, 3, 5, 7, 11, 13, 17, 19]
Factoriales: [1, 1, 2, 6, 24, 120]
Triángulo de Pascal:
[1]
[1, 1]
[1, 2, 1]
[1, 3, 3, 1]
[1, 4, 6, 4, 1]
Serie armónica: [1.00, 0.50, 0.33, 0.25, 0.20, 0.17]
*/
func ejercicio7(n int) {
	// TODO: Implementa aquí tu solución

}

/*
EJERCICIO 8: Simulador de Juego de Cartas
========================================
Simula un mazo de cartas y operaciones básicas.

Instrucciones:
1. Crea un slice para representar un mazo completo (52 cartas)
2. Implementa las siguientes operaciones:
  - Barajar: aleatoriza el orden de las cartas
  - Repartir: retorna n cartas del tope del mazo
  - Clasificar mano: analiza una mano de 5 cartas
  - Calcular puntuación: asigna valores a las cartas

3. Usa maps para representar cartas y suits

Ejemplo:
Mazo creado: 52 cartas
Barajado: orden aleatorio
Mano repartida: [As♠, Rey♥, 10♦, 7♣, 2♠]
Clasificación: Par de Ases
Puntuación: 87 puntos
*/
func ejercicio8() {
	// TODO: Implementa aquí tu solución
	// Define estructuras para cartas y implements operaciones

}

/*
EJERCICIO 9: Analizador de Logs
==============================
Procesa y analiza archivos de log simulados.

Instrucciones:
1. Recibe un slice de strings representando líneas de log
2. Cada línea tiene formato: "TIMESTAMP LEVEL MESSAGE"
3. Analiza y retorna estadísticas:
  - Conteo por nivel de log (INFO, WARN, ERROR, DEBUG)
  - Horas con más actividad
  - Mensajes de error únicos
  - Tendencia de errores por hora

4. Retorna múltiples maps con las estadísticas

Ejemplo logs:
"2024-01-01 10:30:15 INFO Usuario logueado"
"2024-01-01 10:31:22 ERROR Base de datos no disponible"
"2024-01-01 11:15:33 WARN Memoria baja"

Resultado:
Niveles: map[INFO:45 WARN:12 ERROR:8 DEBUG:23]
Actividad por hora: map[10:15 11:20 12:18]
Errores únicos: map[Base de datos no disponible:3 Timeout de conexión:2]
*/
func ejercicio9(logs []string) (map[string]int, map[string]int, map[string]int) {
	// TODO: Implementa aquí tu solución
	// Pista: usa strings.Split(), strings.Fields()

	return nil, nil, nil
}

/*
EJERCICIO 10: Sistema de Reservas
================================
Gestiona un sistema de reservas para un hotel.

Instrucciones:
1. Usa maps anidados para gestionar: map[fecha]map[habitacion]cliente
2. Implementa las siguientes operaciones:
  - Hacer reserva: verifica disponibilidad y reserva
  - Cancelar reserva: libera la habitación
  - Consultar disponibilidad: lista habitaciones libres por fecha
  - Obtener reservas de cliente: lista todas sus reservas
  - Generar reporte de ocupación: estadísticas por fecha

3. Retorna múltiples valores según la operación

Ejemplo:
Reservas: map[2024-01-15:map[101:Juan 102:Maria] 2024-01-16:map[101:Pedro]]
Hacer reserva Juan, habitación 103, 2024-01-15: exitoso
Disponibilidad 2024-01-15: [104, 105, 106, 107, 108, 109, 110]
Reservas de Juan: [(103, 2024-01-15)]
Ocupación 2024-01-15: 3/10 habitaciones (30%)
*/
func ejercicio10() {
	// TODO: Implementa aquí tu solución
	// Define la estructura de datos y demuestra todas las operaciones

}

// Funciones de prueba - NO MODIFICAR
func testExercise1() {
	fmt.Println("\n--- EJERCICIO 1 ---")
	numeros := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	duplicados, invertido, ordenado, suma, promedio, mediana, moda := ejercicio1(numeros)
	fmt.Printf("Original: %v\n", numeros)
	fmt.Printf("Sin duplicados: %v\n", duplicados)
	fmt.Printf("Invertido: %v\n", invertido)
	fmt.Printf("Ordenado: %v\n", ordenado)
	fmt.Printf("Suma: %d, Promedio: %.2f, Mediana: %.2f, Moda: %v\n", suma, promedio, mediana, moda)
}

func testExercise2() {
	fmt.Println("\n--- EJERCICIO 2 ---")
	ejercicio2()
}

func testExercise3() {
	fmt.Println("\n--- EJERCICIO 3 ---")
	texto := "Hola mundo. ¿Cómo estás? ¡Muy bien!"
	palabras, chars, palabraMasComun, charMasComun, oraciones, parrafos := ejercicio3(texto)
	fmt.Printf("Texto: %s\n", texto)
	fmt.Printf("Palabras: %v\n", palabras)
	fmt.Printf("Caracteres: %v\n", chars)
	fmt.Printf("Palabra más común: %s\n", palabraMasComun)
	fmt.Printf("Carácter más común: %c\n", charMasComun)
	fmt.Printf("Oraciones: %d, Párrafos: %d\n", oraciones, parrafos)
}

func testExercise4() {
	fmt.Println("\n--- EJERCICIO 4 ---")
	matrizA := [][]int{{1, 2}, {3, 4}}
	matrizB := [][]int{{5, 6}, {7, 8}}

	suma, exitoso := ejercicio4(matrizA, matrizB, "suma")
	fmt.Printf("Suma exitosa: %v, Resultado: %v\n", exitoso, suma)

	producto, exitoso := ejercicio4(matrizA, matrizB, "multiplicacion")
	fmt.Printf("Multiplicación exitosa: %v, Resultado: %v\n", exitoso, producto)
}

func testExercise5() {
	fmt.Println("\n--- EJERCICIO 5 ---")
	ejercicio5()
}

func testExercise6() {
	fmt.Println("\n--- EJERCICIO 6 ---")
	numeros := []int{64, 34, 25, 12, 22, 11, 90}
	ejercicio6(numeros, 22)
}

func testExercise7() {
	fmt.Println("\n--- EJERCICIO 7 ---")
	ejercicio7(6)
}

func testExercise8() {
	fmt.Println("\n--- EJERCICIO 8 ---")
	ejercicio8()
}

func testExercise9() {
	fmt.Println("\n--- EJERCICIO 9 ---")
	logs := []string{
		"2024-01-01 10:30:15 INFO Usuario logueado",
		"2024-01-01 10:31:22 ERROR Base de datos no disponible",
		"2024-01-01 11:15:33 WARN Memoria baja",
		"2024-01-01 11:20:44 ERROR Base de datos no disponible",
		"2024-01-01 12:05:11 INFO Operación completada",
	}
	niveles, actividad, errores := ejercicio9(logs)
	fmt.Printf("Niveles: %v\n", niveles)
	fmt.Printf("Actividad: %v\n", actividad)
	fmt.Printf("Errores: %v\n", errores)
}

func testExercise10() {
	fmt.Println("\n--- EJERCICIO 10 ---")
	ejercicio10()
}
