# Constantes en Go

Las constantes en Go son valores inmutables que se evalúan en tiempo de compilación. Una vez definidas, no pueden ser modificadas durante la ejecución del programa.

## Declaración de Constantes

### Sintaxis Básica

```go
const PI = 3.14159
const GRAVITY = 9.81
const COMPANY_NAME = "TechCorp"
```

### Constantes Tipadas vs No Tipadas

#### Constantes No Tipadas (Untyped)
```go
const PI = 3.14        // No tipada, puede ser float32 o float64
const COUNT = 100      // No tipada, puede ser cualquier tipo entero
const MESSAGE = "Hello" // No tipada, será string
```

#### Constantes Tipadas (Typed)
```go
const E float64 = 2.71828    // Específicamente float64
const MAX_SIZE int = 1024    // Específicamente int
const DEBUG bool = true      // Específicamente bool
```

## Declaración en Grupo

Puedes declarar múltiples constantes usando la sintaxis de grupo:

```go
const (
    MONDAY    = 1
    TUESDAY   = 2
    WEDNESDAY = 3
    THURSDAY  = 4
    FRIDAY    = 5
)
```

## Iota - Generador Automático

`iota` es un identificador predeclarado que se incrementa automáticamente:

```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)
```

### Ejemplos Avanzados con iota

```go
// Potencias de 2
const (
    KB = 1 << (10 * iota)  // 1024
    MB                     // 1048576
    GB                     // 1073741824
)

// Estados de conexión
const (
    DISCONNECTED = iota
    CONNECTING
    CONNECTED
    ERROR
)

// Saltando valores
const (
    A = iota    // 0
    B           // 1
    _           // 2 (valor ignorado)
    D           // 3
)
```

## Ámbito de Constantes

### Constantes Globales (Exportadas)
```go
const PublicConstant = "Visible desde otros paquetes"
const API_VERSION = "v1.2.3"
```

### Constantes Privadas (No Exportadas)
```go
const privateConstant = "Solo visible en este paquete"
const internalConfig = "configuración interna"
```

## Tipos de Constantes

### Constantes Numéricas
```go
const (
    // Enteros
    MAX_USERS = 1000
    DEFAULT_PORT = 8080
    
    // Flotantes
    PI = 3.14159265
    E = 2.71828
    
    // Complejos
    COMPLEX_CONSTANT = 3 + 4i
)
```

### Constantes de Cadena
```go
const (
    WELCOME_MESSAGE = "¡Bienvenido!"
    ERROR_MESSAGE = "Ha ocurrido un error"
    API_ENDPOINT = "https://api.example.com"
)
```

### Constantes Booleanas
```go
const (
    DEBUG_MODE = true
    PRODUCTION = false
    ENABLE_CACHE = true
)
```

## Ventajas de las Constantes

1. **Inmutabilidad**: No pueden ser modificadas accidentalmente
2. **Rendimiento**: Se evalúan en tiempo de compilación
3. **Legibilidad**: Hacen el código más expresivo
4. **Mantenibilidad**: Centralizan valores importantes
5. **Seguridad de tipos**: Previenen errores de asignación

## Limitaciones

1. **Solo valores básicos**: No puedes usar constantes para slices, maps o structs
2. **Evaluación en compilación**: Deben ser calculables en tiempo de compilación
3. **Sin funciones**: No puedes usar llamadas a funciones en constantes

```go
// ✓ Válido
const VALID = 10 * 20

// ✗ Inválido
const INVALID = len("hello")        // Error: no se puede evaluar en compilación
const SLICE = []int{1, 2, 3}       // Error: slice no es un tipo básico
```

## Ejemplo Práctico Completo

```go
package main

import "fmt"

// Constantes globales
const (
    // Configuración de la aplicación
    APP_NAME = "MiAplicacion"
    VERSION = "1.0.0"
    
    // Estados usando iota
    INACTIVE = iota
    ACTIVE
    PENDING
    SUSPENDED
)

// Constantes matemáticas tipadas
const (
    PI float64 = 3.14159265359
    E float64 = 2.71828182846
)

// Constantes de configuración
const (
    MAX_CONNECTIONS = 100
    TIMEOUT_SECONDS = 30
    RETRY_ATTEMPTS = 3
)

func main() {
    fmt.Println("Aplicación:", APP_NAME)
    fmt.Println("Versión:", VERSION)
    fmt.Println("Pi:", PI)
    fmt.Println("E:", E)
    
    // Usando constantes en cálculos
    radius := 5.0
    area := PI * radius * radius
    fmt.Printf("Área del círculo: %.2f\n", area)
    
    // Estados
    fmt.Println("Estado Inactivo:", INACTIVE)    // 0
    fmt.Println("Estado Activo:", ACTIVE)        // 1
    fmt.Println("Estado Pendiente:", PENDING)    // 2
    fmt.Println("Estado Suspendido:", SUSPENDED) // 3
    
    // Configuración
    fmt.Printf("Máximo %d conexiones con timeout de %d segundos\n", 
               MAX_CONNECTIONS, TIMEOUT_SECONDS)
}
```

## Convenciones de Nomenclatura

1. **UPPER_CASE**: Para constantes globales importantes
2. **PascalCase**: Para constantes exportadas
3. **camelCase**: Para constantes privadas
4. **Nombres descriptivos**: Evita abreviaciones confusas

```go
// ✓ Buenas prácticas
const MAX_RETRY_ATTEMPTS = 3
const DefaultTimeout = 30
const apiBaseURL = "https://api.example.com"

// ✗ Evitar
const MRA = 3              // No descriptivo
const T = 30               // Muy corto
const url = "..."          // Podría confundirse con variable
```

## Constantes vs Variables

| Aspecto | Constantes | Variables |
|---------|------------|-----------|
| **Mutabilidad** | Inmutable | Mutable |
| **Evaluación** | Tiempo de compilación | Tiempo de ejecución |
| **Rendimiento** | Más eficiente | Menos eficiente |
| **Memoria** | No usa memoria en runtime | Usa memoria |
| **Tipos permitidos** | Solo tipos básicos | Todos los tipos |

Las constantes son una herramienta fundamental en Go para crear código más seguro, eficiente y mantenible. Úsalas siempre que tengas valores que no cambiarán durante la ejecución del programa.
