package main

import "fmt"

func main() {
	num := 12
	var p *int
	p = &num

	fmt.Println(num) // Valor de la variable num
	fmt.Println(p)   // Dirección de memoria de la variable num
	fmt.Println(*p)  // Valor de la variable a la que apunta p

	fmt.Println("Haciendo cambios en puntero...")

	*p = 22
	fmt.Println(num) // Nuevo valor de la variable num
	fmt.Println(p)   // Dirección de memoria de la variable num
	fmt.Println(*p)  // Valor de la variable a la que apunta p

	numeros := []int{}
	numeros = append(numeros, 10, 11, 12)
	fmt.Println(numeros)

	edades := []int{19, 20, 21, 18, 19}
	nombres := []string{"Felipe", "Banana", "Bananado", "Pera", "Banaismo"}

	diccionario := make(map[string]int)

	for i := 0; i < len(nombres); i++ {
		diccionario[nombres[i]] = edades[i]
	}

	fmt.Println(diccionario)

}
