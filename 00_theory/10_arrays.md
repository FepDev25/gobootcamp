# Arrays

Los arrays en Go son colecciones de elementos del mismo tipo con tamaño fijo determinado en tiempo de compilación.

## Declaración

Especifica el tipo y la longitud:

```go
var numeros [5]int           // [0 0 0 0 0]
var nombres [3]string        // ["" "" ""]
```

Inicialización con valores:

```go
var numeros = [5]int{1, 2, 3, 4, 5}
nombres := [3]string{"Ana", "Luis", "María"}
```

## Inferencia de Longitud

El compilador cuenta los elementos con `[...]`:

```go
valores := [...]int{10, 20, 30}    // Array de longitud 3
```

## Acceso a Elementos

Los índices comienzan en cero:

```go
numeros := [5]int{10, 20, 30, 40, 50}

primero := numeros[0]    // 10
ultimo := numeros[4]     // 50

numeros[2] = 35          // Modificar elemento
```

Acceso fuera de rango produce error de compilación o pánico en tiempo de ejecución.

## Longitud del Array

La función `len` retorna el número de elementos:

```go
numeros := [5]int{1, 2, 3, 4, 5}
longitud := len(numeros)    // 5
```

## Arrays Multidimensionales

```go
// Matriz 3x3
var matriz [3][3]int

matriz[0][0] = 1
matriz[1][1] = 5

// Inicialización
matriz := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

## Comparación de Arrays

Los arrays son comparables si su tipo elemento es comparable:

```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [3]int{3, 2, 1}

fmt.Println(a == b)    // true
fmt.Println(a == c)    // false
```

Arrays de diferentes longitudes son tipos distintos e incomparables:

```go
var a [3]int
var b [4]int

// a == b    // Error: tipos incompatibles
```

## Copia de Arrays

La asignación crea una copia completa:

```go
original := [3]int{1, 2, 3}
copia := original

copia[0] = 100

fmt.Println(original)    // [1 2 3]
fmt.Println(copia)       // [100 2 3]
```

## Limitaciones

- Tamaño fijo: no puede cambiar después de la declaración
- Pasar a funciones copia todos los elementos (costoso para arrays grandes)
- Por estas razones, se usan poco directamente; los slices son más comunes

## Recorrido con Range

```go
numeros := [5]int{10, 20, 30, 40, 50}

for indice, valor := range numeros {
    fmt.Printf("numeros[%d] = %d\n", indice, valor)
}
```

## Resumen

- Tamaño fijo definido en la declaración: `[n]tipo`
- Los arrays son valores; la asignación copia todos los elementos
- Son comparables si tienen la misma longitud y tipo comparable
- Usar slices para colecciones de tamaño dinámico
- La función `len` retorna el número de elementos
