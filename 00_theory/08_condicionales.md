# Condicionales

Las estructuras condicionales permiten ejecutar código basado en el resultado de expresiones booleanas.

## If

La sentencia `if` evalúa una condición y ejecuta el bloque si es verdadera:

```go
if edad >= 18 {
    fmt.Println("Mayor de edad")
}
```

### If con Inicialización

Go permite declarar variables en la sentencia `if`, limitando su alcance al bloque:

```go
if edad := calcularEdad(nacimiento); edad >= 18 {
    fmt.Println("Mayor de edad:", edad)
}
// edad no está disponible aquí
```

## If-Else

Ejecuta un bloque u otro según la condición:

```go
if nota >= 60 {
    fmt.Println("Aprobado")
} else {
    fmt.Println("Reprobado")
}
```

## If-Else If-Else

Para múltiples condiciones excluyentes:

```go
if nota >= 90 {
    fmt.Println("Excelente")
} else if nota >= 70 {
    fmt.Println("Bueno")
} else if nota >= 60 {
    fmt.Println("Suficiente")
} else {
    fmt.Println("Insuficiente")
}
```

Las condiciones se evalúan en orden; se ejecuta solo el primer bloque cuya condición sea verdadera.

## Switch

El `switch` evalúa una expresión contra múltiples casos:

```go
dia := 3

switch dia {
case 1:
    fmt.Println("Lunes")
case 2:
    fmt.Println("Martes")
case 3:
    fmt.Println("Miércoles")
default:
    fmt.Println("Otro día")
}
```

### Switch sin Expresión

Equivalente a múltiples if-else, más legible:

```go
switch {
case nota >= 90:
    fmt.Println("A")
case nota >= 80:
    fmt.Println("B")
case nota >= 70:
    fmt.Println("C")
default:
    fmt.Println("F")
}
```

### Fallthrough

Por defecto, los casos no caen al siguiente. Use `fallthrough` para continuar:

```go
switch valor {
case 1:
    fmt.Println("Uno")
    fallthrough
case 2:
    fmt.Println("Dos")
}
// Si valor es 1, imprime: Uno, Dos
```

### Múltiples Valores en Case

```go
switch dia {
case 1, 2, 3, 4, 5:
    fmt.Println("Día laborable")
case 6, 7:
    fmt.Println("Fin de semana")
}
```

## Declaración de Variables en Switch

Similar a `if`, puede declarar variables:

```go
switch hora := time.Now().Hour(); {
case hora < 12:
    fmt.Println("Mañana")
case hora < 18:
    fmt.Println("Tarde")
default:
    fmt.Println("Noche")
}
```

## Tipo Switch

Evalúa el tipo dinámico de una interfaz:

```go
var valor interface{} = "texto"

switch v := valor.(type) {
case int:
    fmt.Println("Entero:", v)
case string:
    fmt.Println("Cadena:", v)
case bool:
    fmt.Println("Booleano:", v)
default:
    fmt.Printf("Tipo desconocido: %T\n", v)
}
```

## Resumen

- `if` puede incluir declaración de variables antes de la condición
- Las variables declaradas en `if` o `switch` tienen alcance limitado al bloque
- `switch` sin expresión actúa como cadena de if-else
- Los casos no caen automáticamente; use `fallthrough` explícitamente
- El `type switch` permite ramificar según el tipo de dato
