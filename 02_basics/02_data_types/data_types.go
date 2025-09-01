package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Tipos de datos
func main() {
	primitivos()
	compuestos()
	especiales()
}

func primitivos() {
	// -------------------- Tipos Primitivos --------------------

	// ********** Enteros **********
	// Enteros con signo
	var int8 int8 = 12                    // 8 bits
	var int16 int16 = 1234                // 16 bits
	var int32 int32 = 12345678            // 32 bits
	var int64 int64 = 1234567890123456789 // 64 bits
	var int = 123456                      // Depende de la arquitectura (32 o 64 bits)

	fmt.Println("Int8:", int8)
	fmt.Println("Int16:", int16)
	fmt.Println("Int32:", int32)
	fmt.Println("Int64:", int64)
	fmt.Println("Int:", int)

	// Enteros sin signo
	var uint8 uint8 = 12                    // 8 bits
	var uint16 uint16 = 1234                // 16 bits
	var uint32 uint32 = 12345678            // 32 bits
	var uint64 uint64 = 1234567890123456789 // 64 bits
	var uint uint = 123456                  // Depende de la arquitectura (32 o 64 bits)
	var uintptr uintptr = 0                 // Tamaño de un puntero

	fmt.Println("Uint8:", uint8)
	fmt.Println("Uint16:", uint16)
	fmt.Println("Uint32:", uint32)
	fmt.Println("Uint64:", uint64)
	fmt.Println("Uint:", uint)
	fmt.Println("Uintptr:", uintptr)

	// ********** Flotantes **********
	var float32 float32 = 3.14                   // 32 bits
	var float64 float64 = 3.14159265358979323846 // 64 bits

	fmt.Println("Float32:", float32)
	fmt.Println("Float64:", float64)

	// ********** Complejos **********
	var complex64 complex64 = 1 + 2i   // 32 bits para cada parte
	var complex128 complex128 = 1 + 2i // 64 bits para cada parte

	fmt.Println("Complex64:", complex64)
	fmt.Println("Complex128:", complex128)

	// ********** Booleanos **********
	var bool1 bool = true
	var bool2 bool = false

	fmt.Println("Bool1:", bool1)
	fmt.Println("Bool2:", bool2)

	// ********** Strings **********
	var str1 string = "Hola mundo"
	var str2 string = "Felipe Peralta"

	fmt.Println("Str1:", str1)
	fmt.Println("Str2:", str2)
}

func compuestos() {
	// -------------------- Tipos Compuestos --------------------

	// ********** Constantes **********
	const E float64 = 2.718281828459045
	const EMPRESA string = "Mi Empresa"
	const ACTIVO bool = true

	fmt.Println("E:", E)
	fmt.Println("EMPRESA:", EMPRESA)
	fmt.Println("ACTIVO:", ACTIVO)

	// ********** Arrays **********
	var numeros []int = []int{21, 20, 0, 52, 52}
	var nombres [3]string = [3]string{"Felipe", "Emilia", "Karen"}
	var matriz [2][2]int = [2][2]int{{2, 2}, {1, 10}}

	fmt.Println("Numeros:", numeros)
	fmt.Println("Nombres:", nombres)
	fmt.Println("Matriz:", matriz)

	// ********** Structs **********
	type Persona struct {
		Nombre string
		Edad   int
		Email  string
	}
	var me Persona = Persona{Nombre: "Felipe", Edad: 20, Email: "felipe@example.com"}
	fmt.Println("Yo:", me)

	// ********** Punteros **********
	var x int = 45
	var ptr *int = &x    // & obtiene la dirección de memoria
	var valor int = *ptr // * obtiene el valor en la dirección de memoria

	fmt.Println("Valor:", valor)
	fmt.Println("Puntero:", ptr)

	// ********** Maps **********
	var edades map[string]int = make(map[string]int)
	edades["Felipe"] = 20
	edades["Emilia"] = 21

	fmt.Println("Edades:", edades)

	colores := map[string]string{
		"rojo":     "#FF0000",
		"verde":    "#00FF00",
		"azul":     "#0000FF",
		"amarillo": "#FFFF00",
	}
	fmt.Println("Colores:", colores)

	// ********** Slices **********
	var nums []int = []int{1, 2, 3, 4, 5}
	nums = append(nums, 6)

	fmt.Println("Numeros:", nums)

	// ********** Funciones **********
	imprimirNumeros(nums)

	var multiplicar = func(x int, y int) int {
		return x * y
	}

	fmt.Println("Multiplicar 3 * 4 =", multiplicar(3, 4))
}

func imprimirNumeros(numeros []int) {
	for _, n := range numeros {
		fmt.Println("Numero:", n)
	}
}

func especiales() {
	// -------------------- Tipos Especiales --------------------

	// ********** JSON **********
	type Usuario struct {
		Nombre string
		Edad   int
	}

	usuario := Usuario{Nombre: "Felipe", Edad: 20}
	jsonData, err := json.Marshal(usuario)

	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	fmt.Println("JSON:", string(jsonData))

	// ********** Texto **********

	// Manipulacion de texto
	texto := strings.ToUpper("rick and morty")
	partes := strings.Split(texto, " AND ")

	fmt.Println("Texto en mayúsculas:", texto)
	fmt.Println("Partes:", partes)

	// Conversion de tipos
	numero, err := strconv.Atoi("12345")              // string a int
	texto_num := strconv.Itoa(67890)                  // int a string
	decimal, err := strconv.ParseFloat("3.14159", 64) // string a float64

	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	fmt.Println("Numero:", numero)
	fmt.Println("Texto:", texto_num)
	fmt.Println("Decimal:", decimal)

	// ********** Zero Values **********
	var entero int
	var flotante float64
	var booleano bool
	var cadena string
	var arreglo [5]int
	var slice []int
	var mapa map[string]int
	var estructura Usuario
	var puntero *int

	fmt.Println("Entero:", entero)
	fmt.Println("Flotante:", flotante)
	fmt.Println("Booleano:", booleano)
	fmt.Println("Cadena:", cadena)
	fmt.Println("Arreglo:", arreglo)
	fmt.Println("Slice:", slice)
	fmt.Println("Mapa:", mapa)
	fmt.Println("Estructura:", estructura)
	fmt.Println("Puntero:", puntero)

	// ********** Verificación de Tipos en Runtime **********
	var interfaz any = "Hola"

	if texto, ok := interfaz.(string); ok {
		fmt.Println("Es un string:", texto)
	}

	switch v := interfaz.(type) {
	case string:
		fmt.Println("String:", v)
	case int:
		fmt.Println("Integer:", v)
	default:
		fmt.Println("Tipo desconocido")
	}
}
