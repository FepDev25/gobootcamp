package main

import (
	"fmt"
)

func main() {

	fmt.Println("Range in array")
	ages := [3]int{19, 18, 20}
	for index, value := range ages {
		fmt.Println(index, value)
	}

	fmt.Println("Range in slice")
	agesSlice := []int{19, 18, 20}
	for index, value := range agesSlice {
		fmt.Println(index, value)
	}

	fmt.Println("Range in string")
	message := "Hola mundo"
	for i, v := range message {
		fmt.Printf("Index: %d, Value unicode: %d, Rune: %c\n", i, v, v) // Indice y valor unicode
	}

	fmt.Println("Range in map")
	agesMap := map[string]int{"Felipe": 19, "Juan": 18, "Maria": 20}
	for key, value := range agesMap {
		fmt.Println(key, value)
	}

}
