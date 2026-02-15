# Maps

Los maps son colecciones desordenadas de pares clave-valor. Las claves son únicas y permiten búsqueda, inserción y eliminación eficientes.

## Declaración e Inicialización

```go
// Declaración (nil, no usable)
var edades map[string]int

// Inicialización con make
edades = make(map[string]int)

// Literal
edades := map[string]int{
    "Ana": 25,
    "Luis": 30,
    "María": 28,
}
```

## Operaciones Básicas

### Inserción y Actualización

```go
edades["Pedro"] = 35    // Insertar
edades["Ana"] = 26      // Actualizar
```

### Acceso

```go
edad := edades["Ana"]    // 26
```

Si la clave no existe, retorna el valor cero del tipo:

```go
edad := edades["Carlos"]    // 0 (valor cero de int)
```

### Comprobar Existencia

El segundo valor indica si la clave existe:

```go
edad, existe := edades["Carlos"]
if existe {
    fmt.Println("Edad:", edad)
} else {
    fmt.Println("No encontrado")
}

// Forma idiomática
if edad, ok := edades["Carlos"]; ok {
    fmt.Println("Edad:", edad)
}
```

### Eliminación

```go
delete(edades, "Luis")
```

Eliminar una clave inexistente no produce error.

## Tipos de Clave Permitidos

Las claves deben ser comparables: no pueden ser slices, maps ni funciones.

```go
// Válidos
map[string]int
map[int]string
map[bool]float64

// Inválidos (no compilan)
// map[[]int]string
// map[map[string]int]bool
```

## Recorrido

El orden de iteración no está garantizado:

```go
for clave, valor := range edades {
    fmt.Printf("%s: %d años\n", clave, valor)
}

// Solo claves
for clave := range edades {
    fmt.Println(clave)
}
```

## Longitud

```go
cantidad := len(edades)
```

## Maps Anidados

```go
// Map de maps
empleados := map[string]map[string]string{
    "E001": {
        "nombre": "Ana",
        "departamento": "Ventas",
    },
    "E002": {
        "nombre": "Luis",
        "departamento": "IT",
    },
}

depto := empleados["E001"]["departamento"]    // "Ventas"
```

## Nil Maps

Un map no inicializado es `nil`. Lectura es segura, escritura produce pánico:

```go
var nilMap map[string]int

// Válido: retorna 0
valor := nilMap["clave"]

// Inválido: panic
// nilMap["clave"] = 10

// Solución: inicializar antes de usar
nilMap = make(map[string]int)
nilMap["clave"] = 10    // Ahora sí es válido
```

## Conjuntos (Sets)

Go no tiene tipo set nativo. Se simula con maps:

```go
// Conjunto de strings
conjunto := make(map[string]bool)

// Agregar
conjunto["rojo"] = true
conjunto["azul"] = true

// Verificar membresía
if conjunto["rojo"] {
    fmt.Println("Contiene rojo")
}

// Tamaño
cantidad := len(conjunto)
```

## Resumen

- `make(map[tipoClave]tipoValor)` para inicializar
- Acceso: `map[clave]` retorna valor y opcionalmente existe
- `delete(map, clave)` para eliminar
- Las claves deben ser tipos comparables
- Lectura de map nil es segura; escritura produce panic
- El orden de iteración sobre maps es no determinístico
