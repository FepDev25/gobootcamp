# Init en Go

## ¿Qué es la función init?

La función `init` es una función especial en Go que se ejecuta automáticamente **antes** de la función `main`. Se utiliza para inicializar variables, configurar el entorno, registrar drivers, o realizar cualquier configuración necesaria antes de que comience la ejecución principal del programa.

## Características Fundamentales

### 1. Ejecución Automática

```go
package main

import "fmt"

func init() {
    fmt.Println("Esta línea se ejecuta ANTES de main")
}

func main() {
    fmt.Println("Esta línea se ejecuta en main")
}
```

**Salida:**

```bash
Esta línea se ejecuta ANTES de main
Esta línea se ejecuta en main
```

### 2. No Tiene Parámetros ni Valor de Retorno

```go
// ✅ Correcto
func init() {
    // Código de inicialización
}

// ❌ Incorrecto - init no puede tener parámetros
func init(config string) {
    // Error de compilación
}

// ❌ Incorrecto - init no puede retornar valores
func init() error {
    // Error de compilación
    return nil
}
```

### 3. Múltiples Funciones init por Paquete

```go
package main

import "fmt"

func init() {
    fmt.Println("Primera función init")
}

func init() {
    fmt.Println("Segunda función init")
}

func init() {
    fmt.Println("Tercera función init")
}

func main() {
    fmt.Println("Función main")
}
```

**Salida:**

```bash
Primera función init
Segunda función init
Tercera función init
Función main
```

## Orden de Ejecución

### Dentro de un Paquete

1. **Inicialización de variables del paquete**
2. **Funciones `init` en orden de aparición**
3. **Función `main` (solo en paquete main)**

```go
package main

import "fmt"

// 1. Variables del paquete se inicializan primero
var packageVar = initializeVar()

func initializeVar() string {
    fmt.Println("1. Inicializando variable del paquete")
    return "valor inicial"
}

// 2. Funciones init se ejecutan después
func init() {
    fmt.Println("2. Primera función init")
    fmt.Println("   Variable del paquete:", packageVar)
}

func init() {
    fmt.Println("3. Segunda función init")
}

// 3. Main se ejecuta al final
func main() {
    fmt.Println("4. Función main ejecutándose")
}
```

### Entre Múltiples Paquetes

```go
// archivo: config/config.go
package config

import "fmt"

var AppConfig = "configuración cargada"

func init() {
    fmt.Println("init de paquete config")
    AppConfig = "configuración inicializada"
}
```

```go
// archivo: database/db.go
package database

import (
    "fmt"
    "myapp/config"
)

func init() {
    fmt.Println("init de paquete database")
    fmt.Println("Config disponible:", config.AppConfig)
}
```

```go
// archivo: main.go
package main

import (
    "fmt"
    _ "myapp/database" // Importación para efectos secundarios
    "myapp/config"
)

func init() {
    fmt.Println("init de paquete main")
}

func main() {
    fmt.Println("main ejecutándose")
    fmt.Println("Config final:", config.AppConfig)
}
```

**Salida:**

```bash
init de paquete config
init de paquete database
Config disponible: configuración inicializada
init de paquete main
main ejecutándose
Config final: configuración inicializada
```

## Casos de Uso Comunes

### 1. Inicialización de Variables Globales

```go
package main

import (
    "fmt"
    "os"
    "strconv"
)

var (
    debug     bool
    maxRetries int
    apiURL    string
)

func init() {
    // Configurar debug mode
    debug = os.Getenv("DEBUG") == "true"
    
    // Configurar máximo de reintentos
    if retries := os.Getenv("MAX_RETRIES"); retries != "" {
        if parsed, err := strconv.Atoi(retries); err == nil {
            maxRetries = parsed
        }
    } else {
        maxRetries = 3 // Valor por defecto
    }
    
    // Configurar URL de API
    apiURL = os.Getenv("API_URL")
    if apiURL == "" {
        apiURL = "https://api.example.com" // URL por defecto
    }
    
    if debug {
        fmt.Printf("Configuración inicializada:\n")
        fmt.Printf("  Debug: %v\n", debug)
        fmt.Printf("  Max Retries: %d\n", maxRetries)
        fmt.Printf("  API URL: %s\n", apiURL)
    }
}

func main() {
    fmt.Println("Aplicación iniciada con configuración:")
    fmt.Printf("Debug: %v, Retries: %d, URL: %s\n", debug, maxRetries, apiURL)
}
```

### 2. Registro de Drivers de Base de Datos

```go
// archivo: drivers/mysql.go
package drivers

import (
    "database/sql"
    "database/sql/driver"
    "fmt"
)

// Driver personalizado (simplificado)
type MySQLDriver struct{}

func (d MySQLDriver) Open(name string) (driver.Conn, error) {
    // Implementación del driver
    return nil, nil
}

func init() {
    fmt.Println("Registrando driver MySQL")
    sql.Register("mysql-custom", &MySQLDriver{})
}
```

