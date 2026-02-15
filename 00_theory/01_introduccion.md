# Introducción a Go

Go es un lenguaje de programación desarrollado por Google en 2009. Combina la eficiencia de los lenguajes compilados con la simplicidad de sintaxis de lenguajes interpretados.

## Características Principales

- **Compilado**: Genera binarios nativos sin dependencias externas
- **Tipado estático**: Detección de errores en tiempo de compilación
- **Garbage collector**: Gestión automática de memoria
- **Concurrencia nativa**: Goroutines y channels integrados
- **Sintaxis minimalista**: Legibilidad como prioridad

## Estructura de un Programa

Todo programa Go requiere:

1. **Paquete principal**: `package main` indica el punto de entrada
2. **Función main**: `func main()` donde inicia la ejecución

```go
package main

import "fmt"

func main() {
    fmt.Println("Hola, Go")
}
```

## Compilación y Ejecución

Desde la terminal en el directorio del archivo:

```bash
# Compilar y ejecutar directamente
go run main.go

# Compilar a binario
go build main.go

# Compilar con nombre específico
go build -o mi_programa main.go
```

## Organización de Código

```
mi_proyecto/
├── go.mod          # Definición del módulo
├── main.go         # Punto de entrada
└── interno/        # Paquetes internos
    └── util.go
```

Inicializar un nuevo módulo:

```bash
go mod init nombre_modulo
```

## Convenciones de Formato

Go incluye `gofmt`, una herramienta que estandariza el formato:

```bash
gofmt -w archivo.go   # Formatear archivo
go fmt ./...          # Formatear todo el proyecto
```

No se usan punto y coma al final de línea. Las llaves de apertura van en la misma línea.

## Resumen

- Todo programa inicia en `package main` con función `main()`
- `go run` ejecuta sin generar binario; `go build` compila
- El formato es parte del lenguaje, no una preferencia personal
- Go prioriza la claridad sobre la expresividad compleja
