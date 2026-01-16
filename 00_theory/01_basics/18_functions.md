# Funciones en Go

Las funciones son bloques de código reutilizables que realizan tareas específicas. En Go, las funciones son ciudadanos de primera clase, lo que significa que pueden ser asignadas a variables, pasadas como argumentos y retornadas desde otras funciones.

## Declaración Básica de Funciones

### Sintaxis

```go
func nombreFuncion(parametros) tipoRetorno {
    // cuerpo de la función
    return valor
}
```

### Función Simple

```go
func saludar() {
    fmt.Println("¡Hola, mundo!")
}

func main() {
    saludar() // Llamada a la función
}
```

### Función con Parámetros

```go
func saludarPersona(nombre string) {
    fmt.Printf("¡Hola, %s!\n", nombre)
}

func main() {
    saludarPersona("Felipe")
    saludarPersona("María")
}
```

### Función con Retorno

```go
func sumar(a int, b int) int {
    return a + b
}

func main() {
    resultado := sumar(3, 4)
    fmt.Println("Suma:", resultado) // 7
}
```

## Parámetros de Funciones

### Parámetros del Mismo Tipo

```go
// Declaración larga
func multiplicar(a int, b int) int {
    return a * b
}

// Declaración corta (mismo tipo)
func dividir(a, b float64) float64 {
    return a / b
}
```

### Múltiples Parámetros de Diferentes Tipos

```go
func presentar(nombre string, edad int, activo bool) {
    fmt.Printf("Nombre: %s, Edad: %d, Activo: %t\n", nombre, edad, activo)
}

func main() {
    presentar("Alice", 25, true)
}
```

## Valores de Retorno

### Retorno Simple

```go
func obtenerPi() float64 {
    return 3.14159
}
```

### Múltiples Valores de Retorno

```go
func dividirConResto(a, b int) (int, int) {
    cociente := a / b
    resto := a % b
    return cociente, resto
}

func main() {
    c, r := dividirConResto(17, 5)
    fmt.Printf("17 ÷ 5 = %d con resto %d\n", c, r) // 17 ÷ 5 = 3 con resto 2
}
```

### Valores de Retorno Nombrados

```go
func calcularRectangulo(largo, ancho float64) (area, perimetro float64) {
    area = largo * ancho
    perimetro = 2 * (largo + ancho)
    return // return implícito de area y perimetro
}

func main() {
    a, p := calcularRectangulo(5.0, 3.0)
    fmt.Printf("Área: %.1f, Perímetro: %.1f\n", a, p)
}
```

### Ignorar Valores de Retorno

```go
func obtenerDatos() (string, int, error) {
    return "datos", 42, nil
}

func main() {
    nombre, _, err := obtenerDatos() // Ignora el segundo valor
    if err != nil {
        // manejar error
    }
    fmt.Println("Nombre:", nombre)
}
```

## Funciones Variádicas

Funciones que pueden recibir un número variable de argumentos:

```go
func sumarVarios(numeros ...int) int {
    total := 0
    for _, num := range numeros {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sumarVarios(1, 2, 3))          // 6
    fmt.Println(sumarVarios(10, 20, 30, 40))   // 100
    
    // Pasar slice como argumentos variádicos
    valores := []int{5, 10, 15, 20}
    fmt.Println(sumarVarios(valores...))       // 50
}
```

### Parámetros Mixtos

```go
func imprimirMensaje(prefijo string, mensajes ...string) {
    for _, mensaje := range mensajes {
        fmt.Printf("%s: %s\n", prefijo, mensaje)
    }
}

func main() {
    imprimirMensaje("INFO", "Sistema iniciado", "Base de datos conectada")
    imprimirMensaje("ERROR", "Conexión perdida")
}
```

## Funciones como Valores

### Asignar Función a Variable

```go
func sumar(a, b int) int {
    return a + b
}

func main() {
    // Asignar función a variable
    operacion := sumar
    resultado := operacion(5, 3)
    fmt.Println("Resultado:", resultado) // 8
}
```

### Funciones Anónimas

