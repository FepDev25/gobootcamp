package main

import "fmt"

func main() {

	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{10, 20, 30}
	fmt.Println("Numbers1:", numbers1)
	fmt.Println("Numbers2:", numbers2)

	// Llamado a la función con un slice
	fmt.Println("Sum of numbers1:", sum(numbers1...))
	fmt.Println("Sum of numbers2:", sum(numbers2...))

	// Llamado a la función con múltiples argumentos
	fmt.Println("Sum of individual numbers:", sum(1, 2, 3, 4, 5, 6))

	// Llamado a la función con strings
	greet("Alice", "Bob", "Charlie")
	greet("Dave")

	// Llamado a la función con un parámetro normal y parámetros variádicos
	playerAndClubes("Lionel Messi", "FC Barcelona", "Paris Saint-Germain", "Inter Miami CF")
	playerAndClubes("Cristiano Ronaldo", "Sporting CP", "Manchester United", "Real Madrid", "Juventus", "Al Nassr")

	// Llamado a la función que multiplica una secuencia por un factor
	result := multiplicarSecuenciaPorFactor(3, 1, 2, 3, 4, 5) // 3 es el factor
	fmt.Println("Multiplying sequence by factor 3:", result)

	result = multiplicarSecuenciaPorFactor(5, 10, 20, 30)
	fmt.Println("Multiplying sequence by factor 5:", result)

}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func greet(messages ...string) {
	for _, message := range messages {
		fmt.Println("Hello", message)
	}
}

func playerAndClubes(player string, clubs ...string) {
	for _, club := range clubs {
		fmt.Printf("%s has played for %s\n", player, club)
	}
}

func multiplicarSecuenciaPorFactor(factor int, numeros ...int) []int {
	resultados := make([]int, len(numeros))
	for i, numero := range numeros {
		resultados[i] = numero * factor
	}
	return resultados
}
