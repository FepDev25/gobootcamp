# Panic y Recover en Go

`panic` y `recover` son mecanismos de manejo de errores excepcionales en Go. Mientras que el manejo de errores típico utiliza valores de error, `panic` y `recover` están diseñados para situaciones verdaderamente excepcionales.

## ¿Qué es Panic?

`panic` es una función built-in que detiene el flujo normal de ejecución y comienza a "desenrollar" la pila de llamadas, ejecutando todas las funciones `defer` en el camino.

```go
func ejemploPanic() {
    fmt.Println("Antes del panic")
    panic("¡Algo salió terriblemente mal!")
    fmt.Println("Esto nunca se ejecutará")
}
```

### Cuándo Ocurre Panic Automáticamente

Go puede entrar en panic automáticamente en ciertas situaciones:

```go
// Acceso fuera de límites en array/slice
slice := []int{1, 2, 3}
fmt.Println(slice[10]) // panic: runtime error: index out of range

// División por cero en enteros
x := 5
y := 0
fmt.Println(x / y) // panic: runtime error: integer divide by zero

// Desreferencia de puntero nil
var ptr *int
fmt.Println(*ptr) // panic: runtime error: invalid memory address

// Type assertion incorrecta
var i interface{} = "string"
num := i.(int) // panic: interface conversion

// Envío a channel cerrado
ch := make(chan int)
close(ch)
ch <- 1 // panic: send on closed channel
```

## Crear Panic Manualmente

```go
func validarEdad(edad int) {
    if edad < 0 {
        panic("La edad no puede ser negativa")
    }
    if edad > 150 {
        panic("Edad imposible")
    }
    fmt.Printf("Edad válida: %d\n", edad)
}

func main() {
    validarEdad(25)  // OK
    validarEdad(-5)  // panic!
}
```

## ¿Qué es Recover?

`recover` es una función built-in que puede "capturar" un panic y recuperar el control normal del programa. **Solo funciona dentro de funciones `defer`**.

```go
func funcionRiesgosa() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recuperado de panic: %v\n", r)
        }
    }()
    
    fmt.Println("Antes del panic")
    panic("¡Error crítico!")
    fmt.Println("Nunca se ejecuta")
}

func main() {
    funcionRiesgosa()
    fmt.Println("El programa continúa normalmente")
}
// Salida:
// Antes del panic
// Recuperado de panic: ¡Error crítico!
// El programa continúa normalmente
```

## Patrón Básico de Recover

```go
func manejarPanic() {
    defer func() {
        if r := recover(); r != nil {
            // Manejar el panic
            fmt.Printf("Error recuperado: %v\n", r)
            
            // Opcionalmente, re-lanzar el panic
            // panic(r)
        }
    }()
    
    // Código que puede hacer panic
    realizarOperacionRiesgosa()
}
```

## Ejemplos Prácticos

### 1. Validación con Panic y Recover

```go
type Usuario struct {
    Nombre string
    Email  string
    Edad   int
}

func (u *Usuario) Validar() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("error de validación: %v", r)
        }
    }()
    
    if u.Nombre == "" {
        panic("nombre es requerido")
    }
    
    if !strings.Contains(u.Email, "@") {
        panic("email inválido")
    }
    
    if u.Edad < 0 || u.Edad > 120 {
        panic("edad debe estar entre 0 y 120")
    }
    
    return nil
}

func main() {
    usuario := &Usuario{
        Nombre: "",
        Email:  "invalid-email",
        Edad:   -5,
    }
    
    if err := usuario.Validar(); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
```

### 2. Parser Recursivo

```go
type Parser struct {
    tokens []string
    pos    int
}

func (p *Parser) parse() (result string, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("error de parsing: %v", r)
        }
    }()
    
    return p.parseExpression(), nil
}

func (p *Parser) parseExpression() string {
    if p.pos >= len(p.tokens) {
        panic("fin inesperado de entrada")
    }
    
    token := p.tokens[p.pos]
    p.pos++
    
    if token == "(" {
        result := p.parseExpression()
        p.expect(")")
        return result
    }
    
    return token
}

func (p *Parser) expect(expected string) {
    if p.pos >= len(p.tokens) {
        panic(fmt.Sprintf("esperaba '%s' pero llegó al final", expected))
    }
    
    if p.tokens[p.pos] != expected {
        panic(fmt.Sprintf("esperaba '%s' pero encontré '%s'", expected, p.tokens[p.pos]))
    }
    
    p.pos++
}
```

### 3. Worker Pool con Recuperación

```go
func worker(id int, jobs <-chan func(), results chan<- string) {
    for job := range jobs {
        result := func() (result string) {
            defer func() {
                if r := recover(); r != nil {
                    result = fmt.Sprintf("Worker %d: panic recuperado: %v", id, r)
                }
            }()
            
            job() // Ejecutar trabajo que puede hacer panic
            return fmt.Sprintf("Worker %d: trabajo completado", id)
        }()
        
        results <- result
    }
}

func main() {
    jobs := make(chan func(), 100)
    results := make(chan string, 100)
    
    // Iniciar workers
    for i := 1; i <= 3; i++ {
        go worker(i, jobs, results)
    }
    
    // Enviar trabajos (algunos que causan panic)
    for i := 0; i < 5; i++ {
        num := i
        jobs <- func() {
            if num%2 == 0 {
                panic(fmt.Sprintf("trabajo %d falló", num))
            }
            fmt.Printf("Trabajo %d completado\n", num)
        }
    }
    close(jobs)
    
    // Recoger resultados
    for i := 0; i < 5; i++ {
        fmt.Println(<-results)
    }
}
```

### 4. Servidor HTTP con Recuperación

