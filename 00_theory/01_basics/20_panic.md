# Panic en Go

## ¿Qué es Panic?

`panic` es una función integrada en Go que detiene la ejecución normal de un programa. Cuando se llama a `panic`, el programa deja de ejecutar el código actual, ejecuta cualquier función `defer` pendiente en orden inverso (LIFO - Last In, First Out), y luego termina el programa con un mensaje de error.

## Sintaxis

```go
panic(v interface{})
```

- `v` puede ser cualquier valor (string, error, int, etc.)
- Generalmente se usa un string descriptivo o un error

## ¿Cuándo ocurre un Panic?

### Panic Explícito

Cuando el programador llama directamente a `panic()`:

```go
func divide(a, b int) int {
    if b == 0 {
        panic("división por cero no permitida")
    }
    return a / b
}
```

### Panic Implícito

Go puede generar panic automáticamente en situaciones como:

- **Acceso fuera de límites en arrays/slices:**

```go
arr := [3]int{1, 2, 3}
fmt.Println(arr[5]) // panic: runtime error: index out of range
```

- **Desreferencia de puntero nil:**

```go
var p *int
fmt.Println(*p) // panic: runtime error: invalid memory address
```

- **Conversión de tipo inválida:**

```go
var i interface{} = "hello"
num := i.(int) // panic: interface conversion
```

- **Operaciones en canales cerrados:**

```go
ch := make(chan int)
close(ch)
ch <- 1 // panic: send on closed channel
```

## Flujo de Ejecución con Panic

1. Se ejecuta `panic()`
2. La función actual se detiene inmediatamente
3. Se ejecutan todas las funciones `defer` de la función actual (en orden inverso)
4. Se propaga hacia arriba en la pila de llamadas
5. El proceso se repite en cada función hasta llegar a `main()`
6. El programa termina con un stack trace

### Ejemplo Detallado

```go
package main

import "fmt"

func main() {
    fmt.Println("Inicio del programa")
    levelOne()
    fmt.Println("Esta línea nunca se ejecutará")
}

func levelOne() {
    defer fmt.Println("Defer en levelOne")
    fmt.Println("En levelOne")
    levelTwo()
    fmt.Println("Esta línea nunca se ejecutará")
}

func levelTwo() {
    defer fmt.Println("Defer en levelTwo")
    fmt.Println("En levelTwo")
    levelThree()
    fmt.Println("Esta línea nunca se ejecutará")
}

func levelThree() {
    defer fmt.Println("Defer en levelThree")
    fmt.Println("En levelThree")
    panic("¡Algo salió mal!")
    fmt.Println("Esta línea nunca se ejecutará")
}
```

**Salida:**

```bash
Inicio del programa
En levelOne
En levelTwo
En levelThree
Defer en levelThree
Defer en levelTwo
Defer en levelOne
panic: ¡Algo salió mal!

goroutine 1 [running]:
main.levelThree()
    /path/to/file.go:26 +0x95
main.levelTwo()
    /path/to/file.go:19 +0x20
main.levelOne()
    /path/to/file.go:12 +0x20
main.main()
    /path/to/file.go:6 +0x20
```

## Casos de Uso Apropiados

### 1. Errores de Programación

Usar panic para errores que indican bugs en el código:

```go
func getElement(slice []int, index int) int {
    if index < 0 || index >= len(slice) {
        panic(fmt.Sprintf("índice %d fuera de rango para slice de longitud %d", 
              index, len(slice)))
    }
    return slice[index]
}
```

### 2. Configuración Inválida al Inicio

Para errores críticos durante la inicialización:

```go
func init() {
    configFile := os.Getenv("CONFIG_FILE")
    if configFile == "" {
        panic("variable de entorno CONFIG_FILE es requerida")
    }
}
```

### 3. Precondiciones Críticas

Cuando las precondiciones son absolutamente necesarias:

```go
func processPayment(amount float64) {
    if amount <= 0 {
        panic("el monto debe ser positivo")
    }
    // procesar pago...
}
```

## Buenas Prácticas

### ✅ Cuándo Usar Panic

1. **Errores irrecuperables de programación**
2. **Fallas en la inicialización crítica**
3. **Violaciones de invariantes importantes**
4. **En bibliotecas, solo para errores de uso incorrecto**

### ❌ Cuándo NO Usar Panic

1. **Para manejar errores esperados**
2. **En APIs públicas (preferir devolver errores)**
3. **Para validación de entrada de usuario**
4. **En operaciones de red o E/O**

### Ejemplo de Uso Incorrecto vs Correcto

**❌ Incorrecto:**

```go
func openFile(filename string) *os.File {
    file, err := os.Open(filename)
    if err != nil {
        panic(err) // ¡NO hagas esto!
    }
    return file
}
```

**✅ Correcto:**

```go
func openFile(filename string) (*os.File, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("error abriendo archivo %s: %w", filename, err)
    }
    return file, nil
}
```

## Diferencias con Otros Lenguajes

| Concepto | Go (panic) | Java (Exception) | Python (Exception) |
| ---------- | ------------ | ------------------ | ------------------- |
| Propagación | Automática hacia arriba | try/catch puede manejar | try/except puede manejar |
| Recuperación | Solo con `recover()` | Múltiples catch blocks | Múltiples except blocks |
| Uso recomendado | Errores de programación | Control de flujo normal | Control de flujo normal |

## Consideraciones de Rendimiento

- **Panic es costoso:** Detiene la ejecución y desenrolla la pila
- **Evitar panic en bucles críticos**
- **Usar errores normales para casos frecuentes**

## Ejemplo Completo

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Ejemplo de uso apropiado de panic
    defer func() {
        fmt.Println("Limpieza final ejecutada")
    }()

    fmt.Println("Iniciando aplicación...")
    
    // Simular configuración crítica
    initializeApp()
    
    fmt.Println("Aplicación inicializada correctamente")
}

func initializeApp() {
    defer func() {
        fmt.Println("Limpieza de inicialización")
    }()

    // Verificar dependencias críticas
    checkCriticalDependencies()
    
    fmt.Println("Dependencias verificadas")
}

func checkCriticalDependencies() {
    defer func() {
        fmt.Println("Verificación de dependencias completada")
    }()

    // Simular verificación de base de datos crítica
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        panic("DATABASE_URL es requerida para el funcionamiento de la aplicación")
    }
    
    fmt.Println("Base de datos configurada:", dbURL)
}
```

## Resumen

- `panic` es para errores graves e irrecuperables
- Detiene la ejecución inmediata y ejecuta defers
- Usar con moderación y solo para errores de programación
- Preferir errores normales para casos recuperables
- Siempre considerar si `recover()` es necesario para limpieza
