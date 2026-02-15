# Funciones Avanzadas

Go extiende las capacidades de las funciones con múltiples valores de retorno, parámetros variádicos y funciones nombradas de retorno.

## Múltiples Valores de Retorno

Las funciones pueden retornar más de un valor, separados por comas:

```go
func dividir(dividendo, divisor float64) (float64, error) {
    if divisor == 0 {
        return 0, errors.New("división por cero")
    }
    return dividendo / divisor, nil
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

### Ignorar Valores de Retorno

Use guión bajo para ignorar valores no necesarios:

```go
resultado, _ := dividir(10, 2)    // Ignora el error
```

## Retornos Nombrados

Asigne nombres a los valores de retorno para usarlos como variables locales:

```go
func rectangulo(base, altura float64) (area, perimetro float64) {
    area = base * altura
    perimetro = 2 * (base + altura)
    return    // Retorno implícito de las variables nombradas
}
```

El `return` vacío retorna los valores actuales de las variables nombradas. Aunque válido, `return explicito` suele ser más legible:

```go
func rectangulo(base, altura float64) (area, perimetro float64) {
    area = base * altura
    perimetro = 2 * (base + altura)
    return area, perimetro
}
```

## Funciones Variádicas

Aceptan número variable de argumentos del mismo tipo. El parámetro variádico usa `...` y debe ser el último:

```go
func sumar(numeros ...int) int {
    total := 0
    for _, n := range numeros {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sumar(1, 2, 3))      // 6
    fmt.Println(sumar(10, 20))       // 30
    fmt.Println(sumar())             // 0
}
```

### Pasar Slice a Función Variádica

Expandir slice con `...`:

```go
valores := []int{1, 2, 3, 4, 5}
resultado := sumar(valores...)    // 15
```

### Múltiples Parámetros con Variádico

```go
func imprimir(prefijo string, valores ...interface{}) {
    for _, v := range valores {
        fmt.Printf("%s: %v\n", prefijo, v)
    }
}

imprimir("INFO", "Iniciando", 42, true)
```

## Retornos en Condicionales

El patrón común de retorno temprano mejora la legibilidad:

```go
func obtenerUsuario(id int) (*Usuario, error) {
    if id <= 0 {
        return nil, errors.New("ID inválido")
    }
    
    usuario, err := db.Query(id)
    if err != nil {
        return nil, err
    }
    
    return usuario, nil
}
```

## Funciones como Campos

Las estructuras pueden contener funciones:

```go
type Calculadora struct {
    Operacion func(a, b float64) float64
}

calc := Calculadora{
    Operacion: func(a, b float64) float64 {
        return a + b
    },
}
resultado := calc.Operacion(5, 3)
```

## Defer en Funciones

La palabra clave `defer` pospone la ejecución de una función hasta que la función que la contiene retorne:

```go
func procesarArchivo(nombre string) error {
    archivo, err := os.Open(nombre)
    if err != nil {
        return err
    }
    defer archivo.Close()    // Se ejecuta al retornar
    
    // Procesar archivo...
    return nil
}
```

Los argumentos de `defer` se evalúan inmediatamente, pero la función se ejecuta al final.

## Resumen

- Las funciones pueden retornar múltiples valores separados por comas
- Los retornos nombrados actúan como variables locales; `return` vacío los retorna
- Parámetros variádicos (`...tipo`) aceptan número variable de argumentos
- Expanda slices con `...` para pasarlos a funciones variádicas
- `defer` ejecuta funciones al finalizar la función contenedora
- El patrón de retorno temprano simplifica el flujo de error
