package main

import (
	"fmt"
	"os"
)

func main() {
	// Terminar el programa inmediatamente con un código de salida 0 (éxito).
	// Nota: os.Exit no ejecuta funciones defer.
	// Por lo tanto, cualquier función defer declarada antes de llamar a os.Exit no se ejecutará.
	process()
}

func process() {
	defer fmt.Println("Defer no se ejecuta si usamos os.Exit")

	fmt.Println("Inicio del programa")
	os.Exit(0)
	fmt.Println("Este mensaje no se imprimirá")
}
