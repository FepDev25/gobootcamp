# Slices en Go

Los slices son una abstracción más poderosa y flexible que los arrays en Go. Representan una vista sobre un array subyacente y pueden cambiar de tamaño dinámicamente durante la ejecución del programa.

## ¿Qué es un Slice?

Un slice es una estructura que contiene:
- **Puntero**: Referencia al array subyacente
- **Longitud**: Número de elementos en el slice
- **Capacidad**: Número total de elementos desde el inicio del slice hasta el final del array subyacente

```
slice := []int{1, 2, 3, 4, 5}
       ┌─────────────────────┐
       │ ptr │ len │ cap     │
       │  •  │  5  │  5      │
       └─────┼─────────────────┘
             │
             ▼
       ┌─────┬─────┬─────┬─────┬─────┐
       │  1  │  2  │  3  │  4  │  5  │
       └─────┴─────┴─────┴─────┴─────┘
```

## Declaración e Inicialización

### Slice Vacío

```go
var numbers []int            // Slice vacío (nil)
var names = []string{}       // Slice vacío pero no nil
var scores = make([]int, 0)  // Slice vacío con capacidad 0

fmt.Println(numbers == nil)  // true
fmt.Println(names == nil)    // false
fmt.Println(len(numbers))    // 0
```

### Slice con Valores Iniciales

```go
// Slice literal
numbers := []int{1, 2, 3, 4, 5}
fmt.Println(numbers) // [1 2 3 4 5]

// Slice de strings
fruits := []string{"apple", "banana", "orange"}
fmt.Println(fruits) // [apple banana orange]
```

### Crear Slice con `make`

```go
// make([]type, longitud, capacidad)
numbers := make([]int, 5)      // longitud=5, capacidad=5, [0 0 0 0 0]
numbers2 := make([]int, 3, 10) // longitud=3, capacidad=10, [0 0 0]

fmt.Println("numbers:", numbers, "len:", len(numbers), "cap:", cap(numbers))
fmt.Println("numbers2:", numbers2, "len:", len(numbers2), "cap:", cap(numbers2))
```

### Slice desde Array

```go
array := [6]int{10, 20, 30, 40, 50, 60}
slice := array[1:4] // Desde índice 1 hasta 3 (excluye 4)

fmt.Println("Array:", array)   // [10 20 30 40 50 60]
fmt.Println("Slice:", slice)   // [20 30 40]
```

## Sintaxis de Slicing

```go
slice[bajo:alto]      // Desde bajo hasta alto-1
slice[bajo:]          // Desde bajo hasta el final
slice[:alto]          // Desde el inicio hasta alto-1
slice[:]              // Todo el slice (copia)
```

**Ejemplos:**
```go
numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

fmt.Println(numbers[2:5])    // [2 3 4]
fmt.Println(numbers[3:])     // [3 4 5 6 7 8 9]
fmt.Println(numbers[:4])     // [0 1 2 3]
fmt.Println(numbers[:])      // [0 1 2 3 4 5 6 7 8 9]
```

## Operaciones con Slices

### Función `append`

```go
slice := []int{1, 2, 3}
fmt.Println("Original:", slice)

// Agregar un elemento
slice = append(slice, 4)
fmt.Println("Después de append(4):", slice) // [1 2 3 4]

// Agregar múltiples elementos
slice = append(slice, 5, 6, 7)
fmt.Println("Después de append(5,6,7):", slice) // [1 2 3 4 5 6 7]

// Agregar otro slice
other := []int{8, 9, 10}
slice = append(slice, other...)  // Los ... expanden el slice
fmt.Println("Después de append otro slice:", slice) // [1 2 3 4 5 6 7 8 9 10]
```

### Función `copy`

```go
source := []int{1, 2, 3, 4, 5}
destination := make([]int, len(source))

// Copiar elementos de source a destination
copied := copy(destination, source)
fmt.Printf("Elementos copiados: %d\n", copied)
fmt.Println("Source:", source)           // [1 2 3 4 5]
fmt.Println("Destination:", destination) // [1 2 3 4 5]

// Modificar destination no afecta source
destination[0] = 999
fmt.Println("Source después de modificar destination:", source) // [1 2 3 4 5]
```

### Eliminar Elementos

```go
func removeElement(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

numbers := []int{10, 20, 30, 40, 50}
fmt.Println("Original:", numbers)

// Eliminar elemento en índice 2 (valor 30)
numbers = removeElement(numbers, 2)
fmt.Println("Después de eliminar índice 2:", numbers) // [10 20 40 50]
```

## Comparación de Slices

Los slices NO son directamente comparables con `==`:

```go
slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}

// ❌ Esto NO compila
// fmt.Println(slice1 == slice2) // Error

// ✅ Solo puedes comparar con nil
fmt.Println(slice1 == nil) // false

// ✅ Para comparar contenido, usa slices.Equal (Go 1.21+)
import "slices"
fmt.Println(slices.Equal(slice1, slice2)) // true
```

### Función de Comparación Manual

```go
func equalSlices(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
```

