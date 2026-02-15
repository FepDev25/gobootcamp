# Bucles

Go tiene una única palabra clave para iteración: `for`. Esta palabra cubre todos los patrones de bucle mediante variantes de sintaxis.

## For Clásico

Sintaxis similar a C, con tres componentes separados por punto y coma:

```go
for inicializacion; condicion; post {
    // cuerpo del bucle
}
```

Ejemplo:

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)    // 0, 1, 2, 3, 4
}
```

## While

En Go no existe `while`; se logra omitiendo la inicialización y el post:

```go
contador := 0
for contador < 5 {
    fmt.Println(contador)
    contador++
}
```

## Bucle Infinito

Omitiendo todos los componentes:

```go
for {
    // ejecuta indefinidamente
    if condicionDeSalida {
        break
    }
}
```

Útil para servidores, lectura de streams o cuando la condición de terminación es compleja.

## Break y Continue

- `break`: Termina el bucle inmediatamente
- `continue`: Salta a la siguiente iteración

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue    // Salta el 3
    }
    if i == 7 {
        break       // Termina en 7
    }
    fmt.Println(i)  // 0, 1, 2, 4, 5, 6
}
```

## Range

La construcción `for range` itera sobre colecciones:

### Slices y Arrays

```go
numeros := []int{10, 20, 30}

for indice, valor := range numeros {
    fmt.Printf("[%d] = %d\n", indice, valor)
}
```

Si solo necesita el valor, use guión bajo para ignorar el índice:

```go
for _, valor := range numeros {
    fmt.Println(valor)
}
```

Si solo necesita el índice:

```go
for indice := range numeros {
    fmt.Println(indice)
}
```

### Maps

```go
edades := map[string]int{
    "Ana": 25,
    "Luis": 30,
}

for clave, valor := range edades {
    fmt.Printf("%s tiene %d años\n", clave, valor)
}
```

El orden de iteración sobre maps no está garantizado.

### Cadenas

Itera sobre runes (puntos de código Unicode):

```go
texto := "Go"
for indice, runa := range texto {
    fmt.Printf("[%d] = %c (U+%04X)\n", indice, runa, runa)
}
// [0] = G (U+0047)
// [1] = o (U+006F)
```

## Bucles Anidados y Etiquetas

Para salir de bucles externos, use etiquetas:

```go
externo:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i*j > 2 {
            break externo    // Sale de ambos bucles
        }
        fmt.Println(i, j)
    }
}
```

## Resumen

- `for` es la única palabra clave de iteración en Go
- Sintaxis clásica: `for i := 0; i < n; i++`
- Bucle while: `for condicion`
- Bucle infinito: `for`
- `range` itera sobre slices, maps, strings y channels
- Use `break` y `continue` para controlar el flujo
- Las etiquetas permiten romper bucles externos
