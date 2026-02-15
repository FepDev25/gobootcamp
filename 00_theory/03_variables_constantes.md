# Variables y Constantes

Go es un lenguaje tipado estático donde cada variable tiene un tipo definido en tiempo de compilación.

## Declaración de Variables

### Forma Explícita

```go
var nombre string = "Ana"
var edad int = 30
```

### Inferencia de Tipo

El compilador deduce el tipo del valor asignado:

```go
var nombre = "Ana"    // string
var edad = 30         // int
```

### Declaración Corta (Solo Dentro de Funciones)

Operador `:=` para declarar e inicializar:

```go
func main() {
    nombre := "Ana"       // var nombre string = "Ana"
    edad := 30            // var edad int = 30
    salario := 2500.50    // var salario float64
}
```

### Declaración Múltiple

```go
var x, y, z int = 1, 2, 3

// Con inferencia
var a, b, c = 1, "dos", 3.0

// Declaración corta múltiple
nombre, edad := "Ana", 30
```

## Valor Cero

Las variables declaradas sin inicialización obtienen el valor cero de su tipo:

| Tipo | Valor Cero |
|------|------------|
| `int`, `float` | `0`, `0.0` |
| `string` | `""` (cadena vacía) |
| `bool` | `false` |
| `puntero`, `slice`, `map`, `channel`, `func` | `nil` |

```go
var contador int      // 0
var mensaje string    // ""
var activo bool       // false
```

## Constantes

Las constantes son valores inmutables evaluados en tiempo de compilación.

```go
const Pi = 3.14159
const Saludo = "Hola"
```

### Declaración Múltiple

```go
const (
    Lunes = 1
    Martes = 2
    Miercoles = 3
)
```

### Iota

Generador de constantes incrementales dentro de un bloque `const`:

```go
const (
    Lunes = iota      // 0
    Martes            // 1
    Miercoles         // 2
    Jueves            // 3
)
```

Patrón común para máscaras de bits:

```go
const (
    Lectura = 1 << iota   // 1 (0001)
    Escritura             // 2 (0010)
    Ejecucion             // 4 (0100)
)
```

## Tipado Explícito vs. Inferencia

- Use `var` para variables de paquete o cuando necesite claridad explícita
- Use `:=` para variables locales donde el tipo sea obvio del contexto
- Evite declarar variables que no se utilizarán; el compilador lo detectará como error

## Resumen

- `var nombre tipo = valor` para declaración explícita
- `nombre := valor` para declaración corta (solo funciones)
- Las variables tienen valor cero por defecto; no hay valores indefinidos
- `const` para valores inmutables conocidos en compilación
- `iota` genera secuencias de constantes de forma concisa
