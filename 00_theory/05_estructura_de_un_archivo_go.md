# Estructura de un Archivo Go

## Anatomía de un Programa Go Básico

Analicemos la estructura de un archivo Go usando como ejemplo nuestro primer programa `hello.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("¡Hola, mundo!")
}
```

## Componentes Fundamentales

### 1. Declaración del Paquete

```go
package main
```

**¿Qué es?**
- Toda archivo Go debe comenzar con una declaración de paquete
- Define a qué paquete pertenece el código del archivo
- `package main` indica que este es un programa ejecutable

**Reglas importantes:**
- `package main` es especial: indica que el archivo contiene un programa que puede ser ejecutado
- Otros nombres de paquete (como `package utils`, `package math`) crean bibliotecas reutilizables
- El nombre del paquete debe coincidir con el nombre de la carpeta (excepto para `main`)

### 2. Declaración de Importaciones

```go
import "fmt"
```

**¿Qué hace?**
- Importa paquetes externos que necesitamos usar en nuestro código
- `fmt` es un paquete estándar de Go que proporciona funciones de formateo e impresión
- Sin esta importación, no podríamos usar `fmt.Println()`

**Variaciones de import:**
```go
// Importación simple
import "fmt"

// Múltiples importaciones
import (
    "fmt"
    "os"
    "strings"
)

// Con alias
import f "fmt"

// Importación en blanco (solo por efectos secundarios)
import _ "database/sql"
```

### 3. Función Principal

```go
func main() {
	fmt.Println("¡Hola, mundo!")
}
```

**Características de la función main:**
- Es el punto de entrada del programa
- Solo puede existir una función `main()` por paquete `main`
- No recibe parámetros ni retorna valores (en su forma básica)
- El programa termina cuando la función `main()` termina

## Estructura General de un Archivo Go

```go
// 1. Declaración del paquete (obligatorio)
package nombre_del_paquete

// 2. Importaciones (opcional)
import (
    "paquete1"
    "paquete2"
)

// 3. Declaraciones a nivel de paquete (opcional)
// - Variables globales
// - Constantes
// - Tipos personalizados
// - Funciones

// 4. Función main (solo en package main)
func main() {
    // Código del programa
}
```

## Ejemplo Expandido

```go
package main

import (
    "fmt"
    "time"
)

// Variable global
var mensaje string = "¡Hola desde Go!"

// Constante
const version = "1.0.0"

// Función personalizada
func saludar(nombre string) {
    fmt.Printf("Hola, %s\n", nombre)
}

// Función principal
func main() {
    fmt.Println(mensaje)
    fmt.Println("Versión:", version)
    
    saludar("Mundo")
    
    fmt.Println("Fecha actual:", time.Now())
}
```

## Reglas de Sintaxis Importantes

### Llaves y Formateo
```go
// Correcto - llave de apertura en la misma línea
func main() {
    fmt.Println("Hola")
}

// Incorrecto - Go no permite esto
func main()
{
    fmt.Println("Hola")
}
```

### Punto y Coma
- Go inserta automáticamente puntos y coma al final de las líneas
- No necesitas escribirlos explícitamente en la mayoría de los casos

### Indentación
- Go usa tabs para indentación por convención
- El formateador `gofmt` automatiza esto

## Convenciones de Nomenclatura

### Nombres de Archivos
- Usa snake_case: `hello_world.go`, `data_types.go`
- La extensión siempre debe ser `.go`

### Nombres de Paquetes
- Usa lowercase: `main`, `fmt`, `utils`
- Deben ser descriptivos y cortos

### Nombres de Funciones y Variables
- Usa camelCase: `miVariable`, `calcularTotal()`
- Si comienza con mayúscula, es exportable (público): `MiFunction()`
- Si comienza con minúscula, es privado al paquete: `miFunction()`

## Herramientas Útiles

### go run
```bash
go run hello.go
```
Compila y ejecuta el archivo directamente

### go build
```bash
go build hello.go
```
Compila el archivo y crea un ejecutable

### gofmt
```bash
gofmt -w hello.go
```
Formatea automáticamente el código según las convenciones de Go

## Puntos Clave para Recordar

1. **Orden obligatorio**: package → import → declaraciones → main
2. **Un solo main()**: Por paquete `main`
3. **Importaciones no usadas**: Go no compila si importas algo que no usas
4. **Variables no usadas**: Go no compila si declaras variables que no usas
5. **Mayúsculas = Público**: Los identificadores que comienzan con mayúscula son exportables
6. **Minúsculas = Privado**: Los identificadores que comienzan con minúscula son privados al paquete

Esta estructura simple pero estricta de Go ayuda a mantener el código limpio, legible y consistente en todos los proyectos.
