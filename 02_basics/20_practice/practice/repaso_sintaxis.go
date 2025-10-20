package main

import (
	"fmt"
)

/*
EJERCICIOS DE REPASO DE SINTAXIS GO
====================================
Estos ejercicios están diseñados para recordar la practice básica de Go.
Completa cada función según las instrucciones en los comentarios.
*/

func main() {
	fmt.Println("=== REPASO DE SINTAXIS GO ===")

	ejercicio1Variables()
	ejercicio2Constantes()
	ejercicio3TiposDatos()
	ejercicio4Operadores()
	ejercicio5Condicionales()
	ejercicio6Loops()
	ejercicio7Arrays()
	ejercicio8Slices()
	ejercicio9Maps()
	ejercicio10Range()
	ejercicio11Funciones()
	ejercicio12Defer()
}

// ============================================================
// EJERCICIO 1: VARIABLES
// ============================================================
// Crea las siguientes variables:
// - Una variable "nombre" de tipo string con tu nombre
// - Una variable "edad" de tipo int con tu edad (usa :=)
// - Una variable "altura" de tipo float64 sin inicializar
// - Asigna un valor a altura
// - Imprime todas las variables
func ejercicio1Variables() {
	fmt.Println("\n--- Ejercicio 1: Variables ---")

	var nombre string = "Felipe"
	edad := 20
	var altura float64
	altura = 1.75

	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	fmt.Println("Altura:", altura)

}

// ============================================================
// EJERCICIO 2: CONSTANTES
// ============================================================
// Crea las siguientes constantes:
// - Una constante PI con valor 3.14159
// - Un grupo de constantes con los días de la semana (LUNES = 1, MARTES = 2, etc.)
// - Imprime PI y dos días de la semana
func ejercicio2Constantes() {
	fmt.Println("\n--- Ejercicio 2: Constantes ---")

	const PI = 3.141592
	const (
		LUNES     = 1
		MARTES    = 2
		MIERCOLES = 3
		JUEVES    = 4
		VIERNES   = 5
		SABADO    = 6
		DOMINGO   = 7
	)

	fmt.Println("PI:", PI)
	fmt.Println("Lunes:", LUNES)
	fmt.Println("Martes:", MARTES)
	fmt.Println("Miércoles:", MIERCOLES)
	fmt.Println("Jueves:", JUEVES)
	fmt.Println("Viernes:", VIERNES)
	fmt.Println("Sábado:", SABADO)
	fmt.Println("Domingo:", DOMINGO)
}

// ============================================================
// EJERCICIO 3: TIPOS DE DATOS
// ============================================================
// Declara variables de los siguientes tipos y asígnales valores:
// - int, int32, int64
// - float32, float64
// - bool
// - string
// - Imprime cada una mostrando su tipo y valor
func ejercicio3TiposDatos() {
	fmt.Println("\n--- Ejercicio 3: Tipos de Datos ---")

	entero32 := int32(100)
	entero64 := int64(1000)
	entero := 50

	flotante32 := float32(10.5)
	flotante64 := float64(20.99)

	booleano := true

	cadena := "Hola, Go!"

	fmt.Printf("int: %d (Tipo: %T)\n", entero, entero)
	fmt.Printf("int32: %d (Tipo: %T)\n", entero32, entero32)
	fmt.Printf("int64: %d (Tipo: %T)\n", entero64, entero64)
	fmt.Printf("float32: %.2f (Tipo: %T)\n", flotante32, flotante32)
	fmt.Printf("float64: %.2f (Tipo: %T)\n", flotante64, flotante64)
	fmt.Printf("bool: %t (Tipo: %T)\n", booleano, booleano)
	fmt.Printf("string: %s (Tipo: %T)\n", cadena, cadena)

}