## Slices Bidimensionales

```go
// Crear una matriz usando slices
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

fmt.Println("Matriz:", matrix)
fmt.Println("Elemento [1][2]:", matrix[1][2]) // 6

// Modificar elemento
matrix[0][0] = 100
fmt.Println("Matriz modificada:", matrix)
```

### Crear Matriz Dinámica

```go
rows, cols := 3, 4
matrix := make([][]int, rows)

for i := range matrix {
    matrix[i] = make([]int, cols)
}

// Llenar la matriz
counter := 1
for i := 0; i < rows; i++ {
    for j := 0; j < cols; j++ {
        matrix[i][j] = counter
        counter++
    }
}

fmt.Println("Matriz dinámica:", matrix)
// [[1 2 3 4] [5 6 7 8] [9 10 11 12]]
```

## Iteración sobre Slices

```go
numbers := []int{10, 20, 30, 40, 50}

// Con for tradicional
for i := 0; i < len(numbers); i++ {
    fmt.Printf("Índice: %d, Valor: %d\n", i, numbers[i])
}

// Con range (recomendado)
for index, value := range numbers {
    fmt.Printf("Índice: %d, Valor: %d\n", index, value)
}

// Solo el valor
for _, value := range numbers {
    fmt.Println("Valor:", value)
}
```

## Paquete `slices` (Go 1.21+)

El paquete `slices` proporciona funciones útiles para trabajar con slices:

```go
import "slices"

numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}

// Verificar si contiene un elemento
fmt.Println("Contiene 5:", slices.Contains(numbers, 5)) // true

// Encontrar índice
fmt.Println("Índice de 4:", slices.Index(numbers, 4)) // 2

// Ordenar (modifica el slice original)
slices.Sort(numbers)
fmt.Println("Ordenado:", numbers) // [1 1 2 3 4 5 6 9]

// Invertir
slices.Reverse(numbers)
fmt.Println("Invertido:", numbers) // [9 6 5 4 3 2 1 1]

// Clonar
clone := slices.Clone(numbers)
fmt.Println("Clon:", clone)
```

## Memoria y Rendimiento

### Crecimiento de Capacidad

```go
slice := make([]int, 0, 1)
fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

for i := 0; i < 10; i++ {
    slice = append(slice, i)
    fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}
```

### Slice Leak

Cuidado con mantener referencias a arrays grandes:

```go
// ❌ Problemático: mantiene referencia al array completo
func getFirstThree(data []int) []int {
    return data[:3]  // El array completo permanece en memoria
}

// ✅ Mejor: hacer una copia
func getFirstThreeSafe(data []int) []int {
    result := make([]int, 3)
    copy(result, data[:3])
    return result  // Solo mantiene los 3 elementos necesarios
}
```

## Ejemplos Prácticos

### 1. Filtrar Elementos

```go
func filterEvens(numbers []int) []int {
    var evens []int
    for _, num := range numbers {
        if num%2 == 0 {
            evens = append(evens, num)
        }
    }
    return evens
}

numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evens := filterEvens(numbers)
fmt.Println("Números pares:", evens) // [2 4 6 8 10]
```

### 2. Transformar Elementos

```go
func square(numbers []int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = num * num
    }
    return result
}

numbers := []int{1, 2, 3, 4, 5}
squared := square(numbers)
fmt.Println("Cuadrados:", squared) // [1 4 9 16 25]
```

### 3. Encontrar Duplicados

```go
func findDuplicates(numbers []int) []int {
    seen := make(map[int]bool)
    var duplicates []int
    
    for _, num := range numbers {
        if seen[num] {
            duplicates = append(duplicates, num)
        } else {
            seen[num] = true
        }
    }
    
    return duplicates
}

numbers := []int{1, 2, 3, 2, 4, 3, 5}
duplicates := findDuplicates(numbers)
fmt.Println("Duplicados:", duplicates) // [2 3]
```

## Mejores Prácticas

### 1. Pre-asignar Capacidad

```go
// ❌ Ineficiente: múltiples reasignaciones
var result []int
for i := 0; i < 1000; i++ {
    result = append(result, i)
}

// ✅ Eficiente: capacidad conocida de antemano
result := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    result = append(result, i)
}
```

### 2. Usar `copy` para Duplicar

```go
// ❌ No crea una copia independiente
original := []int{1, 2, 3}
notACopy := original  // Misma referencia

// ✅ Crea una copia independiente
original := []int{1, 2, 3}
independent := make([]int, len(original))
copy(independent, original)
```

### 3. Verificar Nil antes de Usar

```go
func safeProcess(data []int) {
    if data == nil {
        fmt.Println("Slice es nil")
        return
    }
    
    if len(data) == 0 {
        fmt.Println("Slice está vacío")
        return
    }
    
    // Procesar data...
}
```

Los slices son uno de los tipos más utilizados en Go debido a su flexibilidad y eficiencia. Entender cómo funcionan internamente te ayudará a escribir código más eficiente y evitar errores comunes relacionados con referencias compartidas y gestión de memoria.
