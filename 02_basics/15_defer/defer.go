package main

import "fmt"

func main() {
	// Defer es una palabra clave en Go que se utiliza para posponer la ejecución de
	// una función hasta que la función que la contiene haya terminado.
	process()
	processWithMultipleDefers()
	processI(1)

	// Practical uses cases:
	// 1. Cerrar archivos
	// 2. Liberar recursos
	// 3. Desbloquear mutexes
	// 4. Registrar salidas o errores
	// Mejores prácticas;
	// 1. Usar defer para limpiar recursos
	// 2. Evitar defer en bucles intensivos
	// 3. Ser consciente del orden de ejecución
	// 4. No depender de defer para lógica crítica

	// Es como el bloque finally en otros lenguajes
	// pero con la diferencia de que se ejecuta en orden LIFO (Last In, First Out)
	// es decir, el último defer en ser declarado es el primero en ejecutarse.
}

func process() {
	defer fmt.Println("Process ended")
	fmt.Println("Processing...")
	// Simulate some processing work
	for i := range 3 {
		fmt.Println("Working...", i)
	}
	fmt.Println("Process completed")
}

func processWithMultipleDefers() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")

	fmt.Println("Processing with multiple defers...")
}

func processI(i int) {
	// El defer se evalua en el momento de la llamada, por lo que captura i sin el incremento
	defer fmt.Println("Variable i defer:", i)
	i++
	fmt.Println("Processing i:", i)
}
