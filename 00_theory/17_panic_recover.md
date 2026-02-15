# Panic y Recover

Go maneja errores mediante valores de retorno, pero proporciona `panic` y `recover` para situaciones excepcionales que no pueden manejarse de forma elegante.

## Panic

`panic` detiene la ejecución normal y comienza a deshacer la pila de llamadas:

```go
func dividir(a, b float64) float64 {
    if b == 0 {
        panic("división por cero")
    }
    return a / b
}

func main() {
    resultado := dividir(10, 0)
    fmt.Println(resultado)    // No se ejecuta
}
```

El mensaje de panic se imprime y el programa termina con código de salida distinto de cero.

## Recover

`recover` captura un panic y permite que el programa continúe. Solo funciona dentro de funciones diferidas:

```go
func protegido() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado de:", r)
        }
    }()
    
    panic("error grave")
    fmt.Println("No se ejecuta")
}

func main() {
    protegido()
    fmt.Println("Continuando...")    // Se ejecuta
}
```

## Patrón de Protección

Envuelva operaciones riesgosas con defer/recover:

```go
func ejecutarSeguro(operacion func()) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Operación fallida: %v", r)
        }
    }()
    
    operacion()
}

func main() {
    ejecutarSeguro(func() {
        panic("algo salió mal")
    })
    fmt.Println("Programa continúa")
}
```

## Panic en Librerías

Las librerías no deben usar panic para errores predecibles. Retorne error en su lugar:

```go
// Correcto para librerías
func ParseEntero(s string) (int, error) {
    // retornar error si s es inválido
}

// Incorrecto para librerías
func ParseEntero(s string) int {
    if s == "" {
        panic("cadena vacía")    // No hacer esto
    }
    // ...
}
```

Use panic solo para condiciones de programa irrecuperables (bugs, invariantes violados).

## Recuperación Parcial

Después de recover, el programa continúa, pero el estado puede estar inconsistente:

```go
func procesarElementos(elementos []int) {
    for _, e := range elementos {
        func() {
            defer func() {
                if r := recover(); r != nil {
                    fmt.Printf("Error en elemento %d: %v\n", e, r)
                }
            }()
            
            procesar(e)    // Si falla, continúa con el siguiente
        }()
    }
}
```

## Panic vs Error

| Situación | Mecanismo |
|-----------|-----------|
| Entrada inválida | `error` |
| Recurso no encontrado | `error` |
| Error de red temporal | `error` |
| Bug de programación | `panic` |
| Invariante violado | `panic` |
| Inicialización fallida | `panic` (main/init) |

## Re-lanzar Panic

Puede recuperar, registrar y volver a lanzar:

```go
defer func() {
    if r := recover(); r != nil {
        log.Printf("Error crítico: %v", r)
        panic(r)    // Re-lanzar para terminar el programa
    }
}()
```

## Resumen

- `panic` detiene la ejecución y deshace la pila de llamadas
- `recover` captura panics solo dentro de funciones `defer`
- Use `error` para condiciones esperadas; reserve `panic` para bugs irrecuperables
- Las librerías no deben hacer panic por errores de entrada
- Después de recover, el estado del programa puede ser inconsistente
- El patrón defer/recover permite aislar fallos sin terminar el programa
