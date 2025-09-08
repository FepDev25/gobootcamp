# Variables en Go

Las variables en Go son espacios de memoria con nombre que almacenan valores de un tipo específico. Go ofrece varias formas de declarar e inicializar variables.

## Declaración de Variables

### 1. Declaración con `var`

```go
var nombre string        // Declaración sin inicialización
var edad int = 25        // Declaración con inicialización
var altura = 1.75        // Inferencia de tipo
```

### 2. Declaración Corta con `:=`

La forma más común de declarar variables en Go es usando el operador `:=` (declaración corta):

```go
nombre := "Felipe"       // string
edad := 25              // int
activo := true          // bool
```

**Nota:** El operador `:=` solo puede usarse dentro de funciones, no en el ámbito global del paquete.

### 3. Declaración Múltiple

```go
// Con var
var nombre, apellido string
var x, y, z int = 1, 2, 3

// Con :=
nombre, edad := "Felipe", 25
```

## Ámbito (Scope) de Variables

### Variables Globales
Las variables declaradas fuera de cualquier función tienen ámbito global dentro del paquete:

```go
package main

import "fmt"

var petName = "Pedrito" // Ámbito global

func main() {
    // Puede usar petName aquí
    fmt.Println(petName)
}
```

### Variables Locales
Las variables declaradas dentro de una función tienen ámbito local:

```go
func printName(name string) {
    uppName := strings.ToUpper(name) // Ámbito local
    fmt.Println(uppName)
    // uppName no existe fuera de esta función
}
```

## Valor Cero

En Go, todas las variables tienen un "valor cero" por defecto:

| Tipo | Valor Cero |
|------|------------|
| `int` | `0` |
| `float64` | `0.0` |
| `string` | `""` (cadena vacía) |
| `bool` | `false` |
| `pointer` | `nil` |
| `slice` | `nil` |
| `map` | `nil` |

## Buenas Prácticas

1. **Usa nombres descriptivos**: `edad` en lugar de `e`
2. **Prefiere `:=` para variables locales** por su concisión
3. **Declara variables cerca de donde se usan**
4. **Usa `var` para variables globales o cuando necesites un valor cero específico**
5. **Evita variables globales cuando sea posible**

## Ejemplo Completo

```go
package main

import (
    "fmt"
    "strings"
)

var petName = "Pedrito" // Variable global

func main() {
    // Diferentes formas de declarar variables
    var age int                // Valor cero: 0
    var name string = "Felipe" // Inicialización explícita
    var name2 = "Juan"         // Inferencia de tipo
    count := 1                 // Declaración corta
    lastname := "Peralta"      // Declaración corta
    
    fmt.Println(name, name2, lastname, age, count)
    
    // Uso de variable global
    printName(petName)
}

func printName(name string) {
    uppName := strings.ToUpper(name) // Variable local
    fmt.Println(uppName)
}
```

Las variables son fundamentales en cualquier programa Go y entender su ámbito y formas de declaración es esencial para escribir código limpio y eficiente.
