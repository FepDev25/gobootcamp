package main

import (
	"fmt"
	"maps"
	"reflect"
)

func main() {
	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["Felipe"] = 20
	myMap["Pedro"] = 0
	myMap["Pepe"] = 5

	fmt.Println(myMap)
	fmt.Println("Felipe's age:", myMap["Felipe"])

	delete(myMap, "Felipe")
	fmt.Println("After deletion:", myMap)

	myMap["Felipe"] = 20

	checkIfExists("Felipe", myMap)
	checkIfExists("Juan", myMap)

	clear(myMap)
	fmt.Println("After clearing:", myMap)

	// Declare and initialize in the same time
	myMap2 := map[string]int{
		"Chelsea FC":  2,
		"Real Madrid": 15,
		"Barcelona":   5,
	}
	fmt.Println("myMap2:", myMap2)

	iterateMap(myMap2)

	fmt.Println("Length of myMap2:", len(myMap2))
	keys := maps.Keys(myMap2)
	values := maps.Values(myMap2)
	fmt.Println("Keys:", keys)
	fmt.Println("Values:", values)

	typ := reflect.TypeOf(keys)
	fmt.Println("Type:", typ)

	for key := range keys {
		fmt.Println("Key:", key)
	}

	for value := range values {
		fmt.Println("Value:", value)
	}

	// Nested maps
	myMap3 := make(map[string]map[string]int)

	myMap3["Team A"] = map[string]int{
		"Player 1": 30,
		"Player 2": 25,
	}

	myMap3["Team B"] = map[string]int{
		"Player 3": 28,
		"Player 4": 22,
	}

	fmt.Println("myMap3:", myMap3)
	fmt.Println(myMap3["Team A"])
}

func checkIfExists(key string, myMap map[string]int) {
	valor, valorDesconocido := myMap[key]
	if valorDesconocido {
		fmt.Println(key+"'s age:", valor)
	} else {
		fmt.Println(key, "not found")
	}
}

func iterateMap(mapa map[string]int) {
	for key, value := range mapa {
		fmt.Println(key, ":", value)
	}
}
