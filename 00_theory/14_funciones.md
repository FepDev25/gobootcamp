# Funciones

Las funciones son bloques de código reutilizables que realizan tareas específicas. Go soporta funciones con parámetros, valores de retorno y funciones como ciudadanos de primera clase.

## Declaración Básica

```go
func nombre(parametro tipo) tipoRetorno {
    // cuerpo de la función
    return valor
}
```

Ejemplo:

```go
func sumar(a int, b int) int {
    return a + b
}
```

## Parámetros

### Múltiples Parámetros del Mismo Tipo

Cuando parámetros consecutivos comparten tipo, se declara una vez al final:

```go
// Forma explícita
func rectangulo(base int, altura int) int

// Forma abreviada
func rectangulo(base, altura int) int
```

### Paso por Valor

Go pasa argumentos por valor. La función recibe una copia:

```go
func duplicar(x int) {
    x = x * 2
}

func main() {
    n := 5
    duplicar(n)
    fmt.Println(n)    // 5 (sin cambios)
}
```

Para modificar el original, use punteros (tema posterior).

## Valor de Retorno

La declaración especifica el tipo después de los parámetros:

```go
func cuadrado(x int) int {
    return x * x
}
```

## Funciones sin Retorno

Use la palabra clave `void` omitida (no retorna nada):

```go
func saludar(nombre string) {
    fmt.Println("Hola,", nombre)
}
```

## Funciones sin Parámetros

```go
func obtenerVersion() string {
    return "1.0.0"
}
```

## Funciones como Valores

Las funciones son tipos de primera clase, asignables a variables:

```go
var operacion func(int, int) int

operacion = sumar
resultado := operacion(3, 4)    // 7
```

## Funciones Anónimas

Funciones sin nombre, definidas en el lugar:

```go
func main() {
    suma := func(a, b int) int {
        return a + b
    }
    
    fmt.Println(suma(2, 3))    // 5
}
```

## Llamada Inmediata (IIFE)

Función anónima ejecutada al definirse:

```go
resultado := func(a, b int) int {
    return a + b
}(3, 4)    // 7
```

## Funciones como Parámetros

Pasar funciones como argumentos:

```go
func aplicar(a, b int, operacion func(int, int) int) int {
    return operacion(a, b)
}

func main() {
    resultado := aplicar(5, 3, func(x, y int) int {
        return x - y
    })
    fmt.Println(resultado)    // 2
}
```

## Recursividad

Funciones que se llaman a sí mismas:

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

Go optimiza recursión de cola cuando es posible, pero para profundidad grande considere iteración.

## Documentación

Go usa comentarios precedidos por el nombre de la función para documentación:

```go
// CalcularArea calcula el área de un rectángulo dados su base y altura.
// Retorna el área como float64.
func CalcularArea(base, altura float64) float64 {
    return base * altura
}
```

La herramienta `godoc` extrae estos comentarios para generar documentación.

## Resumen

- Declaración: `func nombre(param tipo) tipoRetorno`
- Parámetros consecutivos del mismo tipo pueden compartir la declaración de tipo
- Los argumentos se pasan por valor (copia)
- Las funciones son valores de primera clase, asignables a variables
- Las funciones anónimas permiten definir lógica en el lugar
- Use comentarios con el nombre de la función para documentación
