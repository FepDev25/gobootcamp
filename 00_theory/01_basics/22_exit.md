# Exit en Go

## ¿Qué es os.Exit?

`os.Exit` es una función del paquete `os` que termina inmediatamente el programa con un código de salida específico. Es una forma abrupta de finalizar la ejecución que **no ejecuta funciones `defer`** ni permite limpieza de recursos.

## Sintaxis

```go
func Exit(code int)
```

- `code`: Código de salida que indica el estado de finalización del programa
  - `0`: Éxito
  - `1-255`: Error (por convención, `1` para errores generales)

## Comportamiento Crítico

### No Ejecuta Defer

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    defer fmt.Println("Esta función defer NO se ejecutará")
    defer cleanup() // ¡Esta limpieza NO ocurrirá!
    
    fmt.Println("Programa iniciado")
    os.Exit(1) // Termina inmediatamente
    fmt.Println("Esta línea nunca se ejecutará")
}

func cleanup() {
    fmt.Println("Limpiando recursos...")
}
```

**Salida:**

```bash
Programa iniciado
```

### Comparación con return y panic

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("=== Ejemplo con return ===")
    withReturn()
    
    fmt.Println("\n=== Ejemplo con panic ===")
    withPanic()
    
    fmt.Println("\n=== Ejemplo con os.Exit ===")
    withExit()
}

func withReturn() {
    defer fmt.Println("Defer ejecutado con return")
    fmt.Println("Antes de return")
    return // Las funciones defer SÍ se ejecutan
    fmt.Println("Después de return")
}

func withPanic() {
    defer fmt.Println("Defer ejecutado con panic")
    fmt.Println("Antes de panic")
    panic("error") // Las funciones defer SÍ se ejecutan
}

func withExit() {
    defer fmt.Println("Defer NO ejecutado con exit")
    fmt.Println("Antes de exit")
    os.Exit(1) // Las funciones defer NO se ejecutan
}
```

## Códigos de Salida Estándar

### Convenciones de Unix/Linux

| Código | Significado | Ejemplo de Uso |
| ---------- | ------------- | --------- | ----------- |
| `0` | Éxito | Programa completado correctamente |
| `1` | Error general | Error de aplicación genérico |
| `2` | Uso incorrecto | Argumentos inválidos |
| `126` | Comando no ejecutable | Permisos insuficientes |
| `127` | Comando no encontrado | Comando no existe |
| `128+n` | Señal fatal | Programa terminado por señal n |

### Ejemplos Prácticos

```go
package main

import (
    "fmt"
    "os"
)

const (
    ExitSuccess = 0
    ExitError   = 1
    ExitUsage   = 2
)

func main() {
    args := os.Args[1:]
    
    if len(args) == 0 {
        fmt.Fprintf(os.Stderr, "Uso: %s <comando>\n", os.Args[0])
        os.Exit(ExitUsage)
    }
    
    command := args[0]
    
    switch command {
    case "help":
        showHelp()
        os.Exit(ExitSuccess)
    case "process":
        if err := processData(); err != nil {
            fmt.Fprintf(os.Stderr, "Error procesando datos: %v\n", err)
            os.Exit(ExitError)
        }
        os.Exit(ExitSuccess)
    default:
        fmt.Fprintf(os.Stderr, "Comando desconocido: %s\n", command)
        os.Exit(ExitUsage)
    }
}

func showHelp() {
    fmt.Println("Comandos disponibles:")
    fmt.Println("  help - Muestra esta ayuda")
    fmt.Println("  process - Procesa datos")
}

func processData() error {
    // Simular procesamiento
    return nil
}
```

## Casos de Uso Apropiados

### 1. Herramientas de Línea de Comandos

```go
package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    var (
        input  = flag.String("input", "", "Archivo de entrada")
        output = flag.String("output", "", "Archivo de salida")
        help   = flag.Bool("help", false, "Mostrar ayuda")
    )
    flag.Parse()
    
    if *help {
        flag.Usage()
        os.Exit(0)
    }
    
    if *input == "" {
        fmt.Fprintf(os.Stderr, "Error: archivo de entrada requerido\n")
        flag.Usage()
        os.Exit(2)
    }
    
    if err := processFile(*input, *output); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println("Procesamiento completado exitosamente")
    os.Exit(0)
}

func processFile(input, output string) error {
    // Lógica de procesamiento
    return nil
}
```

### 2. Scripts de Configuración/Instalación

```go
package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    fmt.Println("Iniciando instalación...")
    
    // Verificar prerrequisitos
    if !checkPrerequisites() {
        fmt.Fprintf(os.Stderr, "Prerrequisitos no cumplidos\n")
        os.Exit(1)
    }
    
    // Instalar componentes
    if err := installComponents(); err != nil {
        fmt.Fprintf(os.Stderr, "Error durante instalación: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println("Instalación completada exitosamente")
    os.Exit(0)
}

func checkPrerequisites() bool {
    // Verificar que git esté instalado
    if _, err := exec.LookPath("git"); err != nil {
        fmt.Fprintf(os.Stderr, "Git no está instalado\n")
        return false
    }
    return true
}

func installComponents() error {
    // Lógica de instalación
    return nil
}
```