// ============================================================
// EJERCICIO 4: OPERADORES
// ============================================================
// Declara dos variables numéricas (a = 10, b = 3)
// Realiza e imprime:
// - Suma, resta, multiplicación, división
// - Módulo (resto)
// - Comparaciones: ==, !=, <, >, <=, >=
// - Operadores lógicos: && (AND), || (OR), ! (NOT)
func ejercicio4Operadores() {
	fmt.Println("\n--- Ejercicio 4: Operadores ---")

	a, b := 10, 3
	fmt.Println("Variables: a =", a, ", b =", b)

	// Operaciones aritméticas
	fmt.Println("Suma:", a+b)
	fmt.Println("Resta:", a-b)
	fmt.Println("Multiplicación:", a*b)
	fmt.Println("División:", a/b)
	fmt.Println("Módulo:", a%b)

	// Comparaciones
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a <= b:", a <= b)
	fmt.Println("a >= b:", a >= b)

	// Operadores lógicos
	fmt.Println("(a > 5) && (b < 5):", (a > 5) && (b < 5))
	fmt.Println("(a < 5) || (b < 5):", (a < 5) || (b < 5))
	fmt.Println("!(a > 5):", !(a > 5))

}

// ============================================================
// EJERCICIO 5: CONDICIONALES
// ============================================================
// Crea una variable "temperatura" con valor 25
// Usa if-else para imprimir:
// - "Hace frío" si temperatura < 15
// - "Clima agradable" si temperatura está entre 15 y 25
// - "Hace calor" si temperatura > 25
// Luego crea una variable "dia" (1-7) y usa switch para imprimir el nombre del día
func ejercicio5Condicionales() {
	fmt.Println("\n--- Ejercicio 5: Condicionales ---")

	temperatura := 25
	if temperatura < 15 {
		fmt.Println("Hace frío")
	} else if temperatura >= 15 && temperatura <= 25 {
		fmt.Println("Clima agradable")
	} else {
		fmt.Println("Hace calor")
	}

	dia := 3
	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sábado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Día inválido")
	}

}

// ============================================================
// EJERCICIO 6: LOOPS
// ============================================================
// Escribe los siguientes loops:
// - Un for que imprima números del 1 al 5
// - Un for que imprima números pares del 0 al 10
// - Un while (for con condición) que cuente de 10 a 1
// - Un for infinito que se rompa con break cuando i == 3
func ejercicio6Loops() {
	fmt.Println("\n--- Ejercicio 6: Loops ---")

	fmt.Println("For del 1 al 5")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("For de números pares del 0 al 10")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("While de 10 a 1")
	i := 10
	for i >= 1 {
		fmt.Println(i)
		i--
	}

	fmt.Println("For infinito")
	i = 0
	for {
		fmt.Println("Infinito: ", i)
		i++
		if i == 3 {
			break
		}
	}
}

// ============================================================
// EJERCICIO 7: ARRAYS
// ============================================================
// Crea un array de 5 enteros con valores [10, 20, 30, 40, 50]
// - Imprime el array completo
// - Imprime el tercer elemento (índice 2)
// - Modifica el primer elemento a 100
// - Imprime la longitud del array
func ejercicio7Arrays() {
	fmt.Println("\n--- Ejercicio 7: Arrays ---")

	var numeros = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array completo:", numeros)
	fmt.Println("Tercer elemento:", numeros[2])
	numeros[0] = 100
	fmt.Println("Array modificado:", numeros)
	fmt.Println("Longitud del array:", len(numeros))

	// Otras maneras de declarar arrays
	array2 := [...]int{1, 2, 3, 4, 5} // El compilador infiere la longitud
	fmt.Println("Array 2:", array2)

	var array3 [3]string // Array de strings sin inicializar
	array3[0] = "Go"
	array3[1] = "Python"
	array3[2] = "Java"
	fmt.Println("Array 3:", array3)

}

