# Paquetes e Importaciones

Go organiza el código en paquetes. Un paquete es una colección de archivos fuente en el mismo directorio que comparten un propósito común.

## Declaración de Paquetes

Al inicio de cada archivo:

```go
package main        // Ejecutable
package utils       // Biblioteca reutilizable
```

- `package main`: Define un programa ejecutable
- Cualquier otro nombre: Define una biblioteca importable

## Importación de Paquetes

### Importación Simple

```go
import "fmt"
import "os"
```

### Importación Múltiple (forma recomendada)

```go
import (
    "fmt"
    "os"
    "time"
)
```

### Importación con Alias

Útil para resolver conflictos de nombres o abreviar nombres largos:

```go
import (
    f "fmt"
    "github.com/usuario/paquete-largo/libreria"
)

func main() {
    f.Println("Usando alias")
}
```

### Importación Anónima

Ejecuta el `init()` del paquete sin exponer sus funciones:

```go
import _ "database/sql/driver"
```

## Paquetes de la Librería Estándar

| Paquete | Propósito |
|---------|-----------|
| `fmt` | Formateo de entrada/salida |
| `os` | Interacción con el sistema operativo |
| `io` | Operaciones de entrada/salida |
| `strings` | Manipulación de cadenas |
| `time` | Manejo de fechas y tiempos |
| `encoding/json` | Codificación/decodificación JSON |
| `net/http` | Servidor y cliente HTTP |

## Creación de Paquetes Personalizados

Estructura de directorios:

```
proyecto/
├── main.go
└── matematica/
    └── operaciones.go
```

`matematica/operaciones.go`:

```go
package matematica

func Sumar(a, b int) int {
    return a + b
}
```

`main.go`:

```go
package main

import (
    "fmt"
    "proyecto/matematica"
)

func main() {
    resultado := matematica.Sumar(3, 4)
    fmt.Println(resultado)
}
```

## Visibilidad de Identificadores

- **Mayúscula inicial**: Público (exportado)
- **Minúscula inicial**: Privado al paquete

```go
package util

func Publica() {}    // Accesible desde otros paquetes
func privada() {}    // Solo dentro de este paquete
```

## Resumen

- Cada archivo pertenece a un paquete declarado en la primera línea
- Use paréntesis para importar múltiples paquetes
- Los alias resuelven conflictos y simplifican nombres largos
- La capitalización determina la visibilidad: mayúscula es pública
- La importación anónima ejecuta código de inicialización del paquete