### 3. Programas de Validación

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Uso: %s <archivo.json>\n", os.Args[0])
        os.Exit(2)
    }
    
    filename := os.Args[1]
    
    if err := validateJSON(filename); err != nil {
        fmt.Fprintf(os.Stderr, "JSON inválido: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Printf("JSON válido: %s\n", filename)
    os.Exit(0)
}

func validateJSON(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    
    var obj interface{}
    return json.Unmarshal(data, &obj)
}
```

## Buenas Prácticas

### ✅ Cuándo Usar os.Exit

1. **Al final de `main()` en herramientas CLI**
2. **Para errores fatales irrecuperables**
3. **En scripts de una sola ejecución**
4. **Cuando se requiere un código de salida específico**

### ❌ Cuándo NO Usar os.Exit

1. **En bibliotecas/paquetes reutilizables**
2. **En servidores web/aplicaciones de larga duración**
3. **Cuando hay recursos que necesitan limpieza**
4. **En funciones que no sean `main()`**

### Ejemplo de Uso Incorrecto vs Correcto

**❌ Incorrecto:**

```go
// En una biblioteca
func ProcessData(data []byte) []byte {
    if len(data) == 0 {
        os.Exit(1) // ¡NO hagas esto en una biblioteca!
    }
    return process(data)
}
```

**✅ Correcto:**

```go
// En una biblioteca
func ProcessData(data []byte) ([]byte, error) {
    if len(data) == 0 {
        return nil, errors.New("datos vacíos")
    }
    return process(data), nil
}

// En main
func main() {
    result, err := ProcessData(data)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
    // usar result...
}
```

## Alternativas a os.Exit

### 1. return en main()

```go
func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}

func run() error {
    // Lógica principal con manejo de errores
    return nil
}
```

### 2. log.Fatal (usa os.Exit internamente)

```go
import "log"

func main() {
    if err := initialize(); err != nil {
        log.Fatal(err) // Equivale a log.Print(err); os.Exit(1)
    }
}
```

### 3. flag.Usage() + os.Exit

```go
import "flag"

func main() {
    flag.Parse()
    if flag.NArg() == 0 {
        flag.Usage()
        os.Exit(2)
    }
}
```

## Interacción con el Sistema

### Captura del Código de Salida en Scripts

```bash
# Bash script
./mi-programa
exit_code=$?

if [ $exit_code -eq 0 ]; then
    echo "Programa ejecutado exitosamente"
elif [ $exit_code -eq 1 ]; then
    echo "Error general"
elif [ $exit_code -eq 2 ]; then
    echo "Uso incorrecto"
fi
```

### Testing de Códigos de Salida

```go
// main.go
package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    exitCode := flag.Int("exit", 0, "Código de salida")
    flag.Parse()
    
    fmt.Printf("Saliendo con código %d\n", *exitCode)
    os.Exit(*exitCode)
}
```

```bash
# Probar diferentes códigos
go run main.go -exit=0; echo "Exit code: $?"
go run main.go -exit=1; echo "Exit code: $?"
go run main.go -exit=2; echo "Exit code: $?"
```

## Ejemplo Completo: Herramienta CLI

```go
package main

import (
    "flag"
    "fmt"
    "io"
    "os"
    "strings"
)

const (
    ExitSuccess = 0
    ExitError   = 1
    ExitUsage   = 2
)

func main() {
    var (
        input    = flag.String("input", "", "Archivo de entrada")
        output   = flag.String("output", "", "Archivo de salida")
        uppercase = flag.Bool("upper", false, "Convertir a mayúsculas")
        help     = flag.Bool("help", false, "Mostrar ayuda")
    )
    
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Uso: %s [opciones]\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "Herramienta de procesamiento de texto\n\n")
        flag.PrintDefaults()
    }
    
    flag.Parse()
    
    if *help {
        flag.Usage()
        os.Exit(ExitSuccess)
    }
    
    if *input == "" {
        fmt.Fprintf(os.Stderr, "Error: archivo de entrada requerido\n\n")
        flag.Usage()
        os.Exit(ExitUsage)
    }
    
    if err := processText(*input, *output, *uppercase); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(ExitError)
    }
    
    fmt.Println("Procesamiento completado exitosamente")
    os.Exit(ExitSuccess)
}

func processText(inputFile, outputFile string, toUpper bool) error {
    // Leer archivo de entrada
    input, err := os.Open(inputFile)
    if err != nil {
        return fmt.Errorf("no se puede abrir archivo de entrada: %w", err)
    }
    defer input.Close()
    
    data, err := io.ReadAll(input)
    if err != nil {
        return fmt.Errorf("error leyendo archivo: %w", err)
    }
    
    // Procesar texto
    text := string(data)
    if toUpper {
        text = strings.ToUpper(text)
    }
    
    // Escribir resultado
    var output io.Writer = os.Stdout
    if outputFile != "" {
        file, err := os.Create(outputFile)
        if err != nil {
            return fmt.Errorf("no se puede crear archivo de salida: %w", err)
        }
        defer file.Close()
        output = file
    }
    
    _, err = fmt.Fprint(output, text)
    return err
}
```

## Resumen

- `os.Exit` termina el programa inmediatamente sin ejecutar `defer`
- Usar códigos de salida estándar (0=éxito, 1=error, 2=uso incorrecto)
- Apropiado para herramientas CLI y scripts
- Evitar en bibliotecas y servidores de larga duración
- Considerar alternativas como `return` o `log.Fatal`
- Siempre documentar los códigos de salida utilizados
