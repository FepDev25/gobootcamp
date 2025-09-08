# Bucles (Loops) en Go

Go tiene un único tipo de bucle: el bucle `for`. Sin embargo, es muy versátil y puede usarse de diferentes maneras para cubrir todos los casos de uso de bucles tradicionales.

## Sintaxis Básica del Bucle `for`

### 1. Bucle `for` Tradicional (Estilo C)

```go
for inicialización; condición; incremento {
    // código a ejecutar
}
```

**Ejemplo:**
```go
for i := 0; i <= 5; i++ {
    fmt.Println("Iteración:", i)
}
// Salida: Iteración: 0, 1, 2, 3, 4, 5
```

### 2. Bucle `for` como `while`

```go
for condición {
    // código a ejecutar
}
```

**Ejemplo:**
```go
i := 0
for i < 5 {
    fmt.Println("i =", i)
    i++
}
```

### 3. Bucle Infinito

```go
for {
    // código que se ejecuta infinitamente
    // usar break para salir
}
```

**Ejemplo:**
```go
counter := 0
for {
    if counter >= 3 {
        break
    }
    fmt.Println("Counter:", counter)
    counter++
}
```

## Bucle `for` con `range`

El `range` se usa para iterar sobre colecciones como arrays, slices, maps, strings y channels.

### Iterando sobre Slices/Arrays

```go
numbers := []int{1, 2, 3, 4, 5}

// Con índice y valor
for index, value := range numbers {
    fmt.Printf("Índice: %d, Valor: %d\n", index, value)
}

// Solo el valor (ignorar índice)
for _, value := range numbers {
    fmt.Println("Valor:", value)
}

// Solo el índice
for index := range numbers {
    fmt.Println("Índice:", index)
}
```

### Iterando sobre Maps

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}

for name, age := range ages {
    fmt.Printf("%s tiene %d años\n", name, age)
}
```

### Iterando sobre Strings

```go
text := "Hola"
for index, char := range text {
    fmt.Printf("Posición %d: %c\n", index, char)
}
// Salida:
// Posición 0: H
// Posición 1: o
// Posición 2: l
// Posición 3: a
```

### Go 1.22+: Range sobre Enteros

```go
// Iterar de 0 a n-1
for i := range 5 {
    fmt.Println("i =", i) // 0, 1, 2, 3, 4
}
```

## Control de Flujo en Bucles

### `break` - Terminar el Bucle

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Sale del bucle cuando i es 5
    }
    fmt.Println(i)
}
// Salida: 0, 1, 2, 3, 4
```

### `continue` - Saltar Iteración

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Salta los números pares
    }
    fmt.Println("Número impar:", i)
}
// Salida: 1, 3, 5, 7, 9
```

## Ejemplos Prácticos

### 1. Encontrar un Número Secreto

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    secretNumber := rand.Intn(21) // 0-20
    
    for i := range 21 {
        if i == secretNumber {
            fmt.Println("¡Encontrado el número secreto:", secretNumber)
            break
        }
        
        if i%2 == 0 {
            continue // Saltar números pares
        }
        
        fmt.Println("Probando número impar:", i)
    }
}
```

### 2. Procesamiento de Datos

```go
func processNumbers(numbers []int) {
    sum := 0
    count := 0
    
    for _, num := range numbers {
        if num < 0 {
            continue // Ignorar números negativos
        }
        
        sum += num
        count++
        
        if count >= 10 {
            break // Procesar máximo 10 números
        }
    }
    
    if count > 0 {
        average := float64(sum) / float64(count)
        fmt.Printf("Promedio de %d números: %.2f\n", count, average)
    }
}
```

### 3. Bucle While Simulado

```go
func countdown() {
    i := 10
    for i > 0 {
        fmt.Println("Cuenta regresiva:", i)
        i--
        time.Sleep(time.Second)
    }
    fmt.Println("¡Despegue!")
}
```

### 4. Procesamiento de Menú

```go
func showMenu() {
    for {
        fmt.Println("\n--- MENÚ ---")
        fmt.Println("1. Opción A")
        fmt.Println("2. Opción B") 
        fmt.Println("3. Salir")
        
        var choice int
        fmt.Print("Seleccione una opción: ")
        fmt.Scanln(&choice)
        
        switch choice {
        case 1:
            fmt.Println("Ejecutando Opción A")
        case 2:
            fmt.Println("Ejecutando Opción B")
        case 3:
            fmt.Println("¡Adiós!")
            return // Sale de la función y termina el bucle
        default:
            fmt.Println("Opción inválida")
        }
    }
}
```

## Bucles Anidados

```go
// Tabla de multiplicar
for i := 1; i <= 5; i++ {
    for j := 1; j <= 5; j++ {
        fmt.Printf("%d x %d = %d\t", i, j, i*j)
    }
    fmt.Println() // Nueva línea después de cada fila
}
```

### Etiquetas (Labels) para Bucles Anidados

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer // Sale de ambos bucles
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

## Consideraciones de Rendimiento

### 1. Evita Cálculos Innecesarios en la Condición

```go
// ❌ Malo: len() se ejecuta en cada iteración
for i := 0; i < len(slice); i++ {
    // procesamiento
}

// ✅ Mejor: calcula len() una sola vez
length := len(slice)
for i := 0; i < length; i++ {
    // procesamiento
}

// ✅ Mejor aún: usa range cuando sea apropiado
for i, value := range slice {
    // procesamiento
}
```

### 2. Range vs Índice

```go
// Para acceso secuencial, range es más idiomático
for _, value := range slice {
    process(value)
}

// Para índices específicos, usa bucle tradicional
for i := len(slice) - 1; i >= 0; i-- {
    process(slice[i])
}
```

## Errores Comunes

### 1. Variable de Loop Capturada en Goroutines

```go
// ❌ Problemático
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i) // Posiblemente imprime 5 cinco veces
    }()
}

// ✅ Correcto
for i := 0; i < 5; i++ {
    go func(val int) {
        fmt.Println(val) // Imprime 0, 1, 2, 3, 4
    }(i)
}
```

### 2. Modificar Slice Durante Iteración

```go
// ❌ Problemático
numbers := []int{1, 2, 3, 4, 5}
for i, num := range numbers {
    if num%2 == 0 {
        numbers = append(numbers[:i], numbers[i+1:]...) // Modifica durante iteración
    }
}

// ✅ Mejor: iterar hacia atrás o usar un nuevo slice
var odds []int
for _, num := range numbers {
    if num%2 != 0 {
        odds = append(odds, num)
    }
}
```

## Resumen de Mejores Prácticas

1. **Usa `range` cuando sea posible** - Es más idiomático y seguro
2. **Prefiere `for` simple para bucles infinitos o condiciones complejas**
3. **Usa etiquetas solo cuando sea necesario** para bucles anidados complejos
4. **Ten cuidado con las variables capturadas** en goroutines
5. **Evita modificar colecciones durante la iteración**
6. **Usa `break` y `continue` para control de flujo claro**
7. **Considera el rendimiento** al elegir entre diferentes enfoques

Los bucles en Go son simples pero poderosos. El bucle `for` único y versátil, combinado con `range`, proporciona toda la funcionalidad necesaria para iterar eficientemente sobre datos y controlar el flujo de tu programa.
