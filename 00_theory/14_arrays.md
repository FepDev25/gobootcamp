# Arrays en Go

Los arrays en Go son colecciones de elementos del mismo tipo con un tamaño fijo que se define en tiempo de compilación. Son la base para tipos de datos más flexibles como los slices.

## Declaración e Inicialización

### Declaración Básica

```go
// Declarar un array de 5 enteros (inicializado con valores cero)
var numbers [5]int
fmt.Println(numbers) // [0 0 0 0 0]

// Declarar e inicializar con valores específicos
var fruits [4]string{"Apple", "Banana", "Grapes", "Orange"}
fmt.Println(fruits) // [Apple Banana Grapes Orange]
```

### Inicialización con `{}`

```go
// Array literal
numbers := [6]int{42, 27, 33, 19, 55, 78}
fmt.Println(numbers) // [42 27 33 19 55 78]

// Inicialización parcial (resto se llena con valores cero)
partial := [5]int{1, 2}
fmt.Println(partial) // [1 2 0 0 0]
```

### Array con Tamaño Inferido

```go
// El compilador calcula el tamaño basado en los elementos
auto := [...]int{10, 20, 30, 40, 50}
fmt.Printf("Tamaño: %d, Array: %v\n", len(auto), auto)
// Tamaño: 5, Array: [10 20 30 40 50]
```

### Inicialización con Índices Específicos

```go
// Inicializar posiciones específicas
sparse := [10]int{0: 1, 5: 10, 9: 100}
fmt.Println(sparse) // [1 0 0 0 0 10 0 0 0 100]
```

## Operaciones con Arrays

### Acceder a Elementos

```go
fruits := [4]string{"Apple", "Banana", "Grapes", "Orange"}

// Acceder por índice
fmt.Println(fruits[2]) // Grapes

// Modificar elementos
fruits[1] = "Banana (Updated)"
fmt.Println(fruits[1]) // Banana (Updated)
```

### Propiedades del Array

```go
numbers := [5]int{1, 2, 3, 4, 5}

// Longitud del array
fmt.Println("Longitud:", len(numbers)) // 5

// Capacidad del array (siempre igual a la longitud)
fmt.Println("Capacidad:", cap(numbers)) // 5
```

## Iteración sobre Arrays

### Usando `for` tradicional

```go
numbers := [5]int{10, 20, 30, 40, 50}

for i := 0; i < len(numbers); i++ {
    fmt.Printf("Índice: %d, Valor: %d\n", i, numbers[i])
}
```

### Usando `for range`

```go
numbers := [5]int{10, 20, 30, 40, 50}

// Con índice y valor
for index, value := range numbers {
    fmt.Printf("Índice: %d, Valor: %d\n", index, value)
}

// Solo el valor
for _, value := range numbers {
    fmt.Println("Valor:", value)
}

// Solo el índice
for index := range numbers {
    fmt.Println("Índice:", index)
}
```

## Comparación de Arrays

Los arrays son comparables si son del mismo tipo y tamaño:

```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [3]int{1, 2, 4}

fmt.Println(a == b) // true
fmt.Println(a == c) // false

// Arrays de diferentes tamaños NO son comparables
// var d [4]int{1, 2, 3, 4}
// fmt.Println(a == d) // Error de compilación
```

## Copia de Arrays

### Copia por Valor

Los arrays se copian por valor en Go:

```go
original := [5]int{2, 4, 6, 8, 10}
copy := original    // Copia completa del array
copy[4] = 100       // Modificar la copia

fmt.Println("Original:", original) // [2 4 6 8 10]
fmt.Println("Copia:", copy)        // [2 4 6 8 100]
```

### Copia con Punteros

Para modificar el array original, usa punteros:

```go
original := [5]int{2, 4, 6, 8, 10}
ptr := &original    // Puntero al array original
ptr[4] = 100        // Modificar a través del puntero

fmt.Println("Original:", original) // [2 4 6 8 100]
```

## Arrays Multidimensionales

### Array 2D (Matriz)

```go
// Declarar una matriz 3x3
var matrix [3][3]int

// Inicializar con valores
matrix2 := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Acceder a elementos
fmt.Println("Elemento [1][2]:", matrix2[1][2]) // 6

// Modificar elementos
matrix2[0][0] = 100
```

### Iteración en Arrays 2D

