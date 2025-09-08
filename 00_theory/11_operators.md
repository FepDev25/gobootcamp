# Operadores en Go

Go proporciona varios tipos de operadores para realizar diferentes operaciones sobre variables y valores. Los operadores se clasifican en varias categorías según su funcionalidad.

## Operadores Aritméticos

Realizan operaciones matemáticas básicas:

| Operador | Descripción | Ejemplo | Resultado |
|----------|-------------|---------|-----------|
| `+` | Suma | `5 + 3` | `8` |
| `-` | Resta | `5 - 3` | `2` |
| `*` | Multiplicación | `5 * 3` | `15` |
| `/` | División | `15 / 3` | `5` |
| `%` | Módulo (residuo) | `17 % 5` | `2` |

```go
package main

import "fmt"

func main() {
    a, b := 17, 5
    
    fmt.Println("Suma:", a + b)          // 22
    fmt.Println("Resta:", a - b)         // 12
    fmt.Println("Multiplicación:", a * b) // 85
    fmt.Println("División:", a / b)       // 3 (división entera)
    fmt.Println("Módulo:", a % b)         // 2
    
    // División con flotantes
    x, y := 17.0, 5.0
    fmt.Println("División float:", x / y) // 3.4
}
```

## Operadores de Asignación

Asignan valores a variables:

| Operador | Descripción | Equivalente | Ejemplo |
|----------|-------------|-------------|---------|
| `=` | Asignación simple | | `x = 5` |
| `+=` | Suma y asigna | `x = x + y` | `x += 3` |
| `-=` | Resta y asigna | `x = x - y` | `x -= 3` |
| `*=` | Multiplica y asigna | `x = x * y` | `x *= 3` |
| `/=` | Divide y asigna | `x = x / y` | `x /= 3` |
| `%=` | Módulo y asigna | `x = x % y` | `x %= 3` |

```go
func main() {
    x := 10
    
    x += 5    // x = 15
    x -= 3    // x = 12
    x *= 2    // x = 24
    x /= 4    // x = 6
    x %= 5    // x = 1
    
    fmt.Println("Resultado final:", x) // 1
}
```

## Operadores de Comparación

Comparan dos valores y devuelven un booleano:

| Operador | Descripción | Ejemplo | Resultado |
|----------|-------------|---------|-----------|
| `==` | Igual a | `5 == 5` | `true` |
| `!=` | No igual a | `5 != 3` | `true` |
| `<` | Menor que | `3 < 5` | `true` |
| `<=` | Menor o igual | `5 <= 5` | `true` |
| `>` | Mayor que | `7 > 3` | `true` |
| `>=` | Mayor o igual | `5 >= 5` | `true` |

```go
func main() {
    a, b := 10, 20
    
    fmt.Println("a == b:", a == b)  // false
    fmt.Println("a != b:", a != b)  // true
    fmt.Println("a < b:", a < b)    // true
    fmt.Println("a <= b:", a <= b)  // true
    fmt.Println("a > b:", a > b)    // false
    fmt.Println("a >= b:", a >= b)  // false
}
```

## Operadores Lógicos

Realizan operaciones lógicas con valores booleanos:

| Operador | Descripción | Ejemplo | Resultado |
|----------|-------------|---------|-----------|
| `&&` | AND lógico | `true && false` | `false` |
| `\|\|` | OR lógico | `true \|\| false` | `true` |
| `!` | NOT lógico | `!true` | `false` |

```go
func main() {
    a, b := true, false
    
    fmt.Println("a && b:", a && b)  // false
    fmt.Println("a || b:", a || b)  // true
    fmt.Println("!a:", !a)          // false
    fmt.Println("!b:", !b)          // true
    
    // Uso práctico
    age := 25
    hasLicense := true
    
    canDrive := age >= 18 && hasLicense
    fmt.Println("Puede conducir:", canDrive) // true
}
```

## Operadores Bitwise

Operan a nivel de bits:

| Operador | Descripción | Ejemplo | Resultado |
|----------|-------------|---------|-----------|
| `&` | AND bitwise | `5 & 3` | `1` |
| `\|` | OR bitwise | `5 \| 3` | `7` |
| `^` | XOR bitwise | `5 ^ 3` | `6` |
| `&^` | AND NOT bitwise | `5 &^ 3` | `4` |
| `<<` | Desplazamiento izq. | `5 << 1` | `10` |
| `>>` | Desplazamiento der. | `5 >> 1` | `2` |

