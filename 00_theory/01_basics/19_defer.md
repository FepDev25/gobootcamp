# Defer en Go

`defer` es una palabra clave única en Go que permite posponer la ejecución de una función hasta que la función que la contiene haya terminado. Es una característica poderosa para la gestión de recursos y la limpieza de código.

## ¿Qué es Defer?

`defer` programa la ejecución de una función para que se ejecute **justo antes** de que la función actual termine, sin importar cómo termine (return normal, panic, etc.).

```go
func ejemplo() {
    fmt.Println("Inicio")
    defer fmt.Println("Esto se ejecuta al final")
    fmt.Println("Medio")
    fmt.Println("Fin")
}
// Salida:
// Inicio
// Medio  
// Fin
// Esto se ejecuta al final
```

## Sintaxis Básica

```go
defer función()
defer función(argumentos)
defer objeto.método()
```

## Orden de Ejecución (LIFO)

Las declaraciones `defer` se ejecutan en orden **LIFO (Last In, First Out)** - como una pila:

```go
func ejemploOrden() {
    defer fmt.Println("Primera defer")
    defer fmt.Println("Segunda defer") 
    defer fmt.Println("Tercera defer")
    fmt.Println("Función normal")
}
// Salida:
// Función normal
// Tercera defer
// Segunda defer
// Primera defer
```

### Ejemplo Práctico del Orden

```go
func procesarMultiplesDefers() {
    for i := 1; i <= 3; i++ {
        defer fmt.Printf("Defer %d\n", i)
    }
    fmt.Println("Procesando...")
}
// Salida:
// Procesando...
// Defer 3
// Defer 2
// Defer 1
```

## Casos de Uso Principales

### 1. Cerrar Archivos

```go
func leerArchivo(nombre string) error {
    archivo, err := os.Open(nombre)
    if err != nil {
        return err
    }
    defer archivo.Close() // Se ejecuta sin importar cómo termine la función
    
    // Leer contenido del archivo
    contenido, err := io.ReadAll(archivo)
    if err != nil {
        return err // archivo.Close() se ejecuta automáticamente
    }
    
    fmt.Println("Contenido:", string(contenido))
    return nil // archivo.Close() se ejecuta automáticamente
}
```

### 2. Liberar Recursos

```go
func trabajarConRecurso() error {
    recurso := adquirirRecurso()
    defer liberarRecurso(recurso) // Garantiza liberación
    
    // Trabajar con el recurso
    if err := procesarRecurso(recurso); err != nil {
        return err // recurso se libera automáticamente
    }
    
    return nil // recurso se libera automáticamente
}
```

### 3. Desbloquear Mutexes

```go
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (sc *SafeCounter) Increment() {
    sc.mu.Lock()
    defer sc.mu.Unlock() // Garantiza desbloqueo
    
    sc.count++
    // Cualquier return o panic aquí desbloqueará el mutex
}
```

### 4. Logging y Medición de Tiempo

```go
func funcionLenta() {
    defer func() {
        inicio := time.Now()
        defer func() {
            fmt.Printf("Función tardó: %v\n", time.Since(inicio))
        }()
    }()
    
    // Simular trabajo lento
    time.Sleep(2 * time.Second)
    fmt.Println("Trabajo completado")
}
```

### 5. Recuperación de Panics

```go
func funcionRiesgosa() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Se recuperó de panic: %v\n", r)
        }
    }()
    
    // Código que puede causar panic
    panic("¡Algo salió mal!")
}
```

## Evaluación de Argumentos

Los argumentos de las funciones `defer` se evalúan **inmediatamente**, no cuando se ejecuta la función:

```go
func ejemploEvaluacion() {
    x := 10
    defer fmt.Println("Valor de x:", x) // x se evalúa como 10 ahora
    
    x = 20
    fmt.Println("x cambió a:", x)
}
// Salida:
// x cambió a: 20
// Valor de x: 10
```

### Capturar Variables por Referencia

```go
func capturarPorReferencia() {
    x := 10
    
    // Usando closure para capturar por referencia
    defer func() {
        fmt.Println("Valor final de x:", x) // Se evalúa cuando se ejecuta
    }()
    
    x = 20
    fmt.Println("x cambió a:", x)
}
// Salida:
// x cambió a: 20
// Valor final de x: 20
```

## Defer con Valores de Retorno Nombrados

```go
func ejemploRetornoNombrado() (resultado int) {
    defer func() {
        resultado *= 2 // Modifica el valor de retorno
    }()
    
    resultado = 10
    return resultado // Se devuelve 20, no 10
}
```

### Modificar Error de Retorno

```go
func operacionConError() (err error) {
    defer func() {
        if err != nil {
            err = fmt.Errorf("operación falló: %w", err)
        }
    }()
    
    // Simular error
    return errors.New("error interno")
}
```

## Patrones Avanzados

### 1. Patrón de Adquisición/Liberación

```go
func adquirirRecurso() *Recurso {
    // Adquirir recurso
    return &Recurso{}
}

func (r *Recurso) liberar() {
    // Lógica de liberación
    fmt.Println("Recurso liberado")
}

func usarRecurso() *Recurso {
    r := adquirirRecurso()
    defer r.liberar()
    return r // El recurso se libera aquí
}
```

### 2. Defer Condicional

```go
func operacionCondicional(necesitaLimpieza bool) {
    if necesitaLimpieza {
        defer func() {
            fmt.Println("Realizando limpieza")
        }()
    }
    
    // Lógica de la función
    fmt.Println("Ejecutando operación")
}
```