```go
// archivo: main.go
package main

import (
    "database/sql"
    "fmt"
    _ "myapp/drivers" // Importar para ejecutar init
)

func main() {
    // El driver ya está registrado gracias a init
    db, err := sql.Open("mysql-custom", "connection-string")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    defer db.Close()
    
    fmt.Println("Base de datos conectada")
}
```

### 3. Configuración de Logging

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func init() {
    // Configurar logger global
    logFile := os.Getenv("LOG_FILE")
    if logFile != "" {
        file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err != nil {
            log.Printf("Error abriendo archivo de log: %v", err)
        } else {
            log.SetOutput(file)
            fmt.Printf("Logging configurado para escribir a: %s\n", logFile)
        }
    }
    
    // Configurar formato de log
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetPrefix("[MYAPP] ")
}

func main() {
    log.Println("Aplicación iniciada")
    fmt.Println("Aplicación ejecutándose...")
    log.Println("Aplicación terminada")
}
```

### 4. Validación de Prerrequisitos

```go
package main

import (
    "fmt"
    "os"
    "os/exec"
)

func init() {
    // Verificar que las herramientas necesarias estén instaladas
    requiredTools := []string{"git", "docker", "kubectl"}
    
    for _, tool := range requiredTools {
        if _, err := exec.LookPath(tool); err != nil {
            fmt.Fprintf(os.Stderr, "Error: %s no está instalado\n", tool)
            os.Exit(1)
        }
    }
    
    // Verificar variables de entorno críticas
    requiredEnvVars := []string{"KUBECONFIG", "DOCKER_HOST"}
    
    for _, envVar := range requiredEnvVars {
        if os.Getenv(envVar) == "" {
            fmt.Fprintf(os.Stderr, "Error: variable de entorno %s no está configurada\n", envVar)
            os.Exit(1)
        }
    }
    
    fmt.Println("Todos los prerrequisitos están satisfechos")
}

func main() {
    fmt.Println("Aplicación ejecutándose...")
}
```

### 5. Inicialización de Pools de Conexiones

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    connectionPool *ConnectionPool
    once          sync.Once
)

type ConnectionPool struct {
    connections chan *Connection
    maxSize     int
}

type Connection struct {
    ID     int
    Active bool
}

func init() {
    fmt.Println("Inicializando pool de conexiones...")
    
    poolSize := 10
    connectionPool = &ConnectionPool{
        connections: make(chan *Connection, poolSize),
        maxSize:     poolSize,
    }
    
    // Llenar el pool con conexiones
    for i := 0; i < poolSize; i++ {
        conn := &Connection{
            ID:     i + 1,
            Active: true,
        }
        connectionPool.connections <- conn
    }
    
    fmt.Printf("Pool de conexiones inicializado con %d conexiones\n", poolSize)
}

func (cp *ConnectionPool) GetConnection() *Connection {
    select {
    case conn := <-cp.connections:
        return conn
    case <-time.After(5 * time.Second):
        return nil // Timeout
    }
}

func (cp *ConnectionPool) ReturnConnection(conn *Connection) {
    if conn != nil && conn.Active {
        cp.connections <- conn
    }
}

func main() {
    fmt.Println("Usando pool de conexiones...")
    
    conn := connectionPool.GetConnection()
    if conn != nil {
        fmt.Printf("Obtenida conexión %d\n", conn.ID)
        // Usar conexión...
        connectionPool.ReturnConnection(conn)
    }
}
```

## Importaciones para Efectos Secundarios

### Blank Import (_)

```go
package main

import (
    "fmt"
    _ "database/sql/driver" // Solo para ejecutar init
    _ "github.com/lib/pq"   // Driver PostgreSQL
)

func main() {
    // Los drivers ya están registrados
    fmt.Println("Drivers disponibles")
}
```

### Paquete de Plugins

```go
// archivo: plugins/auth.go
package plugins

import "fmt"

func init() {
    fmt.Println("Plugin de autenticación cargado")
    registerAuthPlugin()
}

func registerAuthPlugin() {
    // Registrar plugin en el sistema
}
```

```go
// archivo: main.go
package main

import (
    "fmt"
    _ "myapp/plugins" // Cargar todos los plugins
)

func main() {
    fmt.Println("Aplicación con plugins cargados")
}
```

## Buenas Prácticas

### ✅ Uso Apropiado

1. **Configuración inicial del programa**
2. **Registro de drivers/plugins**
3. **Inicialización de variables globales complejas**
4. **Configuración de logging**
5. **Validación de prerrequisitos críticos**

### ❌ Uso Inapropiado

1. **Lógica compleja de negocio**
2. **Operaciones que pueden fallar sin recovery**
3. **Inicialización costosa que puede evitarse**
4. **Operaciones que requieren entrada del usuario**

### Ejemplo de Uso Incorrecto vs Correcto

**❌ Incorrecto:**