```go
func recuperarMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Panic en %s %s: %v\n", r.Method, r.URL.Path, r)
                
                // Stack trace
                stack := debug.Stack()
                fmt.Printf("Stack trace:\n%s\n", stack)
                
                // Respuesta de error
                http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
            }
        }()
        
        next(w, r)
    }
}

func handlerPeligroso(w http.ResponseWriter, r *http.Request) {
    // Simular error que causa panic
    if r.URL.Query().Get("error") == "panic" {
        panic("¡Error simulado!")
    }
    
    fmt.Fprintf(w, "Todo bien")
}

func main() {
    http.HandleFunc("/", recuperarMiddleware(handlerPeligroso))
    fmt.Println("Servidor en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Información de Stack Trace

```go
import (
    "fmt"
    "runtime"
    "runtime/debug"
)

func obtenerStackTrace() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Panic: %v\n", r)
            
            // Stack trace completo
            fmt.Printf("Stack trace:\n%s\n", debug.Stack())
            
            // Información específica del caller
            _, archivo, linea, ok := runtime.Caller(0)
            if ok {
                fmt.Printf("Archivo: %s, Línea: %d\n", archivo, linea)
            }
        }
    }()
    
    funcionQueHacePanic()
}

func funcionQueHacePanic() {
    panic("Error detallado")
}
```

## Tipos de Panic Values

```go
func ejemplosTiposPanic() {
    defer func() {
        if r := recover(); r != nil {
            switch v := r.(type) {
            case string:
                fmt.Printf("Panic con string: %s\n", v)
            case error:
                fmt.Printf("Panic con error: %v\n", v)
            case int:
                fmt.Printf("Panic con int: %d\n", v)
            default:
                fmt.Printf("Panic con tipo desconocido: %T, valor: %v\n", v, v)
            }
        }
    }()
    
    // Diferentes tipos de panic
    panic("string panic")
    // panic(errors.New("error panic"))
    // panic(42)
    // panic(struct{ msg string }{"struct panic"})
}
```

## Re-panic (Propagar Panic)

```go
func operacionCritica() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Limpieza después de panic: %v\n", r)
            
            // Realizar limpieza necesaria
            limpiarRecursos()
            
            // Re-lanzar el panic para que lo maneje un nivel superior
            panic(r)
        }
    }()
    
    // Operación que puede fallar
    realizarOperacionCritica()
}

func limpiarRecursos() {
    fmt.Println("Limpiando recursos...")
}

func realizarOperacionCritica() {
    panic("Fallo crítico del sistema")
}
```

## Cuándo Usar Panic vs Error

### ✅ Usar Panic Para:

1. **Errores de programación** (bugs)
2. **Situaciones imposibles de recuperar**
3. **Inicialización que no debe fallar**
4. **Violaciones de invariantes**

```go
func configurarServidor(config *Config) {
    if config == nil {
        panic("config no puede ser nil") // Error de programación
    }
    
    if config.Puerto <= 0 || config.Puerto > 65535 {
        panic("puerto inválido") // Configuración imposible
    }
}
```

### ❌ NO Usar Panic Para:

1. **Errores esperados** (archivo no encontrado, conexión perdida)
2. **Validación de entrada de usuario**
3. **Errores de red**
4. **Errores de E/O**

```go
// ❌ Mal uso de panic
func leerArchivo(nombre string) []byte {
    data, err := os.ReadFile(nombre)
    if err != nil {
        panic(err) // ¡No hagas esto!
    }
    return data
}

// ✅ Manejo apropiado con error
func leerArchivo(nombre string) ([]byte, error) {
    return os.ReadFile(nombre)
}
```

## Mejores Prácticas

### 1. Recover Solo en defer

```go
// ❌ No funciona
func malRecover() {
    if r := recover(); r != nil { // No funciona fuera de defer
        fmt.Println("Nunca se ejecuta")
    }
}

// ✅ Correcto
func buenRecover() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recuperado: %v\n", r)
        }
    }()
}
```

### 2. No Ignorar Panics

```go
// ❌ Ignorar panic silenciosamente
func ignorarPanic() {
    defer func() {
        recover() // ¡No hagas esto!
    }()
}

// ✅ Manejar apropiadamente
func manejarPanic() {
    defer func() {
        if r := recover(); r != nil {
            // Log del error
            log.Printf("Error recuperado: %v", r)
            
            // Posiblemente re-panic o convertir a error
            // panic(r) o return error
        }
    }()
}
```

### 3. Documentar Funciones que Pueden Hacer Panic

```go
// ParseInt convierte un string a entero.
// Hace panic si el string no es un número válido.
func ParseInt(s string) int {
    num, err := strconv.Atoi(s)
    if err != nil {
        panic(fmt.Sprintf("número inválido: %s", s))
    }
    return num
}
```

## Alternativas a Panic

### 1. Usar Must para Inicialización

```go
// Patrón "Must" para inicialización
func Must[T any](value T, err error) T {
    if err != nil {
        panic(err)
    }
    return value
}

// Uso
var re = Must(regexp.Compile(`\d+`))
```

### 2. Verificaciones de Invariantes

```go
func assert(condition bool, message string) {
    if !condition {
        panic(message)
    }
}

func procesarSlice(items []int) {
    assert(len(items) > 0, "slice no puede estar vacío")
    
    for i, item := range items {
        assert(item >= 0, fmt.Sprintf("item %d debe ser positivo", i))
        // procesar item
    }
}
```

`panic` y `recover` son herramientas poderosas pero deben usarse con moderación. Son más apropiados para errores excepcionales y situaciones de las que no es posible recuperarse normalmente. Para la mayoría de casos de error, el patrón idiomático de Go de retornar un valor de error es más apropiado.
