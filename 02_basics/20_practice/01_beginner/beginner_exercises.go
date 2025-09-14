package main

import "fmt"

/*
EJERCICIOS NIVEL PRINCIPIANTE
=============================

Estos ejercicios cubren los conceptos básicos de Go que has aprendido:
- Variables y tipos de datos
- Operadores aritméticos y lógicos
- Constantes
- Condicionales básicos
- Bucles simples

Instrucciones: Implementa cada función según la descripción en los comentarios.
Puedes probar tu código ejecutando: go run beginner_exercises.go
*/

func main() {
	fmt.Println("=== EJERCICIOS NIVEL PRINCIPIANTE ===")

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
EJERCICIO 1: Variables y Tipos de Datos
=======================================
Declara variables de diferentes tipos y asígnalas valores apropiados.

Instrucciones:
1. Declara una variable 'nombre' de tipo string con tu nombre
2. Declara una variable 'edad' de tipo int con tu edad
3. Declara una variable 'altura' de tipo float64 con tu altura en metros
4. Declara una variable 'esEstudiante' de tipo bool indicando si eres estudiante
5. Imprime todas las variables con un formato legible

Ejemplo de salida:
Nombre: Juan
Edad: 25 años
Altura: 1.75 metros
Es estudiante: true
*/
func ejercicio1() {
	// TODO: Implementa aquí tu solución

}

/*
EJERCICIO 2: Constantes y Operadores Aritméticos
===============================================
Utiliza constantes para valores matemáticos y realiza cálculos.

Instrucciones:
1. Declara una constante PI con el valor 3.14159
2. Declara una constante GRAVEDAD con el valor 9.81
3. Recibe un radio como parámetro y calcula el área de un círculo
4. Recibe una altura y tiempo, calcula la distancia recorrida en caída libre
5. Imprime los resultados formateados

Fórmulas:
- Área del círculo: PI * radio * radio
- Distancia en caída libre: 0.5 * GRAVEDAD * tiempo * tiempo

Ejemplo de salida para radio=5, altura=10, tiempo=2:
Área del círculo con radio 5.00: 78.54
Distancia en caída libre tras 2.00 segundos: 19.62 metros
*/
func ejercicio2(radio float64, tiempo float64) {
	// TODO: Implementa aquí tu solución

}

/*
EJERCICIO 3: Condicionales y Operadores de Comparación
=====================================================
Implementa un sistema de calificaciones.

Instrucciones:
1. Recibe una calificación numérica (0-100)
2. Determina la letra correspondiente según estos criterios:
  - 90-100: A (Excelente)
  - 80-89: B (Bueno)
  - 70-79: C (Regular)
  - 60-69: D (Suficiente)
  - 0-59: F (Insuficiente)

3. Valida que la calificación esté en el rango correcto
4. Imprime el resultado con un mensaje descriptivo

Ejemplo de salida para calificacion=85:
Calificación: 85
Letra: B
Nivel: Bueno
*/
func ejercicio3(calificacion int) {
	// TODO: Implementa aquí tu solución

}

/*
EJERCICIO 4: Bucles y Operadores Lógicos
=======================================
Analiza una serie de números enteros.

Instrucciones:
1. Recibe un slice de números enteros
2. Usando un bucle for, calcula:
  - La suma total de todos los números
  - El promedio (suma / cantidad)
  - Cuántos números son pares
  - Cuántos números son impares
  - El número mayor
  - El número menor

3. Imprime un reporte completo

Ejemplo de salida para números = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]:
Números analizados: [1 2 3 4 5 6 7 8 9 10]
Suma total: 55
Promedio: 5.50
Números pares: 5
Números impares: 5
Número mayor: 10
Número menor: 1
*/
func ejercicio4(numeros []int) {
	// TODO: Implementa aquí tu solución

}

/*
EJERCICIO 5: Manipulación de Strings y Bucles
============================================
Procesamiento de texto básico.

Instrucciones:
1. Recibe una palabra o frase
2. Cuenta cuántas vocales (a, e, i, o, u) contiene (insensible a mayúsculas)
3. Cuenta cuántas consonantes contiene
4. Determina si la palabra es un palíndromo (se lee igual al derecho y al revés)
5. Convierte la primera letra de cada palabra a mayúscula
6. Imprime todos los resultados

Ejemplo de salida para texto="hola mundo":
Texto original: "hola mundo"
Vocales: 4
Consonantes: 6
Es palíndromo: false
Texto capitalizado: "Hola Mundo"
*/
func ejercicio5(texto string) {
	// TODO: Implementa aquí tu solución
	// Pista: Puedes usar strings.ToLower(), strings.ToUpper(), strings.Fields()

}

/*
EJERCICIO 6: Bucles Anidados y Patrones
=====================================
Genera patrones de asteriscos.

Instrucciones:
1. Recibe un número entero n
2. Genera un triángulo de asteriscos de altura n
3. Genera un triángulo invertido de asteriscos de altura n
4. Genera un rombo de asteriscos (triángulo + triángulo invertido)

Ejemplo de salida para n=4:
Triángulo:
*
**
***
****

Triángulo invertido:
****
***
**
*

Rombo:

	  *
	 **
	***

****

	***
	 **
	  *
*/
func ejercicio6(n int) {
	// TODO: Implementa aquí tu solución

}

/*
EJERCICIO 7: Switch y Operadores
===============================
Sistema de conversión de unidades.

Instrucciones:
1. Recibe un valor numérico y una unidad de origen
2. Recibe una unidad de destino
3. Realiza la conversión usando switch
4. Soporta las siguientes conversiones de temperatura:
  - Celsius a Fahrenheit: (C * 9/5) + 32
  - Fahrenheit a Celsius: (F - 32) * 5/9
  - Celsius a Kelvin: C + 273.15
  - Kelvin a Celsius: K - 273.15
  - Fahrenheit a Kelvin: (F - 32) * 5/9 + 273.15
  - Kelvin a Fahrenheit: (K - 273.15) * 9/5 + 32

Ejemplo de salida para valor=25, origen="C", destino="F":
25.00°C = 77.00°F
*/
func ejercicio7(valor float64, origen, destino string) {
	// TODO: Implementa aquí tu solución

}

/*
EJERCICIO 8: Validación de Datos
===============================
Sistema de validación de entrada de usuario.

Instrucciones:
1. Recibe un nombre de usuario (string)
2. Recibe una contraseña (string)
3. Recibe una edad (int)
4. Valida que:
  - El nombre de usuario tenga entre 3 y 20 caracteres
  - La contraseña tenga al menos 8 caracteres
  - La contraseña contenga al menos una letra mayúscula
  - La contraseña contenga al menos una letra minúscula
  - La contraseña contenga al menos un número
  - La edad esté entre 13 y 120 años

5. Imprime si cada validación es exitosa o falla

Ejemplo de salida para usuario="Juan123", password="MiPass123", edad=25:
Validación de usuario: Juan123
✓ Longitud de usuario válida (7 caracteres)
✓ Longitud de contraseña válida (9 caracteres)
✓ Contraseña contiene mayúscula
✓ Contraseña contiene minúscula
✓ Contraseña contiene número
✓ Edad válida (25 años)
Resultado: TODAS LAS VALIDACIONES EXITOSAS
*/
func ejercicio8(usuario, password string, edad int) {
	// TODO: Implementa aquí tu solución

}

/*
EJERCICIO 9: Calculadora Básica
=============================
Implementa una calculadora con operaciones básicas.

Instrucciones:
1. Recibe dos números (float64) y un operador (string)
2. Realiza la operación correspondiente:
  - "+" para suma
  - "-" para resta
  - "*" para multiplicación
  - "/" para división
  - "%" para módulo (solo para enteros)
  - "^" para potencia (implementa tu propia función de potencia usando bucles)

3. Valida que no se divida por cero
4. Maneja operadores inválidos
5. Imprime el resultado de la operación

Ejemplo de salida para a=10, b=3, operador="/":
10.00 / 3.00 = 3.33
*/
func ejercicio9(a, b float64, operador string) {
	// TODO: Implementa aquí tu solución

}

/*
EJERCICIO 10: Números Primos y Perfectos
=======================================
Análisis matemático de números.

Instrucciones:
1. Recibe un número entero n
2. Determina si n es un número primo
3. Determina si n es un número perfecto (suma de sus divisores = número)
4. Lista todos los números primos menores que n
5. Lista todos los números perfectos menores que n
6. Calcula el factorial de n (si n <= 20 para evitar overflow)

Ejemplo de salida para n=12:
Análisis del número 12:
¿Es primo? false
¿Es perfecto? false
Números primos menores que 12: [2 3 5 7 11]
Números perfectos menores que 12: [6]
Factorial de 12: 479001600
*/
func ejercicio10(n int) {
	// TODO: Implementa aquí tu solución

}

// Funciones de prueba - NO MODIFICAR
func testExercise1() {
	fmt.Println("\n--- EJERCICIO 1 ---")
	ejercicio1()
}

func testExercise2() {
	fmt.Println("\n--- EJERCICIO 2 ---")
	ejercicio2(5.0, 2.0)
}

func testExercise3() {
	fmt.Println("\n--- EJERCICIO 3 ---")
	ejercicio3(85)
	ejercicio3(92)
	ejercicio3(67)
	ejercicio3(45)
}

func testExercise4() {
	fmt.Println("\n--- EJERCICIO 4 ---")
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ejercicio4(numeros)
}

func testExercise5() {
	fmt.Println("\n--- EJERCICIO 5 ---")
	ejercicio5("hola mundo")
	ejercicio5("reconocer")
}

func testExercise6() {
	fmt.Println("\n--- EJERCICIO 6 ---")
	ejercicio6(4)
}

func testExercise7() {
	fmt.Println("\n--- EJERCICIO 7 ---")
	ejercicio7(25, "C", "F")
	ejercicio7(77, "F", "C")
	ejercicio7(298.15, "K", "C")
}

func testExercise8() {
	fmt.Println("\n--- EJERCICIO 8 ---")
	ejercicio8("Juan123", "MiPass123", 25)
	ejercicio8("ab", "pass", 150)
}

func testExercise9() {
	fmt.Println("\n--- EJERCICIO 9 ---")
	ejercicio9(10, 3, "/")
	ejercicio9(5, 4, "+")
	ejercicio9(2, 8, "^")
}

func testExercise10() {
	fmt.Println("\n--- EJERCICIO 10 ---")
	ejercicio10(12)
	ejercicio10(7)
}
