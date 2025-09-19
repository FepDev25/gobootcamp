package ejercicios

import "fmt"

/*
EJERCICIO 2: Sistema de Inventario con Maps
==========================================
Gestiona un inventario de productos usando maps.

Instrucciones:
1. Define un map donde la clave es el nombre del producto (string) y el valor es la cantidad (int)
2. Implementa las siguientes operaciones que retornen bool indicando éxito:
  - Agregar producto: si ya existe, suma la cantidad
  - Remover producto: elimina del inventario
  - Actualizar cantidad: cambia la cantidad de un producto existente
  - Buscar producto: verifica si existe y retorna la cantidad

3. Implementa función para mostrar reporte del inventario

Ejemplo de uso:
Inventario inicial: map[manzanas:50 naranjas:30 plátanos:25]
Agregar manzanas (20): true → manzanas: 70
Remover naranjas: true
Buscar plátanos: true, cantidad: 25
Inventario final: map[manzanas:70 plátanos:25]
*/
func Ejercicio2() {
	// TODO: Implementa aquí tu solución
	// Crea el inventario inicial y ejecuta las operaciones
	inventario := map[string]int{
		"Platano": 15,
		"Limon":   5,
	}
	fmt.Println("Inventario inicial: ", inventario)

	// --- Agregar ---
	manzana1 := agregarProducto(inventario, "Manzana", 10)
	fmt.Println("Agregar manzanas (10): ", manzana1)

	pera1 := agregarProducto(inventario, "Pera", 5)
	fmt.Println("Agregar peras (5): ", pera1)

	durazno1 := agregarProducto(inventario, "Durazno", 5)
	fmt.Println("Agregar duraznos (5): ", durazno1)

	manzana2 := agregarProducto(inventario, "Manzana", 5)
	fmt.Println("Agregar manzanas (5): ", manzana2)

	// --- Remover ---
	removePera := removerProducto(inventario, "Pera")
	fmt.Println("Eliminar Pera: ", removePera)

	// --- Actualizar cantidad ---
	actManzana := actualizarCantidad(inventario, "Manzana", 20)
	fmt.Println("Actualizar cantidad Manzana (20):", actManzana)

	actSandia := actualizarCantidad(inventario, "Sandia", 20)
	fmt.Println("Actualizar Sandia (20):", actSandia)

	// --- Buscar ---
	encontrado, cantidadDurazno := buscarProducto(inventario, "Durazno")
	fmt.Println("Cantidad del durazno: ", encontrado, cantidadDurazno)

	fmt.Println("Inventario final: ", inventario)

}

func agregarProducto(inventario map[string]int, producto string, cantidad int) bool {
	v, ok := inventario[producto]
	if ok {
		inventario[producto] = v + cantidad
		return false
	} else {
		inventario[producto] = cantidad
		return true
	}
}

func removerProducto(inventario map[string]int, producto string) bool {
	_, ok := inventario[producto]
	if ok {
		delete(inventario, producto)
		return true
	}
	return false

}

func actualizarCantidad(inventario map[string]int, producto string, cantidad int) bool {
	v, ok := inventario[producto]
	if ok {
		inventario[producto] = v + cantidad
		return true
	}
	return false

}

func buscarProducto(inventario map[string]int, producto string) (bool, int) {
	v, ok := inventario[producto]
	return ok, v
}
