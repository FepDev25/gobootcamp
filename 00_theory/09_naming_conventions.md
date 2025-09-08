# Convenciones de Nomenclatura en Go

Go tiene convenciones específicas para nombrar identificadores que no solo mejoran la legibilidad del código, sino que también afectan la visibilidad y el comportamiento del programa.

## Reglas Generales

### 1. Caracteres Válidos
- Debe comenzar con una letra o guión bajo `_`
- Puede contener letras, números y guiones bajos
- **Case-sensitive**: `variable` y `Variable` son diferentes

### 2. Palabras Reservadas
No puedes usar palabras clave de Go como nombres:
```go
break    default     func    interface  select
case     defer       go      map        struct
chan     else        goto    package    switch
const    fallthrough if      range      type
continue for         import  return     var
```

## Convenciones por Tipo

### Variables
- **Locales**: camelCase comenzando con minúscula
- **Globales**: PascalCase comenzando con mayúscula (público) o camelCase (privado)

```go
// Variables locales
var userName string
var totalCount int
var isActive bool

// Variables globales públicas (exportadas)
var GlobalConfig string

// Variables globales privadas (no exportadas)
var internalCounter int
```

### Constantes
- Tradicionalmente en UPPER_CASE o PascalCase
- Para constantes exportadas: PascalCase
- Para constantes internas: camelCase o UPPER_CASE

```go
const PI = 3.14159
const GRAVITY = 9.81
const MaxConnections = 100

// Constantes privadas
const defaultTimeout = 30
const MAX_RETRIES = 3
```

### Funciones
- **Públicas**: PascalCase (primera letra mayúscula)
- **Privadas**: camelCase (primera letra minúscula)

```go
// Función pública (exportada)
func CalculateArea(radius float64) float64 {
    return PI * radius * radius
}

// Función privada (no exportada)
func validateInput(input string) bool {
    return len(input) > 0
}
```

### Paquetes
- Nombres cortos, descriptivos y en minúsculas
- Sin guiones bajos ni camelCase
- Preferir una sola palabra

```go
package http    // ✓ Bueno
package json    // ✓ Bueno
package fmt     // ✓ Bueno

package httpServer  // ✗ Evitar
package HTTP_Server // ✗ Evitar
```

## Visibilidad (Público vs Privado)

En Go, la visibilidad se determina por la primera letra del identificador:

### Públicos (Exportados)
- Primera letra en **mayúscula**
- Accesibles desde otros paquetes

```go
var PublicVariable int
func PublicFunction() {}
type PublicStruct struct {
    PublicField string
}
```

### Privados (No Exportados)
- Primera letra en **minúscula**
- Solo accesibles dentro del mismo paquete

```go
var privateVariable int
func privateFunction() {}
type privateStruct struct {
    privateField string
}
```

## Convenciones Específicas

### Nombres de Variables Cortas
Para variables de vida corta, se aceptan nombres de una letra:

```go
// En loops
for i := 0; i < len(slice); i++ {}
for k, v := range map {}

// Variables temporales
if err := doSomething(); err != nil {}
```

### Acrónimos
Los acrónimos deben mantener su formato:

```go
// ✓ Correcto
var HTTPClient *http.Client
var URLPath string
var XMLParser Parser

// ✗ Incorrecto
var HttpClient *http.Client
var UrlPath string
var XmlParser Parser
```

### Interfaces
Comúnmente terminan en "-er":

```go
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type Stringer interface {
    String() string
}
```

## Ejemplos Prácticos

```go
package main

import "fmt"

// Constantes
const (
    MaxRetries = 5
    apiTimeout = 30 // segundos
)

// Variable global pública
var GlobalCounter int

// Variable global privada
var internalState string

// Función pública
func ProcessData(input []byte) error {
    // Variables locales
    var result string
    isValid := validateData(input)
    
    if !isValid {
        return fmt.Errorf("invalid data")
    }
    
    // Procesamiento...
    return nil
}

// Función privada
func validateData(data []byte) bool {
    return len(data) > 0
}

// Estructura pública
type UserAccount struct {
    Name     string // Campo público
    Email    string // Campo público
    password string // Campo privado
}

// Método público
func (u *UserAccount) GetInfo() string {
    return fmt.Sprintf("%s (%s)", u.Name, u.Email)
}

// Método privado
func (u *UserAccount) validatePassword() bool {
    return len(u.password) >= 8
}
```

## Herramientas

### go fmt
Automáticamente formatea el código siguiendo las convenciones de Go:
```bash
go fmt ./...
```

### golint
Herramienta que verifica el cumplimiento de las convenciones:
```bash
golint ./...
```

## Resumen de Mejores Prácticas

1. **Consistencia**: Mantén un estilo consistente en todo el proyecto
2. **Claridad**: Los nombres deben ser descriptivos y claros
3. **Brevedad**: No sacrifiques claridad, pero evita nombres excesivamente largos
4. **Contexto**: Los nombres cortos están bien en contextos pequeños
5. **Convenciones**: Sigue las convenciones establecidas de Go
6. **Herramientas**: Usa `go fmt` y `golint` regularmente

Seguir estas convenciones hace que tu código sea más legible, mantenible y consistente con el ecosistema Go.