```go
func main() {
    // Función anónima asignada a variable
    saludar := func(nombre string) {
        fmt.Printf("¡Hola, %s!\n", nombre)
    }
    saludar("Ana")
    
    // Función anónima ejecutada inmediatamente
    func(mensaje string) {
        fmt.Println("Mensaje:", mensaje)
    }("¡Función anónima!")
}
```

## Funciones como Parámetros

### Función que Recibe Otra Función

```go
func aplicarOperacion(a, b int, operacion func(int, int) int) int {
    return operacion(a, b)
}

func sumar(a, b int) int     { return a + b }
func restar(a, b int) int    { return a - b }
func multiplicar(a, b int) int { return a * b }

func main() {
    fmt.Println(aplicarOperacion(10, 5, sumar))       // 15
    fmt.Println(aplicarOperacion(10, 5, restar))      // 5
    fmt.Println(aplicarOperacion(10, 5, multiplicar)) // 50
    
    // Con función anónima
    resultado := aplicarOperacion(10, 5, func(a, b int) int {
        return a % b
    })
    fmt.Println("Módulo:", resultado) // 0
}
```

## Funciones que Retornan Funciones (Closures)

```go
func crearMultiplicador(factor int) func(int) int {
    return func(numero int) int {
        return numero * factor
    }
}

func main() {
    duplicar := crearMultiplicador(2)
    triplicar := crearMultiplicador(3)
    
    fmt.Println("5 * 2 =", duplicar(5))   // 10
    fmt.Println("5 * 3 =", triplicar(5))  // 15
}
```

### Closure con Estado

```go
func crearContador() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    contador1 := crearContador()
    contador2 := crearContador()
    
    fmt.Println(contador1()) // 1
    fmt.Println(contador1()) // 2
    fmt.Println(contador2()) // 1 (independiente)
    fmt.Println(contador1()) // 3
}
```

## Recursión

Funciones que se llaman a sí mismas:

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    fmt.Println("5! =", factorial(5))    // 120
    fmt.Println("Fibonacci(7) =", fibonacci(7)) // 13
}
```

## Manejo de Errores en Funciones

### Patrón Idiomático

```go
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("división por cero")
    }
    return a / b, nil
}

func main() {
    resultado, err := dividir(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Resultado:", resultado)
}
```

### Función con Múltiples Posibles Errores

```go
func procesarUsuario(nombre string, edad int) (*Usuario, error) {
    if nombre == "" {
        return nil, errors.New("nombre no puede estar vacío")
    }
    
    if edad < 0 {
        return nil, errors.New("edad no puede ser negativa")
    }
    
    if edad > 150 {
        return nil, errors.New("edad no puede ser mayor a 150")
    }
    
    return &Usuario{
        Nombre: nombre,
        Edad:   edad,
    }, nil
}
```

## Métodos en Structs

```go
type Rectangulo struct {
    Largo, Ancho float64
}

// Método con receptor de valor
func (r Rectangulo) Area() float64 {
    return r.Largo * r.Ancho
}

// Método con receptor de puntero
func (r *Rectangulo) Escalar(factor float64) {
    r.Largo *= factor
    r.Ancho *= factor
}

func main() {
    rect := Rectangulo{Largo: 5, Ancho: 3}
    
    fmt.Println("Área:", rect.Area()) // 15
    
    rect.Escalar(2)
    fmt.Println("Nueva área:", rect.Area()) // 60
}
```

## Funciones de Orden Superior

### Map

```go
func mapear(slice []int, fn func(int) int) []int {
    resultado := make([]int, len(slice))
    for i, v := range slice {
        resultado[i] = fn(v)
    }
    return resultado
}

func main() {
    numeros := []int{1, 2, 3, 4, 5}
    
    cuadrados := mapear(numeros, func(x int) int {
        return x * x
    })
    
    fmt.Println("Cuadrados:", cuadrados) // [1 4 9 16 25]
}
```

### Filter

```go
func filtrar(slice []int, predicado func(int) bool) []int {
    var resultado []int
    for _, v := range slice {
        if predicado(v) {
            resultado = append(resultado, v)
        }
    }
    return resultado
}

func main() {
    numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    pares := filtrar(numeros, func(x int) bool {
        return x%2 == 0
    })
    
    fmt.Println("Números pares:", pares) // [2 4 6 8 10]
}
```

### Reduce

```go
func reducir(slice []int, fn func(int, int) int, inicial int) int {
    resultado := inicial
    for _, v := range slice {
        resultado = fn(resultado, v)
    }
    return resultado
}

func main() {
    numeros := []int{1, 2, 3, 4, 5}
    
    suma := reducir(numeros, func(acc, x int) int {
        return acc + x
    }, 0)
    
    producto := reducir(numeros, func(acc, x int) int {
        return acc * x
    }, 1)
    
    fmt.Println("Suma:", suma)         // 15
    fmt.Println("Producto:", producto) // 120
}
```

## Ejemplos Prácticos Completos

### 1. Calculadora Funcional

```go
type Calculadora struct{}

func (c Calculadora) Sumar(a, b float64) float64 {
    return a + b
}

func (c Calculadora) Restar(a, b float64) float64 {
    return a - b
}

func (c Calculadora) Multiplicar(a, b float64) float64 {
    return a * b
}

func (c Calculadora) Dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("división por cero")
    }
    return a / b, nil
}

