# Recover en Go

## ¿Qué es Recover?

`recover` es una función integrada en Go que permite recuperar el control de una goroutine en estado de pánico. Es la única forma de "capturar" un panic y evitar que el programa termine abruptamente. **Solo funciona dentro de funciones diferidas (`defer`)**.

## Sintaxis

```go
func recover() interface{}
```

- Retorna el valor pasado a `panic()` si hay un panic activo
- Retorna `nil` si no hay panic o si se llama fuera de una función defer
- Detiene la propagación del panic

## Funcionamiento Básico

### Ejemplo Simple

```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado de panic:", r)
        }
    }()
    
    fmt.Println("Antes del panic")
    panic("¡Algo salió mal!")
    fmt.Println("Esta línea no se ejecutará")
}
```

**Salida:**

```bash
Antes del panic
Recuperado de panic: ¡Algo salió mal!
```

## Reglas Importantes de Recover

### 1. Solo Funciona en Defer

```go
// ❌ INCORRECTO - recover fuera de defer
func badExample() {
    if r := recover(); r != nil { // ¡Esto NO funciona!
        fmt.Println("No funcionará")
    }
    panic("error")
}

// ✅ CORRECTO - recover dentro de defer
func goodExample() {
    defer func() {
        if r := recover(); r != nil { // ¡Esto SÍ funciona!
            fmt.Println("Funciona correctamente")
        }
    }()
    panic("error")
}
```

### 2. Debe Estar en la Misma Goroutine

```go
func main() {
    // ❌ INCORRECTO - recover en diferente goroutine
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("No capturará el panic de la otra goroutine")
        }
    }()
    
    go func() {
        panic("panic en otra goroutine") // No será capturado
    }()
    
    time.Sleep(time.Second)
}
```

### 3. Solo Captura el Panic Más Reciente

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Capturado:", r) // Solo captura "segundo panic"
        }
    }()
    
    defer func() {
        panic("segundo panic")
    }()
    
    panic("primer panic")
}
```

## Patrones de Uso Comunes

### 1. Función de Protección Simple

```go
func safeFunction() (result string, err error) {
    defer func() {
        if r := recover(); r != nil {
            result = ""
            err = fmt.Errorf("error inesperado: %v", r)
        }
    }()
    
    // Código que puede hacer panic
    riskyOperation()
    return "éxito", nil
}

func riskyOperation() {
    panic("operación falló")
}
```

### 2. Wrapper para Funciones de Terceros

```go
func safeThirdPartyCall(fn func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("función de terceros falló: %v", r)
        }
    }()
    
    fn() // Función que podría hacer panic
    return nil
}

// Uso
err := safeThirdPartyCall(func() {
    // Llamada a biblioteca externa que puede hacer panic
    someLibrary.RiskyFunction()
})
if err != nil {
    log.Printf("Error controlado: %v", err)
}
```

### 3. Servidor Web Robusto

```go
func httpHandler(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Panic en handler: %v", r)
            http.Error(w, "Error interno del servidor", 500)
        }
    }()
    
    // Lógica del handler que podría hacer panic
    processRequest(w, r)
}
```

### 4. Worker Pool con Protección

```go
func worker(jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        func() {
            defer func() {
                if r := recover(); r != nil {
                    results <- Result{
                        ID:    job.ID,
                        Error: fmt.Errorf("worker panic: %v", r),
                    }
                }
            }()
            
            // Procesar trabajo
            result := processJob(job)
            results <- result
        }()
    }
}
```

## Recover con Stack Trace

### Capturar Stack Trace Completo

```go
import (
    "fmt"
    "runtime/debug"
)

func functionWithStackTrace() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Panic: %v\n", r)
            fmt.Printf("Stack trace:\n%s", debug.Stack())
        }
    }()
    
    level1()
}

func level1() {
    level2()
}

