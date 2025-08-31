# Tipos de Datos en Go

Go es un lenguaje fuertemente tipado que ofrece una variedad de tipos de datos primitivos y compuestos. Cada tipo tiene características específicas que los hacen apropiados para diferentes casos de uso.

## Tipos Primitivos

### Integer (Enteros)

Los enteros representan números sin parte decimal. Go ofrece varios tamaños para optimizar el uso de memoria.

#### Tipos de Enteros con Signo:
- `int8`: -128 a 127 (8 bits)
- `int16`: -32,768 a 32,767 (16 bits)
- `int32`: -2,147,483,648 a 2,147,483,647 (32 bits)
- `int64`: -9,223,372,036,854,775,808 a 9,223,372,036,854,775,807 (64 bits)
- `int`: Tamaño depende de la arquitectura (32 o 64 bits)

#### Tipos de Enteros sin Signo:
- `uint8` (alias `byte`): 0 a 255 (8 bits)
- `uint16`: 0 a 65,535 (16 bits)
- `uint32`: 0 a 4,294,967,295 (32 bits)
- `uint64`: 0 a 18,446,744,073,709,551,615 (64 bits)
- `uint`: Tamaño depende de la arquitectura (32 o 64 bits)
- `uintptr`: Entero sin signo para almacenar punteros

```go
var edad int = 25
var poblacion uint64 = 7800000000
var letra byte = 'A' // byte es alias de uint8
```

### Float (Números de Punto Flotante)

Representan números con parte decimal usando el estándar IEEE 754.

- `float32`: Precisión simple (32 bits, ~7 dígitos decimales)
- `float64`: Precisión doble (64 bits, ~15 dígitos decimales)

```go
var precio float32 = 19.99
var pi float64 = 3.141592653589793
var temperatura float64 = -15.5
```

### Complex Numbers (Números Complejos)

Para representar números complejos con parte real e imaginaria.

- `complex64`: Componentes real e imaginaria como float32
- `complex128`: Componentes real e imaginaria como float64

```go
var z1 complex64 = 1 + 2i
var z2 complex128 = complex(3.0, 4.0) // 3 + 4i
```

### Booleans (Booleanos)

Representan valores de verdad: `true` o `false`.

- `bool`: Solo puede ser `true` o `false`

```go
var esVerdadero bool = true
var esFalso bool = false
var resultado bool = (5 > 3) // true
```

### Strings (Cadenas de Caracteres)

Secuencias inmutables de bytes, típicamente texto codificado en UTF-8.

- `string`: Cadena de caracteres inmutable

```go
var nombre string = "Juan Pérez"
var saludo string = `Hola,
mundo multilinea`
var emoji string = "🚀"
```

#### Características importantes de strings:
- Inmutables: no puedes cambiar caracteres individuales
- UTF-8 por defecto
- Pueden contener cualquier byte válido

## Tipos Compuestos

### Constants (Constantes)

Valores que no pueden cambiar durante la ejecución del programa.

```go
const PI float64 = 3.14159
const EMPRESA string = "Mi Empresa"
const (
    LUNES = iota    // 0
    MARTES         // 1
    MIERCOLES      // 2
)
```

### Arrays (Arreglos)

Colección de elementos del mismo tipo con tamaño fijo determinado en tiempo de compilación.

```go
var numeros [5]int = [5]int{1, 2, 3, 4, 5}
var nombres [3]string = [3]string{"Ana", "Luis", "Carlos"}
var matriz [2][3]int // Array bidimensional
```

#### Características:
- Tamaño fijo
- Elementos del mismo tipo
- Acceso por índice (base 0)
- Valor por defecto: zero value del tipo

### Structs (Estructuras)

Tipos personalizados que agrupan diferentes campos de datos.

```go
type Persona struct {
    Nombre string
    Edad   int
    Email  string
}

var p1 Persona = Persona{
    Nombre: "María",
    Edad:   30,
    Email:  "maria@email.com",
}
```

#### Características:
- Agrupan campos relacionados
- Pueden tener métodos asociados
- Soporte para embedding (composición)

### Pointers (Punteros)

Variables que almacenan la dirección de memoria de otra variable.

```go
var x int = 42
var ptr *int = &x  // ptr apunta a la dirección de x
var valor int = *ptr // Desreferencia: obtiene el valor (42)
```