func main() {
    calc := Calculadora{}
    
    fmt.Println("Suma:", calc.Sumar(10, 5))
    fmt.Println("Resta:", calc.Restar(10, 5))
    fmt.Println("Multiplicación:", calc.Multiplicar(10, 5))
    
    if resultado, err := calc.Dividir(10, 5); err == nil {
        fmt.Println("División:", resultado)
    }
}
```

### 2. Sistema de Validación

```go
type Validador func(string) error

func validarLongitud(min, max int) Validador {
    return func(valor string) error {
        if len(valor) < min {
            return fmt.Errorf("debe tener al menos %d caracteres", min)
        }
        if len(valor) > max {
            return fmt.Errorf("no debe tener más de %d caracteres", max)
        }
        return nil
    }
}

func validarEmail(email string) error {
    if !strings.Contains(email, "@") {
        return errors.New("email debe contener @")
    }
    return nil
}

func validar(valor string, validadores ...Validador) []error {
    var errores []error
    for _, validador := range validadores {
        if err := validador(valor); err != nil {
            errores = append(errores, err)
        }
    }
    return errores
}

func main() {
    email := "test"
    
    errores := validar(email,
        validarLongitud(5, 50),
        validarEmail,
    )
    
    for _, err := range errores {
        fmt.Println("Error:", err)
    }
}
```

## Mejores Prácticas

### 1. Nombres Descriptivos

```go
// ❌ Poco descriptivo
func calc(a, b int) int { return a + b }

// ✅ Descriptivo
func calcularSuma(a, b int) int { return a + b }
```

### 2. Funciones Pequeñas y Enfocadas

```go
// ❌ Función que hace demasiado
func procesarUsuario(datos string) {
    // validar
    // parsear
    // guardar en DB
    // enviar email
    // generar reporte
}

// ✅ Funciones pequeñas y específicas
func validarDatos(datos string) error { /*...*/ }
func parsearDatos(datos string) Usuario { /*...*/ }
func guardarUsuario(u Usuario) error { /*...*/ }
func enviarEmail(u Usuario) error { /*...*/ }
```

### 3. Manejo Consistente de Errores

```go
// ✅ Patrón consistente
func operacionRiesgosa() (resultado string, err error) {
    // hacer algo que puede fallar
    if err != nil {
        return "", fmt.Errorf("operación falló: %w", err)
    }
    return resultado, nil
}
```

### 4. Usar Interfaces para Mayor Flexibilidad

```go
type Escritor interface {
    Escribir(datos []byte) error
}

func guardarDatos(w Escritor, datos []byte) error {
    return w.Escribir(datos)
}
```

Las funciones en Go son extremadamente flexibles y potentes. Su capacidad de ser tratadas como valores de primera clase permite patrones de programación funcional muy elegantes y expresivos, mientras que su sintaxis simple mantiene el código legible y mantenible.
