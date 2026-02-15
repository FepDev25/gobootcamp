# Exit e Init

Go proporciona mecanismos para controlar el inicio y la terminación de programas: la función especial `init` para inicialización y `os.Exit` para terminación inmediata.

## Función Init

Cada paquete puede contener funciones `init` que se ejecutan automáticamente antes de `main`:

```go
package main

import "fmt"

var mensaje string

func init() {
    mensaje = "Inicializando..."
    fmt.Println(mensaje)
}

func main() {
    fmt.Println("Ejecutando main")
}
// Salida:
// Inicializando...
// Ejecutando main
```

### Múltiples Init

Un archivo puede tener múltiples funciones `init`, ejecutándose en orden de declaración:

```go
func init() {
    fmt.Println("Init 1")
}

func init() {
    fmt.Println("Init 2")
}
```

### Orden de Ejecución

1. Inicialización de variables de paquete
2. Funciones `init` del paquete (en orden de archivo)
3. `main` del paquete `main`

Para imports, el orden es recursivo: dependencias primero.

### Uso Común de Init

- Configurar variables de paquete
- Registrar tipos en factories
- Verificar dependencias en tiempo de ejecución
- Inicializar conexiones a servicios externos

```go
var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }
}
```

## os.Exit

Termina el programa inmediatamente con un código de salida:

```go
import "os"

func main() {
    if errorCritico {
        os.Exit(1)    // Termina con código de error
    }
    os.Exit(0)        // Éxito (implícito si no se llama)
}
```

| Código | Significado |
|--------|-------------|
| 0 | Éxito |
| 1 | Error general |
| 2 | Uso incorrecto de comando |
| Otros | Definidos por la aplicación |

### Diferencia con Return

- `return` desde `main`: ejecuta defers, retorna al runtime
- `os.Exit`: termina inmediatamente, defers no se ejecutan

```go
func main() {
    defer fmt.Println("No se ejecuta con Exit")
    os.Exit(0)
}
```

### Exit en Funciones Diferidas

`os.Exit` no se ve afectado por `recover`:

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado")    // No se ejecuta
        }
    }()
    
    panic("error")
    // vs
    os.Exit(1)    // Termina sin recuperación
}
```

## log.Fatal

Combina mensaje de error y terminación:

```go
import "log"

func main() {
    archivo, err := os.Open("config.txt")
    if err != nil {
        log.Fatal(err)    // Imprime error y llama os.Exit(1)
    }
}
```

## Comparación de Mecanismos de Terminación

| Mecanismo | Defers | Recover | Código de Salida |
|-----------|--------|---------|------------------|
| `main` retorna | Sí | N/A | 0 o error |
| `os.Exit` | No | No | Especificado |
| `log.Fatal` | No | No | 1 |
| `panic` | Sí | Sí | 2 (no capturado) |

## Cuándo Usar Cada Uno

- **return**: Terminación normal de main
- **os.Exit**: Código de salida específico, terminación inmediata
- **log.Fatal**: Error irrecuperable con mensaje
- **panic**: Bug o condición inesperada que no puede manejarse localmente

## Resumen

- `init` se ejecuta automáticamente antes de `main` para inicialización de paquetes
- Múltiples `init` en un archivo se ejecutan en orden de declaración
- Las dependencias se inicializan antes que el paquete que las importa
- `os.Exit` termina el programa inmediatamente sin ejecutar defers
- `log.Fatal` combina mensaje de error con `os.Exit(1)`
- Prefiera `return` sobre `os.Exit` en `main` para permitir limpieza con defers
