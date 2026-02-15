# Slices

Los slices son secuencias de longitud variable que proporcionan acceso flexible a colecciones de elementos. Son más comunes que los arrays en código Go.

## Declaración e Inicialización

```go
// Slice vacío (nil)
var numeros []int

// Con literal
numeros := []int{1, 2, 3, 4, 5}

// Con make: make(tipo, longitud, capacidad)
numeros := make([]int, 5)        // Longitud 5, capacidad 5
numeros := make([]int, 5, 10)    // Longitud 5, capacidad 10
```

## Slice vs Array

```go
// Array: tamaño fijo parte del tipo
array := [5]int{1, 2, 3, 4, 5}

// Slice: tamaño dinámico
slice := []int{1, 2, 3, 4, 5}
```

## Creación desde Array

```go
array := [5]int{10, 20, 30, 40, 50}

// Slice completo
slice := array[:]           // [10 20 30 40 50]

// Sub-slice: slice[inicio:fin] (fin es exclusivo)
parcial := array[1:4]       // [20 30 40]
inicio := array[:3]         // [10 20 30]
fin := array[2:]            // [30 40 50]
```

## Longitud y Capacidad

- **Longitud (`len`)**: Número de elementos actuales
- **Capacidad (`cap`)**: Número máximo de elementos sin realocar

```go
slice := make([]int, 3, 5)
fmt.Println(len(slice))    // 3
fmt.Println(cap(slice))    // 5
```

## Modificación de Slices

Los slices son referencias a arrays subyacentes. Modificar un slice afecta el original:

```go
original := []int{10, 20, 30, 40, 50}
referencia := original[1:4]

referencia[0] = 99
fmt.Println(original)      // [10 99 30 40 50]
```

## Agregar Elementos: append

La función `append` agrega elementos, creando un nuevo array si es necesario:

```go
slice := []int{1, 2, 3}
slice = append(slice, 4)           // [1 2 3 4]
slice = append(slice, 5, 6)        // [1 2 3 4 5 6]

// Concatenar slices
otro := []int{7, 8, 9}
slice = append(slice, otro...)     // [1 2 3 4 5 6 7 8 9]
```

## Copiar Slices

`copy` transfiere elementos entre slices:

```go
origen := []int{1, 2, 3, 4, 5}
destino := make([]int, 3)

copiados := copy(destino, origen)    // copiados = 3
destino                              // [1 2 3]
```

## Eliminar Elementos

Go no tiene función de eliminación. Se logra con slicing:

```go
// Eliminar elemento en índice i
slice = append(slice[:i], slice[i+1:]...)

// Ejemplo: eliminar índice 2
numeros := []int{1, 2, 3, 4, 5}
numeros = append(numeros[:2], numeros[3:]...)    // [1 2 4 5]
```

## Nil Slices

Un slice no inicializado es `nil`, diferente de vacío:

```go
var nilSlice []int          // nil
emptySlice := []int{}       // vacío, no nil

fmt.Println(nilSlice == nil)     // true
fmt.Println(len(nilSlice))       // 0
```

Ambos funcionan con `append` y `len`.

## Resumen

- Los slices son referencias a arrays con longitud y capacidad
- Sintaxis de slicing: `[inicio:fin]` donde fin es exclusivo
- `append` agrega elementos; realoca si excede la capacidad
- `copy` transfiere elementos entre slices
- Modificar un slice afecta el array subyacente
- Use `make(tipo, len, cap)` para preasignar capacidad conocida
