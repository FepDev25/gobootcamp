package main

import (
	"fmt"
	"strings"
)

var petName = "Pedrito" // Alcance global dentro del paquete

func main() {
	// Variables
	var age int                // Declarar una variable
	var name string = "Felipe" // Inicializar una variable
	var name2 = "Juan"         // Inicializar otra variable sin especificar tipo
	count := 1                 // Inicializar una variable con el operador :=
	lastname := "Peralta"      // Inicializar otra variable con el operador :=

	fmt.Println(name, name2, lastname, age, count)

	// Scope
	printName(petName)

}

func printName(name string) {
	uppName := strings.ToUpper(name) // Alcance local
	fmt.Println(uppName)
}