#### Operadores:
- `&`: Obtiene la dirección de una variable
- `*`: Desreferencia un puntero (obtiene el valor)

### Maps (Mapas)

Colección de pares clave-valor, similar a hash tables o diccionarios.

```go
var edades map[string]int = make(map[string]int)
edades["Juan"] = 25
edades["María"] = 30

// Declaración e inicialización
colores := map[string]string{
    "rojo":  "#FF0000",
    "verde": "#00FF00",
    "azul":  "#0000FF",
}
```

#### Características:
- Claves únicas
- Acceso rápido O(1) promedio
- Iteración no garantiza orden

### Slices (Rebanadas)

Abstracción sobre arrays que proporciona una vista dinámica y flexible.

```go
var numeros []int = []int{1, 2, 3, 4, 5}
var frutas []string = make([]string, 0, 10) // longitud 0, capacidad 10

// Operaciones comunes
numeros = append(numeros, 6, 7, 8)
subcorte := numeros[1:4] // elementos del índice 1 al 3
```

#### Características:
- Tamaño dinámico
- Referencia a un array subyacente
- Tres componentes: puntero, longitud, capacidad

### Functions (Funciones)

Las funciones son tipos de primera clase en Go.

```go
// Declaración de función
func sumar(a, b int) int {
    return a + b
}

// Función como variable
var operacion func(int, int) int = sumar

// Función anónima
multiplicar := func(x, y int) int {
    return x * y
}
```

#### Características:
- Pueden ser asignadas a variables
- Pueden ser pasadas como parámetros
- Soporte para closures
- Múltiples valores de retorno

### Channels (Canales)

Mecanismo para comunicación y sincronización entre goroutines.

```go
// Canal para enviar/recibir enteros
var ch chan int = make(chan int)

// Canal con buffer
var bufferedCh chan string = make(chan string, 10)

// Envío y recepción
go func() {
    ch <- 42        // Enviar
}()
valor := <-ch       // Recibir
```

#### Tipos de canales:
- `chan T`: Bidireccional
- `chan<- T`: Solo envío
- `<-chan T`: Solo recepción

## Tipos Especiales para Datos

### JSON

Go tiene excelente soporte para trabajar con JSON a través del paquete `encoding/json`.

```go
import "encoding/json"

type Usuario struct {
    Nombre string `json:"nombre"`
    Edad   int    `json:"edad"`
}

// Serialización (struct a JSON)
usuario := Usuario{Nombre: "Ana", Edad: 25}
jsonData, err := json.Marshal(usuario)

// Deserialización (JSON a struct)
var nuevoUsuario Usuario
err = json.Unmarshal(jsonData, &nuevoUsuario)
```

### Text (Texto)

Paquetes especializados para manipulación de texto:

```go
import (
    "strings"
    "strconv"
    "fmt"
)

// Manipulación de strings
texto := strings.ToUpper("hola mundo")
partes := strings.Split("a,b,c", ",")

// Conversión de tipos
numero, err := strconv.Atoi("123")    // string a int
texto := strconv.Itoa(456)           // int a string
decimal, err := strconv.ParseFloat("3.14", 64)
```

## Zero Values (Valores por Defecto)

Cada tipo en Go tiene un valor por defecto cuando se declara sin inicializar:

```go
var entero int        // 0
var flotante float64  // 0.0
var booleano bool     // false
var cadena string     // ""
var puntero *int      // nil
var slice []int       // nil
var mapa map[string]int // nil
var canal chan int    // nil
```

## Conversión de Tipos

Go requiere conversiones explícitas entre tipos diferentes:

```go
var entero int = 42
var flotante float64 = float64(entero)  // Conversión explícita
var texto string = strconv.Itoa(entero) // int a string

// Error: no se permite conversión implícita
// var resultado float64 = entero + 3.14 // ❌ Error de compilación
var resultado float64 = float64(entero) + 3.14 // ✅ Correcto
```

## Verificación de Tipos en Runtime

```go
var interfaz interface{} = "hola"

// Type assertion
if texto, ok := interfaz.(string); ok {
    fmt.Println("Es un string:", texto)
}

// Type switch
switch v := interfaz.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Integer:", v)
default:
    fmt.Println("Tipo desconocido")
}
```