// ============================================================
// EJERCICIO 8: SLICES
// ============================================================
// Crea un slice de strings con 3 nombres
// - Imprime el slice
// - Agrega 2 nombres más usando append
// - Imprime la longitud y capacidad del slice
// - Crea un sub-slice con los elementos del índice 1 al 3
// - Imprime el sub-slice
func ejercicio8Slices() {
	fmt.Println("\n--- Ejercicio 8: Slices ---")

	nombres := []string{"Felipe", "Sami", "Jonny"}
	fmt.Println("Nombres:", nombres)

	nombres = append(nombres, "Esteban", "Juanito")
	fmt.Println("Nombres:", nombres)

	sub_slice := nombres[1:3]
	fmt.Println("Nombres subslice:", sub_slice)

}

// ============================================================
// EJERCICIO 9: MAPS
// ============================================================
// Crea un map de string a int llamado "edades"
// - Agrega 3 personas con sus edades
// - Imprime el map completo
// - Imprime la edad de una persona específica
// - Verifica si existe una persona en el map
// - Elimina una persona del map
// - Imprime el map después de la eliminación
func ejercicio9Maps() {
	fmt.Println("\n--- Ejercicio 9: Maps ---")

	var edades = make(map[string]int)
	edades["Felipe"] = 20
	edades["Sami"] = 30
	edades["Jonny"] = 40
	fmt.Println("Edades:", edades)

	fmt.Println("Edades:", edades["Felipe"])

	valor, existe := edades["Felipe"]
	if existe {
		fmt.Println("Existe! Valor:", valor)
	}

	delete(edades, "Felipe")

	fmt.Println("Edades:", edades)

}

// ============================================================
// EJERCICIO 10: RANGE
// ============================================================
// Crea un slice de números [1, 2, 3, 4, 5]
// - Usa range para imprimir cada número y su índice
// - Crea un map de frutas {"manzana": 5, "banana": 3, "naranja": 8}
// - Usa range para imprimir cada fruta y su cantidad
func ejercicio10Range() {
	fmt.Println("\n--- Ejercicio 10: Range ---")

	numeros := []int{1, 2, 3, 4, 5}
	for i, v := range numeros {
		fmt.Println(i, v)
	}

	frutas := map[string]int{
		"manzana": 5,
		"banana":  3,
		"naranja": 8,
	}

	for k, v := range frutas {
		fmt.Println(k, ":", v)
	}

}

// ============================================================
// EJERCICIO 11: FUNCIONES
// ============================================================
// Crea las siguientes funciones y llámalas desde aquí:
// - sumar(a, b int) int - retorna la suma de dos números
// - dividir(a, b float64) (float64, error) - retorna división y error si b es 0
// - saludar(nombres ...string) - función variádica que saluda a todos los nombres

func ejercicio11Funciones() {
	fmt.Println("\n--- Ejercicio 11: Funciones ---")

	fmt.Println("3 + 6 = ", sumar(3, 6))

	result, err := dividir(3, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("3 / 6 = ", result)
	}

	nombres := []string{"Felipe", "Sami", "Jonny", "Juanito", "Emi"}
	saludar(nombres...)

}

// Funciones auxiliares para el ejercicio 11:
// Escribe aquí tus funciones sumar, dividir y saludar
func sumar(a, b int) int {
	return a + b
}

func dividir(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func saludar(nombres ...string) {
	for _, nombre := range nombres {
		fmt.Println("Hola " + nombre)
	}
}

// ============================================================
// EJERCICIO 12: DEFER
// ============================================================
// Crea una función que:
// - Imprima "Inicio"
// - Use defer para imprimir "Fin"
// - Imprima "Procesando..."
// - Usa múltiples defer para ver el orden de ejecución (LIFO)
func auxiliar() {
	defer fmt.Println("Fin Final")
	defer fmt.Println("Fin")

	fmt.Println("Inicio")
	fmt.Println("Procesando")
}

func ejercicio12Defer() {
	fmt.Println("\n--- Ejercicio 12: Defer ---")
	auxiliar()
}