```go
func main() {
    a, b := 5, 3  // a = 101, b = 011 en binario
    
    fmt.Printf("a & b = %d\n", a & b)   // 1 (001)
    fmt.Printf("a | b = %d\n", a | b)   // 7 (111)
    fmt.Printf("a ^ b = %d\n", a ^ b)   // 6 (110)
    fmt.Printf("a &^ b = %d\n", a &^ b) // 4 (100)
    fmt.Printf("a << 1 = %d\n", a << 1) // 10 (1010)
    fmt.Printf("a >> 1 = %d\n", a >> 1) // 2 (10)
}
```

## Operadores Unarios

Operan sobre un solo operando:

| Operador | Descripción | Ejemplo |
|----------|-------------|---------|
| `+` | Positivo unario | `+x` |
| `-` | Negativo unario | `-x` |
| `!` | NOT lógico | `!x` |
| `^` | NOT bitwise | `^x` |
| `*` | Desreferencia | `*ptr` |
| `&` | Dirección de | `&var` |

```go
func main() {
    x := 10
    y := true
    
    fmt.Println("+x:", +x)  // 10
    fmt.Println("-x:", -x)  // -10
    fmt.Println("!y:", !y)  // false
    fmt.Println("^x:", ^x)  // -11 (complemento bitwise)
    
    // Punteros
    ptr := &x
    fmt.Println("&x:", ptr)   // Dirección de memoria
    fmt.Println("*ptr:", *ptr) // 10 (valor apuntado)
}
```

## Operadores de Incremento y Decremento

**Nota importante**: En Go, `++` y `--` son **declaraciones**, no expresiones.

```go
func main() {
    x := 5
    
    x++  // Equivale a: x = x + 1
    fmt.Println(x) // 6
    
    x--  // Equivale a: x = x - 1
    fmt.Println(x) // 5
    
    // ❌ ESTO NO ES VÁLIDO EN GO:
    // y := x++  // Error: no es una expresión
    // fmt.Println(++x)  // Error: solo postfijo
}
```

## Precedencia de Operadores

Del mayor al menor precedencia:

1. `*`, `/`, `%`, `<<`, `>>`, `&`, `&^`
2. `+`, `-`, `|`, `^`
3. `==`, `!=`, `<`, `<=`, `>`, `>=`
4. `&&`
5. `||`

```go
func main() {
    // Sin paréntesis
    result1 := 2 + 3 * 4  // = 2 + 12 = 14
    
    // Con paréntesis para cambiar precedencia
    result2 := (2 + 3) * 4  // = 5 * 4 = 20
    
    fmt.Println("2 + 3 * 4 =", result1)    // 14
    fmt.Println("(2 + 3) * 4 =", result2)  // 20
}
```

## Operador de Canal (Channel)

Para trabajar con goroutines y canales:

```go
// <- operador de canal
ch <- value    // Enviar valor al canal
value := <-ch  // Recibir valor del canal
```

## Casos de Uso Comunes

### Validación de Rangos
```go
func isValidAge(age int) bool {
    return age >= 0 && age <= 150
}
```

### Operaciones Bitwise para Flags
```go
const (
    Read = 1 << iota  // 1
    Write            // 2
    Execute          // 4
)

func hasPermission(userPerms, requiredPerm int) bool {
    return (userPerms & requiredPerm) == requiredPerm
}
```

### Cálculos Matemáticos
```go
func calculateArea(length, width float64) float64 {
    return length * width
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
```

## Consideraciones Especiales

### División por Cero
```go
// ❌ Esto causará panic en runtime
x := 10 / 0  // panic: integer divide by zero

// ✅ Mejor práctica: verificar antes de dividir
func safeDivide(a, b int) int {
    if b == 0 {
        return 0  // o manejar el error apropiadamente
    }
    return a / b
}
```

### Overflow
```go
var x int8 = 127
x++  // x se convierte en -128 (overflow)
```

Los operadores en Go son fundamentales para manipular datos y controlar el flujo de tus programas. Entender su precedencia y uso correcto es esencial para escribir código eficiente y libre de errores.