```go
func init() {
    // No hacer esto en init
    fmt.Print("Ingrese su nombre: ")
    var name string
    fmt.Scanln(&name) // ¡Mal! Requiere entrada del usuario
    
    // Operación costosa innecesaria
    time.Sleep(5 * time.Second) // ¡Mal! Demora el inicio
    
    // Lógica de negocio
    processPayments() // ¡Mal! No es configuración
}
```

**✅ Correcto:**

```go
var (
    config *AppConfig
    logger *log.Logger
)

func init() {
    // Configuración basada en variables de entorno
    config = &AppConfig{
        Port:     getEnvOrDefault("PORT", "8080"),
        LogLevel: getEnvOrDefault("LOG_LEVEL", "info"),
    }
    
    // Configurar logger
    logger = log.New(os.Stdout, "[APP] ", log.LstdFlags)
    
    // Validar configuración crítica
    if config.Port == "" {
        log.Fatal("PORT es requerido")
    }
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

## Testing con Init

### Problema con Tests

```go
// main.go
package main

import "fmt"

var globalValue string

func init() {
    globalValue = "valor desde init"
    fmt.Println("init ejecutado")
}

func GetGlobalValue() string {
    return globalValue
}
```

```go
// main_test.go
package main

import "testing"

func TestGetGlobalValue(t *testing.T) {
    // init ya se ejecutó antes de este test
    result := GetGlobalValue()
    if result != "valor desde init" {
        t.Errorf("Esperado 'valor desde init', obtenido '%s'", result)
    }
}
```

### Solución: Inicialización Explícita

```go
// config.go
package main

import "sync"

var (
    config *AppConfig
    once   sync.Once
)

type AppConfig struct {
    DatabaseURL string
    Port        string
}

func InitConfig() *AppConfig {
    once.Do(func() {
        config = &AppConfig{
            DatabaseURL: getEnv("DATABASE_URL", "localhost:5432"),
            Port:        getEnv("PORT", "8080"),
        }
    })
    return config
}

func GetConfig() *AppConfig {
    if config == nil {
        return InitConfig()
    }
    return config
}
```

## Ejemplo Completo: Aplicación Web

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "strconv"
    "syscall"
    "time"
)

var (
    server *http.Server
    config *ServerConfig
)

type ServerConfig struct {
    Port            string
    ReadTimeout     time.Duration
    WriteTimeout    time.Duration
    ShutdownTimeout time.Duration
    LogLevel        string
}

func init() {
    fmt.Println("Inicializando servidor web...")
    
    // Configurar logging
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetPrefix("[WEB] ")
    
    // Cargar configuración
    config = &ServerConfig{
        Port:            getEnvOrDefault("PORT", "8080"),
        ReadTimeout:     parseDurationOrDefault("READ_TIMEOUT", "30s"),
        WriteTimeout:    parseDurationOrDefault("WRITE_TIMEOUT", "30s"),
        ShutdownTimeout: parseDurationOrDefault("SHUTDOWN_TIMEOUT", "10s"),
        LogLevel:        getEnvOrDefault("LOG_LEVEL", "info"),
    }
    
    // Validar configuración
    if _, err := strconv.Atoi(config.Port); err != nil {
        log.Fatalf("Puerto inválido: %s", config.Port)
    }
    
    // Configurar servidor
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandler)
    mux.HandleFunc("/health", healthHandler)
    
    server = &http.Server{
        Addr:         ":" + config.Port,
        Handler:      mux,
        ReadTimeout:  config.ReadTimeout,
        WriteTimeout: config.WriteTimeout,
    }
    
    log.Printf("Servidor configurado en puerto %s", config.Port)
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func parseDurationOrDefault(key, defaultValue string) time.Duration {
    value := getEnvOrDefault(key, defaultValue)
    duration, err := time.ParseDuration(value)
    if err != nil {
        log.Printf("Error parseando %s: %v, usando valor por defecto", key, err)
        duration, _ = time.ParseDuration(defaultValue)
    }
    return duration
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Servidor web funcionando en puerto %s", config.Port)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "OK")
}

func main() {
    // Canal para señales del sistema
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    
    // Iniciar servidor en goroutine
    go func() {
        log.Printf("Iniciando servidor en :%s", config.Port)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Error iniciando servidor: %v", err)
        }
    }()
    
    // Esperar señal de cierre
    <-quit
    log.Println("Cerrando servidor...")
    
    // Shutdown graceful
    ctx, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Error cerrando servidor: %v", err)
    }
    
    log.Println("Servidor cerrado exitosamente")
}
```

## Resumen

- `init` se ejecuta automáticamente antes de `main`
- Usar para configuración inicial y prerrequisitos
- Puede haber múltiples funciones `init` por paquete
- Ejecuta en orden de dependencias entre paquetes
- Evitar lógica compleja o operaciones costosas
- Ideal para configuración, registro de drivers y validaciones críticas
