# Operadores de Comparación y Lógicos

Los operadores de comparación evalúan relaciones entre valores retornando un booleano. Los operadores lógicos combinan expresiones booleanas.

## Operadores de Comparación

| Operador | Descripción | Ejemplo |
|----------|-------------|---------|
| `==` | Igual a | `a == b` |
| `!=` | Distinto de | `a != b` |
| `<` | Menor que | `a < b` |
| `<=` | Menor o igual | `a <= b` |
| `>` | Mayor que | `a > b` |
| `>=` | Mayor o igual | `a >= b` |

```go
a, b := 5, 10

igual := a == b        // false
distinto := a != b     // true
menor := a < b         // true
mayorIgual := a >= b   // false
```

### Restricciones

- Los operandos deben ser del mismo tipo o tipos compatibles
- No se pueden comparar slices, maps ni funciones directamente
- Las cadenas se comparan léxicamente (orden alfabético)

```go
// Válido
"abc" < "def"          // true (orden alfabético)
"ABC" == "abc"         // false (distingue mayúsculas)

// Inválido: no compila
// []int{1, 2} == []int{1, 2}
```

## Operadores Lógicos

| Operador | Descripción | Ejemplo |
|----------|-------------|---------|
| `&&` | AND lógico | `a && b` |
| `||` | OR lógico | `a || b` |
| `!` | NOT lógico | `!a` |

```go
edad := 25
tieneLicencia := true

puedeConducir := edad >= 18 && tieneLicencia    // true
esMayor := edad >= 18 || edad <= 65              // true
noTieneLicencia := !tieneLicencia               // false
```

### Evaluación en Cortocircuito

Los operadores `&&` y `||` evalúan de izquierda a derecha y se detienen tan pronto como el resultado es determinable:

- `&&`: Si el operando izquierdo es `false`, no evalúa el derecho
- `||`: Si el operando izquierdo es `true`, no evalúa el derecho

```go
// La función derecha() no se ejecuta porque izquierda() retorna false
if izquierda() && derecha() {
    // ...
}
```

## Combinación de Operadores

Los operadores de comparación tienen menor precedencia que los aritméticos, y los lógicos tienen la menor precedencia de todos:

```go
// Evaluación: (a + b) > (c * d) && (e == f)
resultado := a + b > c * d && e == f

// Equivalente con paréntesis explícitos
resultado := ((a + b) > (c * d)) && (e == f)
```

Para legibilidad, use paréntesis cuando combine múltiples operadores lógicos:

```go
// Menos claro
if a > 0 && b > 0 || c > 0 && d > 0 {}

// Más claro
if (a > 0 && b > 0) || (c > 0 && d > 0) {}
```

## Comparación de Punto Flotante

Evite comparaciones exactas con valores de punto flotante debido a imprecisiones de representación:

```go
import "math"

// Incorrecto
if a == 0.1 + 0.2 {}    // Puede ser false

// Correcto: comparación con tolerancia
if math.Abs(a - (0.1 + 0.2)) < 0.0001 {}
```

## Resumen

- Use `==` para igualdad, `!=` para desigualdad
- Los operadores lógicos combinan condiciones: `&&` (y), `||` (o), `!` (negación)
- La evaluación en cortocircuito evita cálculos innecesarios
- No compare slices, maps ni funciones directamente
- Use paréntesis para clarificar expresiones complejas
- Para flotantes, compare con una tolerancia en lugar de igualdad exacta