```go
matrix := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Iterar con índices
for i := 0; i < len(matrix); i++ {
    for j := 0; j < len(matrix[i]); j++ {
        fmt.Printf("matrix[%d][%d] = %d ", i, j, matrix[i][j])
    }
    fmt.Println()
}

// Iterar con range
for i, row := range matrix {
    for j, value := range row {
        fmt.Printf("matrix[%d][%d] = %d ", i, j, value)
    }
    fmt.Println()
}
```

## Funciones con Arrays

### Pasar Arrays a Funciones

Los arrays se pasan por valor (se copian):

```go
func processArray(arr [5]int) {
    arr[0] = 999  // Modifica solo la copia local
    fmt.Println("Dentro de función:", arr)
}

func main() {
    numbers := [5]int{1, 2, 3, 4, 5}
    processArray(numbers)
    fmt.Println("Array original:", numbers) // [1 2 3 4 5] (sin cambios)
}
```

### Pasar Punteros a Arrays

Para modificar el array original:

```go
func modifyArray(arr *[5]int) {
    arr[0] = 999  // Modifica el array original
    fmt.Println("Dentro de función:", *arr)
}

func main() {
    numbers := [5]int{1, 2, 3, 4, 5}
    modifyArray(&numbers)
    fmt.Println("Array original:", numbers) // [999 2 3 4 5] (modificado)
}
```

## Ejemplos Prácticos

### 1. Búsqueda Linear

```go
func linearSearch(arr [10]int, target int) int {
    for i, value := range arr {
        if value == target {
            return i
        }
    }
    return -1  // No encontrado
}

func main() {
    numbers := [10]int{5, 2, 8, 1, 9, 3, 7, 4, 6, 0}
    index := linearSearch(numbers, 7)
    
    if index != -1 {
        fmt.Printf("Encontrado en índice: %d\n", index)
    } else {
        fmt.Println("No encontrado")
    }
}
```

### 2. Encontrar Máximo y Mínimo

```go
func findMinMax(arr [8]int) (min, max int) {
    min, max = arr[0], arr[0]
    
    for _, value := range arr {
        if value < min {
            min = value
        }
        if value > max {
            max = value
        }
    }
    
    return min, max
}

func main() {
    numbers := [8]int{45, 23, 78, 12, 67, 34, 89, 56}
    min, max := findMinMax(numbers)
    fmt.Printf("Mínimo: %d, Máximo: %d\n", min, max)
}
```

### 3. Suma de Elementos

```go
func sumArray(arr [6]int) int {
    sum := 0
    for _, value := range arr {
        sum += value
    }
    return sum
}

func main() {
    numbers := [6]int{10, 20, 30, 40, 50, 60}
    total := sumArray(numbers)
    fmt.Println("Suma total:", total) // 210
}
```

## Limitaciones de los Arrays

1. **Tamaño fijo**: No pueden cambiar de tamaño en tiempo de ejecución
2. **Tipo específico**: El tamaño es parte del tipo
3. **Copia por valor**: Pueden ser costosos de pasar a funciones
4. **Menos flexibles**: Comparados con slices

```go
// Estos son tipos DIFERENTES
var arr1 [5]int
var arr2 [6]int
// No puedes asignar arr1 = arr2 (error de compilación)
```

## Arrays vs Slices

| Característica | Arrays | Slices |
|---------------|--------|--------|
| **Tamaño** | Fijo | Dinámico |
| **Memoria** | Valor | Referencia |
| **Flexibilidad** | Limitada | Alta |
| **Uso común** | Raro | Frecuente |

**Recomendación**: En la mayoría de casos, usa slices en lugar de arrays por su flexibilidad.

## Ejemplo Completo: Inventario

```go
package main

import "fmt"

const MAX_ITEMS = 10

type Inventory struct {
    items [MAX_ITEMS]string
    count int
}

func (inv *Inventory) addItem(item string) bool {
    if inv.count >= MAX_ITEMS {
        return false  // Inventario lleno
    }
    
    inv.items[inv.count] = item
    inv.count++
    return true
}

func (inv *Inventory) listItems() {
    fmt.Println("Inventario:")
    for i := 0; i < inv.count; i++ {
        fmt.Printf("%d. %s\n", i+1, inv.items[i])
    }
}

func main() {
    var inventory Inventory
    
    inventory.addItem("Espada")
    inventory.addItem("Escudo")
    inventory.addItem("Poción")
    
    inventory.listItems()
}
```

Los arrays en Go son fundamentales pero se usan menos frecuentemente que los slices en aplicaciones reales. Son útiles cuando necesitas un tamaño fijo conocido en tiempo de compilación y quieres las garantías de rendimiento que ofrecen.
