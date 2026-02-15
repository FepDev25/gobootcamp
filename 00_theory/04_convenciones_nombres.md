# Convenciones de Nomenclatura

Go tiene convenciones de nomenclatura estrictas que afectan directamente la visibilidad de identificadores y son parte integral del diseño del lenguaje.

## Reglas de Visibilidad

La primera letra del identificador determina su alcance:

| Convención | Visibilidad | Ejemplo |
|------------|-------------|---------|
| Inicial mayúscula | Público (exportado) | `Contador`, `NombreUsuario` |
| Inicial minúscula | Privado (paquete) | `contador`, `nombreUsuario` |

```go
package util

// Público: accesible desde otros paquetes
func CalcularArea(base, altura float64) float64 {
    return base * altura
}

// Privado: solo dentro de este paquete
func validarEntrada(valor float64) bool {
    return valor > 0
}
```

## Estilo de Nombres

### CamelCase

Go utiliza camelCase para identificadores compuestos:

```go
var nombreUsuario string
var totalVentas float64
func procesarDatos() {}
```

### Mayúsculas para Siglas

Las siglas se escriben completamente en mayúsculas:

```go
var URL string
var HTTPRequest string
func ParseJSON() {}
```

### Nombres Descriptivos

- Variables: cortas en alcance reducido, descriptivas en alcance amplio
- Funciones: verbo + sustantivo que describa la acción

```go
// Bien: variable de bucle corta
for i := 0; i < n; i++ {}

// Bien: variable de paquete descriptiva
var configuracionServidor ServidorConfig

// Bien: función con nombre claro
func CalcularPromedio(valores []float64) float64
```

## Convenciones por Tipo de Identificador

### Variables

```go
var contador int
var nombreCliente string
var listaUsuarios []Usuario
```

### Constantes

Sin convención de mayúsculas forzada como en otros lenguajes:

```go
const maxConexiones = 100
const VersionSistema = "2.0.1"
```

### Interfaces

Los nombres de interfaces suelen terminar en "-er" cuando describen comportamiento:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### Estructuras

Nombre descriptivo del concepto que representan:

```go
type Cliente struct {
    ID       int
    Nombre   string
    Email    string
}
```

## Nombres de Paquetes

- Minúsculas, sin guiones ni guiones bajos
- Nombre del directorio debe coincidir con el paquete
- Cortos y descriptivos

```
src/
├── http/         // package http
├── json/         // package json
└── net/
    └── smtp/     // package smtp
```

## Getters y Setters

Go no requiere getters/setters explícitos. Para propiedades exportadas, acceda directamente. Para propiedades calculadas, omita el prefijo "Get":

```go
type Persona struct {
    nombre string
}

// No: GetNombre()
// Sí: Nombre()
func (p Persona) Nombre() string {
    return p.nombre
}
```

## Errores Comunes a Evitar

- No use guiones bajos en nombres (`nombre_variable`)
- No use mayúsculas para constantes a menos que sean exportadas
- No abrevie innecesariamente (`cfg` en lugar de `config` solo en contextos locales muy claros)

## Resumen

- Inicial mayúscula = exportado, minúscula = privado
- Use camelCase para identificadores compuestos
- Siglas en mayúsculas completas (URL, HTTP)
- Nombres descriptivos proporcionales al alcance
- Paquetes en minúsculas, coincidiendo con directorio
- Interfaces descriptivas con sufijo "-er" cuando aplique
