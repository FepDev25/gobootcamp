# Range en Go

El `range` es una palabra clave en Go que se utiliza para iterar sobre diferentes tipos de datos de manera elegante y eficiente. Es fundamental para recorrer arrays, slices, maps, strings, channels y otros tipos iterables.

## Sintaxis Básica

```go
for índice, valor := range iterable {
    // procesar índice y valor
}

for índice := range iterable {
    // procesar solo índice
}

for _, valor := range iterable {
    // procesar solo valor
}
```

## Range sobre Arrays y Slices

### Sintaxis Completa

```go
numbers := []int{10, 20, 30, 40, 50}

// Con índice y valor
for index, value := range numbers {
    fmt.Printf("Índice: %d, Valor: %d\n", index, value)
}
// Salida:
// Índice: 0, Valor: 10
// Índice: 1, Valor: 20
// ...
```

### Solo el Valor

```go
numbers := []int{10, 20, 30, 40, 50}

// Ignorar índice con _
for _, value := range numbers {
    fmt.Printf("Valor: %d\n", value)
}
```

### Solo el Índice

```go
numbers := []int{10, 20, 30, 40, 50}

// Solo índice
for index := range numbers {
    fmt.Printf("Índice: %d, Valor: %d\n", index, numbers[index])
}
```

### Modificar durante Iteración

```go
numbers := []int{1, 2, 3, 4, 5}

// ❌ Esto NO modifica el slice original
for _, value := range numbers {
    value *= 2  // Solo modifica la copia local
}
fmt.Println("Sin cambios:", numbers) // [1 2 3 4 5]

// ✅ Usar índice para modificar
for i := range numbers {
    numbers[i] *= 2  // Modifica el slice original
}
fmt.Println("Modificado:", numbers) // [2 4 6 8 10]
```

## Range sobre Maps

### Iterar Claves y Valores

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}

for name, age := range ages {
    fmt.Printf("%s tiene %d años\n", name, age)
}
```

### Solo Claves

```go
for name := range ages {
    fmt.Printf("Nombre: %s\n", name)
}
```

### Solo Valores

```go
for _, age := range ages {
    fmt.Printf("Edad: %d\n", age)
}
```

**Importante**: El orden de iteración en maps es **aleatorio**.

## Range sobre Strings

### Iteración por Caracteres (Runes)

```go
text := "Hola 世界"

for index, char := range text {
    fmt.Printf("Posición %d: %c (Unicode: %U)\n", index, char, char)
}
// Salida:
// Posición 0: H (Unicode: U+0048)
// Posición 1: o (Unicode: U+006F)
// Posición 2: l (Unicode: U+006C)
// Posición 3: a (Unicode: U+0061)
// Posición 4:   (Unicode: U+0020)
// Posición 5: 世 (Unicode: U+4E16)
// Posición 8: 界 (Unicode: U+754C)
```

**Nota**: El índice salta posiciones porque los caracteres Unicode pueden ocupar múltiples bytes.

### Solo Caracteres

```go
text := "Go es genial"
for _, char := range text {
    fmt.Printf("%c ", char)
}
// Salida: G o   e s   g e n i a l
```

## Range sobre Channels

```go
func rangeOverChannel() {
    ch := make(chan int, 3)
    
    // Enviar datos
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch) // Importante: cerrar el channel
    
    // Iterar hasta que se cierre
    for value := range ch {
        fmt.Println("Recibido:", value)
    }
}
```

## Range sobre Enteros (Go 1.22+)

```go
// Iterar de 0 a n-1
for i := range 5 {
    fmt.Println("i =", i)
}
// Salida: 0, 1, 2, 3, 4

// Útil para bucles simples
for i := range 10 {
    fmt.Printf("Número: %d\n", i)
}
```

## Patrones Comunes con Range

### 1. Encontrar Elemento

```go
func findIndex(slice []string, target string) int {
    for index, value := range slice {
        if value == target {
            return index
        }
    }
    return -1 // No encontrado
}

fruits := []string{"apple", "banana", "orange"}
index := findIndex(fruits, "banana")
fmt.Println("Índice de banana:", index) // 1
```

### 2. Filtrar Elementos

```go
func filterEvens(numbers []int) []int {
    var result []int
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}

numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evens := filterEvens(numbers)
fmt.Println("Pares:", evens) // [2 4 6 8 10]
```

### 3. Transformar Elementos

```go
func toUpperCase(words []string) []string {
    result := make([]string, len(words))
    for i, word := range words {
        result[i] = strings.ToUpper(word)
    }
    return result
}