### 3. Stack de Defers para Múltiples Recursos

```go
func manejarMultiplesRecursos() error {
    r1, err := adquirirRecurso1()
    if err != nil {
        return err
    }
    defer r1.Close()
    
    r2, err := adquirirRecurso2()
    if err != nil {
        return err // r1 se cierra automáticamente
    }
    defer r2.Close()
    
    r3, err := adquirirRecurso3()
    if err != nil {
        return err // r1 y r2 se cierran automáticamente
    }
    defer r3.Close()
    
    // Usar todos los recursos
    return nil // Todos los recursos se cierran en orden inverso
}
```

## Ejemplos Prácticos Completos

### 1. Sistema de Logging

```go
type Logger struct {
    archivo *os.File
}

func NuevoLogger(nombreArchivo string) (*Logger, error) {
    archivo, err := os.OpenFile(nombreArchivo, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        return nil, err
    }
    
    return &Logger{archivo: archivo}, nil
}

func (l *Logger) Log(mensaje string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Fprintf(l.archivo, "[%s] %s\n", timestamp, mensaje)
}

func (l *Logger) Cerrar() error {
    return l.archivo.Close()
}

func usarLogger() error {
    logger, err := NuevoLogger("app.log")
    if err != nil {
        return err
    }
    defer logger.Cerrar() // Garantiza que el archivo se cierre
    
    logger.Log("Aplicación iniciada")
    logger.Log("Procesando datos")
    logger.Log("Aplicación terminada")
    
    return nil
}
```

### 2. Transacciones de Base de Datos

```go
type Transaction struct {
    db *sql.DB
    tx *sql.Tx
}

func (t *Transaction) Commit() error {
    return t.tx.Commit()
}

func (t *Transaction) Rollback() error {
    return t.tx.Rollback()
}

func ejecutarTransaccion(db *sql.DB) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    // Usar defer para garantizar rollback si no hay commit exitoso
    var committed bool
    defer func() {
        if !committed {
            tx.Rollback()
        }
    }()
    
    // Ejecutar operaciones
    if err := operacion1(tx); err != nil {
        return err // Rollback automático
    }
    
    if err := operacion2(tx); err != nil {
        return err // Rollback automático
    }
    
    // Commit exitoso
    if err := tx.Commit(); err != nil {
        return err
    }
    
    committed = true
    return nil
}
```

### 3. Servidor HTTP con Cleanup

```go
func iniciarServidor() error {
    servidor := &http.Server{
        Addr:    ":8080",
        Handler: http.DefaultServeMux,
    }
    
    // Configurar cierre graceful
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
    
    go func() {
        <-stop
        fmt.Println("Cerrando servidor...")
        
        ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
        defer cancel()
        
        if err := servidor.Shutdown(ctx); err != nil {
            fmt.Printf("Error cerrando servidor: %v\n", err)
        }
    }()
    
    fmt.Println("Servidor iniciado en :8080")
    return servidor.ListenAndServe()
}
```

## Consideraciones de Rendimiento

### 1. Defer no es gratis

```go
// Defer tiene un pequeño costo
func conDefer() {
    defer func() {}() // Pequeño overhead
    // operaciones
}

// Sin defer es más rápido para operaciones muy simples
func sinDefer() {
    // operaciones
    // cleanup manual
}
```

### 2. Evitar Defer en Bucles Intensivos

```go
// ❌ Problemático en bucles intensivos
func procesamientoIntenso() {
    for i := 0; i < 1000000; i++ {
        defer fmt.Println(i) // Acumula muchas funciones defer
        // procesamiento
    }
}

// ✅ Mejor: usar defer fuera del bucle cuando sea posible
func procesamientoOptimizado() {
    defer fmt.Println("Procesamiento completado")
    
    for i := 0; i < 1000000; i++ {
        // procesamiento
        // cleanup manual si es necesario
    }
}
```

## Mejores Prácticas

### 1. Colocar defer Inmediatamente Después de Adquisición

```go
// ✅ Buena práctica
func buenaPractica() error {
    archivo, err := os.Open("archivo.txt")
    if err != nil {
        return err
    }
    defer archivo.Close() // Inmediatamente después de abrir
    
    // usar archivo...
    return nil
}
```

### 2. Verificar Errores en Defer

```go
func manejarErroresEnDefer() error {
    archivo, err := os.Create("temporal.txt")
    if err != nil {
        return err
    }
    
    defer func() {
        if closeErr := archivo.Close(); closeErr != nil {
            fmt.Printf("Error cerrando archivo: %v\n", closeErr)
        }
    }()
    
    // usar archivo...
    return nil
}
```

### 3. Usar defer para Invariantes

```go
func mantenerInvariante() {
    mutex.Lock()
    defer mutex.Unlock() // Garantiza que siempre se desbloquee
    
    // Código que modifica estado compartido
}
```

## Limitaciones y Consideraciones

1. **Overhead de rendimiento**: Pequeño pero existente
2. **Evaluación inmediata de argumentos**: Los valores se capturan cuando se ejecuta defer
3. **No es apropiado para bucles intensivos**: Puede acumular muchas funciones
4. **Orden LIFO**: Puede ser confuso en casos complejos

`defer` es una característica distintiva de Go que simplifica significativamente la gestión de recursos y hace que el código sea más robusto y fácil de mantener. Su uso correcto es esencial para escribir código Go idiomático y libre de leaks de recursos.
