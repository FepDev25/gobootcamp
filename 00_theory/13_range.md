# Range

La construcción `range` proporciona una forma idiomática de iterar sobre colecciones en Go. Funciona con arrays, slices, maps, strings y channels.

## Sintaxis General

```go
for indice, valor := range coleccion {
    // usar indice y valor
}
```

Puede omitir el índice con guión bajo si no lo necesita.

## Arrays y Slices

Retorna índice y valor:

```go
numeros := []int{10, 20, 30}

for i, v := range numeros {
    fmt.Printf("numeros[%d] = %d\n", i, v)
}
// numeros[0] = 10
// numeros[1] = 20
// numeros[2] = 30

// Solo valor
for _, v := range numeros {
    fmt.Println(v)
}

// Solo índice
for i := range numeros {
    fmt.Println(i)
}
```

## Maps

Retorna clave y valor. El orden no está garantizado:

```go
edades := map[string]int{
    "Ana": 25,
    "Luis": 30,
}

for nombre, edad := range edades {
    fmt.Printf("%s tiene %d años\n", nombre, edad)
}

// Solo clave
for nombre := range edades {
    fmt.Println(nombre)
}
```

## Cadenas

Itera sobre runes (puntos de código Unicode), no bytes:

```go
texto := "Go"

for indice, runa := range texto {
    fmt.Printf("[%d] = %c (código: %d)\n", indice, runa, runa)
}
// [0] = G (código: 71)
// [1] = o (código: 111)
```

Con caracteres multibyte, los índices son desplazamientos en bytes:

```go
texto := "日本"

for i, r := range texto {
    fmt.Printf("[%d] = %c\n", i, r)
}
// [0] = 日
// [3] = 本
```

## Channels

Retorna valores recibidos hasta que el channel se cierra:

```go
ch := make(chan int)

go func() {
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
}()

for valor := range ch {
    fmt.Println(valor)
}
// 1, 2, 3
```

## Copia de Valores

`range` crea una copia del valor en cada iteración. Modificar la variable del range no afecta la colección:

```go
numeros := []int{1, 2, 3}

for _, n := range numeros {
    n = n * 10    // No modifica el slice
}
fmt.Println(numeros)    // [1 2 3]

// Para modificar, use el índice
for i := range numeros {
    numeros[i] = numeros[i] * 10
}
fmt.Println(numeros)    // [10 20 30]
```

## Range con Break y Continue

Funcionan como en cualquier bucle for:

```go
for i, v := range numeros {
    if v == 0 {
        continue    // Saltar este elemento
    }
    if i > 10 {
        break       // Terminar iteración
    }
    fmt.Println(v)
}
```

## Consideraciones de Rendimiento

- `range` sobre slices crea una copia del elemento
- Para grandes estructuras, usar índice para evitar copias:

```go
// Menos eficiente con estructuras grandes
for _, item := range items {
    procesar(item)    // Copia item
}

// Más eficiente
for i := range items {
    procesar(items[i])    // Sin copia
}
```

## Resumen

- `range` itera sobre arrays, slices, maps, strings y channels
- Arrays y slices: retorna índice y valor
- Maps: retorna clave y valor (orden no garantizado)
- Strings: retorna índice en bytes y rune
- Channels: retorna valores hasta que se cierra
- Modificar la variable de range no afecta la colección; usar índice para modificar