words := []string{"hello", "world", "go"}
upper := toUpperCase(words)
fmt.Println("Mayúsculas:", upper) // [HELLO WORLD GO]
```

### 4. Suma de Elementos

```go
func sum(numbers []int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

numbers := []int{1, 2, 3, 4, 5}
total := sum(numbers)
fmt.Println("Suma:", total) // 15
```

### 5. Contar Elementos

```go
func countOccurrences(slice []string, target string) int {
    count := 0
    for _, value := range slice {
        if value == target {
            count++
        }
    }
    return count
}

words := []string{"go", "is", "great", "go", "is", "awesome"}
count := countOccurrences(words, "go")
fmt.Println("'go' aparece", count, "veces") // 2
```

## Range con Diferentes Tipos

### Struct con Slices

```go
type Team struct {
    Name    string
    Members []string
}

team := Team{
    Name:    "Development",
    Members: []string{"Alice", "Bob", "Carol"},
}

fmt.Printf("Equipo: %s\n", team.Name)
for index, member := range team.Members {
    fmt.Printf("  %d. %s\n", index+1, member)
}
```

### Slice de Structs

```go
type Product struct {
    Name  string
    Price float64
}

products := []Product{
    {"Laptop", 999.99},
    {"Mouse", 29.99},
    {"Keyboard", 79.99},
}

for index, product := range products {
    fmt.Printf("%d. %s: $%.2f\n", index+1, product.Name, product.Price)
}
```

## Consideraciones de Rendimiento

### 1. Copia vs Referencia

```go
// Para structs grandes, considera usar índice
type LargeStruct struct {
    data [1000]int
    // muchos campos...
}

largeSlice := make([]LargeStruct, 100)

// ❌ Copia cada struct completo
for _, item := range largeSlice {
    // item es una copia completa de LargeStruct
    processLargeStruct(item)
}

// ✅ Mejor: usar índice para evitar copias
for i := range largeSlice {
    // largeSlice[i] accede al original
    processLargeStruct(largeSlice[i])
}

// ✅ O usar punteros
for i := range largeSlice {
    processLargeStructPtr(&largeSlice[i])
}
```

### 2. Range sobre Maps Grandes

```go
// Para maps muy grandes, considera procesar en lotes
largeMap := make(map[string]int)
// ... llenar con muchos elementos

batchSize := 100
count := 0
for key, value := range largeMap {
    processKeyValue(key, value)
    count++
    
    if count%batchSize == 0 {
        // Punto de control cada 100 elementos
        fmt.Printf("Procesados %d elementos\n", count)
    }
}
```

## Errores Comunes

### 1. Captura de Variable en Goroutines

```go
// ❌ Problemático
values := []int{1, 2, 3, 4, 5}
for _, v := range values {
    go func() {
        fmt.Println(v) // Posiblemente imprime 5 cinco veces
    }()
}

// ✅ Correcto
for _, v := range values {
    go func(val int) {
        fmt.Println(val) // Imprime valores correctos
    }(v)
}
```

### 2. Modificar Slice Durante Iteración

```go
// ❌ Problemático
numbers := []int{1, 2, 3, 4, 5}
for i, num := range numbers {
    if num%2 == 0 {
        // Modificar slice durante iteración puede causar problemas
        numbers = append(numbers[:i], numbers[i+1:]...)
    }
}

// ✅ Mejor: crear nuevo slice
var odds []int
for _, num := range numbers {
    if num%2 != 0 {
        odds = append(odds, num)
    }
}
```

### 3. Reutilizar Variable de Range

```go
// ❌ Problemático
var ptrs []*int
for _, v := range []int{1, 2, 3} {
    ptrs = append(ptrs, &v) // Todos apuntan al mismo &v
}

// ✅ Correcto
var ptrs []*int
for _, v := range []int{1, 2, 3} {
    temp := v
    ptrs = append(ptrs, &temp)
}
```

## Alternativas a Range

### For Tradicional

```go
// A veces el for tradicional es más claro
slice := []int{1, 2, 3, 4, 5}

// Iterar hacia atrás
for i := len(slice) - 1; i >= 0; i-- {
    fmt.Println(slice[i])
}

// Iterar cada dos elementos
for i := 0; i < len(slice); i += 2 {
    fmt.Println(slice[i])
}
```

## Ejemplo Práctico Completo

```go
package main

import (
    "fmt"
    "strings"
)

type Student struct {
    Name   string
    Grades []int
}

func main() {
    students := []Student{
        {"Alice", []int{90, 85, 92, 88}},
        {"Bob", []int{78, 82, 75, 80}},
        {"Carol", []int{95, 90, 93, 89}},
    }
    
    // Procesar cada estudiante
    for index, student := range students {
        fmt.Printf("Estudiante #%d: %s\n", index+1, student.Name)
        
        // Calcular promedio
        total := 0
        for _, grade := range student.Grades {
            total += grade
        }
        average := float64(total) / float64(len(student.Grades))
        
        fmt.Printf("  Calificaciones: %v\n", student.Grades)
        fmt.Printf("  Promedio: %.1f\n", average)
        
        // Determinar letra de calificación
        var letterGrade string
        switch {
        case average >= 90:
            letterGrade = "A"
        case average >= 80:
            letterGrade = "B"
        case average >= 70:
            letterGrade = "C"
        default:
            letterGrade = "F"
        }
        
        fmt.Printf("  Calificación: %s\n\n", letterGrade)
    }
    
    // Estadísticas generales
    allGrades := make(map[string]int)
    for _, student := range students {
        // Procesar el nombre
        for _, char := range strings.ToLower(student.Name) {
            if char >= 'a' && char <= 'z' {
                allGrades[string(char)]++
            }
        }
    }
    
    fmt.Println("Frecuencia de letras en nombres:")
    for letter, count := range allGrades {
        fmt.Printf("  %s: %d\n", letter, count)
    }
}
```

El `range` es una herramienta poderosa y elegante en Go que simplifica significativamente la iteración sobre diferentes tipos de datos. Su uso idiomático hace que el código sea más legible y menos propenso a errores comunes como índices fuera de límites.
