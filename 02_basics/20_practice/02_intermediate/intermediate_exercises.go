package main

import (
	"fmt"
	"intermediate/ejercicios"
)

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
	testExercise1()
	testExercise2()
	testExercise3()
	testExercise4()
	testExercise5()
	testExercise6()
	testExercise7()
	testExercise8()
	testExercise9()
	testExercise10()

	fmt.Println("¡Todos los ejercicios completados!")
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
	duplicados, invertido, ordenado, suma, promedio, mediana, moda := ejercicios.Ejercicio1(numeros)
	fmt.Printf("Original: %v\n", numeros)
	fmt.Printf("Sin duplicados: %v\n", duplicados)
	fmt.Printf("Invertido: %v\n", invertido)
	fmt.Printf("Ordenado: %v\n", ordenado)
	fmt.Printf("Suma: %d, Promedio: %.2f, Mediana: %.2f, Moda: %v\n", suma, promedio, mediana, moda)
}

func testExercise2() {
	fmt.Println("\n--- EJERCICIO 2 ---")
	ejercicios.Ejercicio2()
}

func testExercise3() {
	fmt.Println("\n--- EJERCICIO 3 ---")
	texto := "Hola mundo. ¿Cómo estás? ¡Muy bien!\nHola mundo. Chelsea FC."
	palabras, chars, palabraMasComun, charMasComun, oraciones, parrafos := ejercicios.Ejercicio3(texto)
	fmt.Printf("Texto: %s\n", texto)
	fmt.Printf("Palabras: %v\n", palabras)
	fmt.Printf("Caracteres: %v\n", chars)
	fmt.Printf("Palabra más común: %s\n", palabraMasComun)
	fmt.Printf("Carácter más común: %c\n", charMasComun)
	fmt.Printf("Oraciones: %d, Párrafos: %d\n", oraciones, parrafos)
}

func testExercise4() {
	fmt.Println("\n--- EJERCICIO 4 ---")
	matrizA := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrizB := [][]int{{2, 5, 1}, {1, 10, 2}, {0, 11, 12}}

	fmt.Println("Matriz A:", matrizA)
	fmt.Println("Matriz B:", matrizB)

	suma, exitoso := ejercicios.Ejercicio4(matrizA, matrizB, "suma")
	fmt.Printf("Suma exitosa: %v, Resultado: %v\n", exitoso, suma)

	resta, exitoso := ejercicios.Ejercicio4(matrizA, matrizB, "resta")
	fmt.Printf("Resta exitosa: %v, Resultado: %v\n", exitoso, resta)

	producto, exitoso := ejercicios.Ejercicio4(matrizA, matrizB, "multi")
	fmt.Printf("Multiplicación exitosa: %v, Resultado: %v\n", exitoso, producto)

	matriz := [][]int{{0, 0, 4}, {1, 0, 4}, {0, 1, 0}, {0, 3, 2}, {0, 2, 3}, {0, 3, 4}, {3, 3, 1}}
	fmt.Println("Matriz:", matriz)

	transpuesta, exitoso := ejercicios.Ejercicio4(matriz, nil, "trans")
	fmt.Printf("Transpuesta exitosa: %v, Resultado: %v\n", exitoso, transpuesta)
}

func testExercise5() {
	fmt.Println("\n--- EJERCICIO 5 ---")
	ejercicios.Ejercicio5()
}

func testExercise6() {
	fmt.Println("\n--- EJERCICIO 6 ---")
	numeros := []int{64, 34, 25, 12, 22, 11, 90, 11, 10, 102, 14, 7, 6, 0, 12}
	ejercicios.Ejercicio6(numeros, 22)
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
