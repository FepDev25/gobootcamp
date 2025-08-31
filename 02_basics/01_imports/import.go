package main

import (
	"fmt"
	red "net/http"
)

// Ejemplo de uso de múltiples importaciones
func main() {
	fmt.Println("¡Hola, mundo!")

	resp, err := red.Get("http://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Respuesta del servidor: ", resp.Status)

}
