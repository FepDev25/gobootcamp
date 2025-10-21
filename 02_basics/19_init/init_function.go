package main

import "fmt"

var name string

func init() {
	name = "Felipe" // Inicializa la variable name
	fmt.Println("Función init ejecutada")
}

func main() {
	// La función init se ejecuta automáticamente antes de la función main.
	// Se utiliza para inicializar variables o configurar el entorno.
	fmt.Println("Función main ejecutada")
	fmt.Println(name) // Imprime el valor inicializado en init
}
