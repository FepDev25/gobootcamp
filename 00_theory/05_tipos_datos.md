# Tipos de Datos

Go es un lenguaje fuertemente tipado donde cada valor pertenece a un tipo específico. La declaración explícita de tipos permite detectar errores en tiempo de compilación.

## Tipos Numéricos

### Enteros con Signo

| Tipo | Rango | Bits |
|------|-------|------|
| `int8` | -128 a 127 | 8 |
| `int16` | -32,768 a 32,767 | 16 |
| `int32` | -2,147,483,648 a 2,147,483,647 | 32 |
| `int64` | -9e18 a 9e18 | 64 |
| `int` | Dependiente de la arquitectura (32 o 64) | 32/64 |

### Enteros sin Signo

| Tipo | Rango | Bits |
|------|-------|------|
| `uint8` (byte) | 0 a 255 | 8 |
| `uint16` | 0 a 65,535 | 16 |
| `uint32` | 0 a 4,294,967,295 | 32 |
| `uint64` | 0 a 18e18 | 64 |
| `uint` | Dependiente de la arquitectura | 32/64 |

```go
var edad uint8 = 25
var poblacion uint64 = 8000000000
var temperatura int16 = -15
```

### Punto Flotante

| Tipo | Precisión | Bits |
|------|-----------|------|
| `float32` | 6-7 decimales | 32 |
| `float64` | 15-17 decimales | 64 |

```go
var precio float64 = 19.99
var pi float64 = 3.14159265359
```

Use `float64` por defecto a menos que tenga restricciones de memoria específicas.

### Números Complejos

```go
var c complex128 = 3 + 4i
```

## Tipo Booleano

Representa valores de verdad: `true` o `false`.

```go
var activo bool = true
var completo bool = false
```

## Tipo Cadena

Las cadenas en Go son secuencias inmutables de bytes que representan texto Unicode.

```go
var nombre string = "María García"
var mensaje string = `Texto con
saltos de línea`
```

Comillas dobles interpretan secuencias de escape. Backticks crean cadenas literales.

## Alias de Tipos

| Alias | Tipo Base |
|-------|-----------|
| `byte` | `uint8` |
| `rune` | `int32` (punto de código Unicode) |

```go
var caracter byte = 'A'    // Valor ASCII 65
var simbolo rune = 'ñ'     // Unicode U+00F1
```

## Conversión de Tipos

Go no realiza conversiones implícitas. Debe convertir explícitamente:

```go
var entero int = 42
var flotante float64 = float64(entero)

var pequeno int32 = 100
var grande int64 = int64(pequeno)
```

Conversión entre cadenas y números requiere el paquete `strconv`:

```go
import "strconv"

numero, _ := strconv.Atoi("42")
cadena := strconv.Itoa(42)
```

## Tipo de Dato por Defecto

Sin tipo explícito, el compilador asigna:

```go
x := 42          // int
y := 3.14        // float64
z := true        // bool
s := "texto"     // string
```

## Resumen

- Use `int` para enteros generales; especifique tamaño solo cuando sea necesario
- Prefiera `float64` sobre `float32` para mayor precisión
- Las conversiones deben ser explícitas: `tipo(valor)`
- `byte` es `uint8`; `rune` es `int32` para caracteres Unicode
- Las cadenas son inmutables y codificadas en UTF-8