func level2() {
    panic("error en level2")
}
```

## Buenas Prácticas

### ✅ Cuándo Usar Recover

1. **Proteger servidores web/APIs**
2. **Wrappear bibliotecas de terceros**
3. **En workers/goroutines de larga duración**
4. **Para logging y monitoreo de errores**
5. **En puntos de entrada críticos**

### ❌ Cuándo NO Usar Recover

1. **Para control de flujo normal**
2. **Como reemplazo de manejo de errores**
3. **Para "silenciar" bugs**
4. **En cada función (overhead innecesario)**

### Ejemplo de Uso Incorrecto vs Correcto

**❌ Incorrecto:**

```go
func divide(a, b int) int {
    defer func() {
        recover() // ¡Silenciar errores es malo!
    }()
    
    return a / b // Podría hacer panic con b=0
}
```

**✅ Correcto:**

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("división por cero")
    }
    return a / b, nil
}
```

## Patrón de Recover Completo

```go
package main

import (
    "fmt"
    "log"
    "runtime/debug"
)

// Estructura para manejar diferentes tipos de panic
type PanicHandler struct {
    logger *log.Logger
}

func NewPanicHandler(logger *log.Logger) *PanicHandler {
    return &PanicHandler{logger: logger}
}

func (p *PanicHandler) Recover() {
    if r := recover(); r != nil {
        // Log detallado del panic
        p.logger.Printf("PANIC DETECTADO: %v", r)
        p.logger.Printf("Stack trace:\n%s", debug.Stack())
        
        // Clasificar tipo de panic
        switch v := r.(type) {
        case string:
            p.logger.Printf("Panic tipo string: %s", v)
        case error:
            p.logger.Printf("Panic tipo error: %v", v)
        default:
            p.logger.Printf("Panic tipo desconocido: %T", v)
        }
    }
}

// Función que usa el handler
func riskyOperation() (result string, err error) {
    handler := NewPanicHandler(log.Default())
    
    defer func() {
        handler.Recover()
        if r := recover(); r != nil {
            result = ""
            err = fmt.Errorf("operación falló: %v", r)
        }
    }()
    
    // Simular diferentes tipos de panic
    scenarios := []func(){
        func() { panic("string panic") },
        func() { panic(fmt.Errorf("error panic")) },
        func() { panic(42) },
    }
    
    scenarios[0]() // Ejecutar primer escenario
    
    return "éxito", nil
}

func main() {
    result, err := riskyOperation()
    if err != nil {
        fmt.Printf("Error controlado: %v\n", err)
    } else {
        fmt.Printf("Resultado: %s\n", result)
    }
}
```

## Recover vs Error Handling

| Aspecto | Recover | Error Handling |
| --------- | --------- | ---------------- |
| **Uso** | Casos excepcionales | Flujo normal |
| **Performance** | Costoso | Eficiente |
| **Claridad** | Puede ocultar problemas | Explícito y claro |
| **Testeo** | Difícil de testear | Fácil de testear |
| **Propagación** | Detiene panic | Propaga naturalmente |

## Ejemplo Completo: Servidor HTTP Robusto

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "runtime/debug"
    "time"
)

// Middleware para recover
func recoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                // Log del error
                log.Printf("Panic en %s %s: %v", r.Method, r.URL.Path, err)
                log.Printf("Stack trace:\n%s", debug.Stack())
                
                // Respuesta al cliente
                http.Error(w, "Error interno del servidor", 
                          http.StatusInternalServerError)
            }
        }()
        
        next(w, r)
    }
}

// Handler que puede hacer panic
func riskyHandler(w http.ResponseWriter, r *http.Request) {
    time.Sleep(100 * time.Millisecond) // Simular trabajo
    
    if r.URL.Query().Get("panic") == "true" {
        panic("error simulado")
    }
    
    fmt.Fprintf(w, "Solicitud procesada correctamente")
}

func main() {
    http.HandleFunc("/api", recoverMiddleware(riskyHandler))
    
    fmt.Println("Servidor iniciado en :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Consideraciones de Performance

- **Recover tiene overhead mínimo cuando no hay panic**
- **Panic + recover es más costoso que return error**
- **Usar solo en casos excepcionales**
- **No usar para control de flujo normal**

## Resumen

- `recover` solo funciona dentro de funciones `defer`
- Permite capturar y manejar panics graciosamente
- Usar para proteger puntos críticos del sistema
- No usar como reemplazo del manejo normal de errores
- Siempre logear panics recuperados para debugging
- Combinar con stack traces para mejor debugging
