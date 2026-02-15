# Operadores Aritméticos

Los operadores aritméticos permiten realizar operaciones matemáticas básicas sobre valores numéricos.

## Operadores Básicos

| Operador | Descripción | Ejemplo | Resultado |
|----------|-------------|---------|-----------|
| `+` | Suma | `5 + 3` | `8` |
| `-` | Resta | `5 - 3` | `2` |
| `*` | Multiplicación | `5 * 3` | `15` |
| `/` | División | `5 / 3` | `1` (enteros) |
| `%` | Módulo (resto) | `5 % 3` | `2` |

```go
a, b := 10, 3

suma := a + b        // 13
resta := a - b       // 7
multiplicacion := a * b  // 30
division := a / b    // 3 (división entera)
resto := a % b       // 1
```

## División con Tipos Diferentes

La división entre enteros trunca el resultado. Para obtener decimales, al menos un operando debe ser flotante:

```go
var a int = 5
var b int = 2

resultado1 := a / b       // 2 (entero)
resultado2 := float64(a) / float64(b)  // 2.5
```

## Operadores de Asignación Compuesta

| Operador | Equivalente | Descripción |
|----------|-------------|-------------|
| `+=` | `a = a + b` | Suma y asigna |
| `-=` | `a = a - b` | Resta y asigna |
| `*=` | `a = a * b` | Multiplica y asigna |
| `/=` | `a = a / b` | Divide y asigna |
| `%=` | `a = a % b` | Módulo y asigna |

```go
contador := 10
contador += 5    // 15
contador -= 3    // 12
contador *= 2    // 24
contador /= 4    // 6
```

## Operadores de Incremento y Decremento

Go proporciona operadores unarios para modificar una variable en una unidad:

| Operador | Descripción | Ejemplo |
|----------|-------------|---------|
| `++` | Incremento en 1 | `i++` |
| `--` | Decremento en 1 | `i--` |

```go
contador := 0
contador++       // 1
contador--       // 0
```

Estos operadores son sentencias, no expresiones. No pueden usarse dentro de otras expresiones:

```go
// Válido
i++

// Inválido: no compila
// j = i++
// fmt.Println(i++)
```

## Precedencia de Operadores

De mayor a menor prioridad:

1. `*`, `/`, `%`
2. `+`, `-`
3. `=`, `+=`, `-=`, etc.

Use paréntesis para clarificar o modificar el orden:

```go
resultado := 10 + 5 * 2      // 20 (multiplicación primero)
resultado := (10 + 5) * 2    // 30 (paréntesis primero)
```

## Operaciones con Desbordamiento

Para enteros, las operaciones que exceden el rango del tipo producen desbordamiento (wrapping):

```go
var max uint8 = 255
max++    // Desborda a 0
```

El compilador detecta constantes que desbordan su tipo, pero no variables.

## Resumen

- Los operandos deben ser del mismo tipo o convertidos explícitamente
- División entera trunca decimales; use conversión a flotante para precisión
- `++` y `--` son sentencias, no retornan valores
- Use paréntesis para controlar la precedencia de operaciones
- Vigile el desbordamiento en tipos de tamaño fijo
