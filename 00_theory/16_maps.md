# Maps en Go

Los maps (mapas) en Go son estructuras de datos que almacenan pares clave-valor, similares a hash tables, diccionarios o arrays asociativos en otros lenguajes. Proporcionan una forma eficiente de recuperar valores basándose en una clave única.

## Concepto de Map

Un map asocia claves con valores:
- **Clave**: Debe ser un tipo comparable (int, string, bool, array, struct con campos comparables)
- **Valor**: Puede ser de cualquier tipo
- **Acceso**: O(1) en promedio para operaciones básicas

```
myMap["Felipe"] → 25
myMap["María"]  → 30  
myMap["Juan"]   → 22
```

## Declaración e Inicialización

### Crear Map Vacío

```go
// Usando make
ages := make(map[string]int)
fmt.Println(ages) // map[]

// Declaración vacía (nil map)
var scores map[string]int
fmt.Println(scores == nil) // true
// scores["test"] = 100    // ❌ panic: asignación a nil map
```

### Map Literal

```go
// Inicializar con valores
ages := map[string]int{
    "Felipe": 25,
    "María":  30,
    "Juan":   22,
}
fmt.Println(ages) // map[Felipe:25 Juan:22 María:30]

// Map vacío pero no nil
empty := map[string]int{}
fmt.Println(empty == nil) // false
```

### Diferentes Tipos de Maps

```go
// Map de strings a enteros
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}

// Map de enteros a strings
names := map[int]string{
    1: "Uno",
    2: "Dos",
    3: "Tres",
}

// Map de strings a slices
groups := map[string][]string{
    "developers": {"Alice", "Bob"},
    "designers":  {"Carol", "Dave"},
}
```

## Operaciones Básicas

### Agregar y Modificar Elementos

```go
ages := make(map[string]int)

// Agregar elementos
ages["Felipe"] = 25
ages["María"] = 30
ages["Juan"] = 22

fmt.Println(ages) // map[Felipe:25 Juan:22 María:30]

// Modificar elemento existente
ages["Felipe"] = 26
fmt.Println("Edad de Felipe:", ages["Felipe"]) // 26
```

### Acceder a Elementos

```go
ages := map[string]int{
    "Felipe": 25,
    "María":  30,
}

// Acceso directo
fmt.Println("Edad de Felipe:", ages["Felipe"]) // 25

// Verificar si existe (idioma "comma ok")
if age, exists := ages["Pedro"]; exists {
    fmt.Println("Edad de Pedro:", age)
} else {
    fmt.Println("Pedro no encontrado")
}

// Valor cero si no existe
fmt.Println("Edad de Pedro:", ages["Pedro"]) // 0 (valor cero de int)
```

### Eliminar Elementos

```go
ages := map[string]int{
    "Felipe": 25,
    "María":  30,
    "Juan":   22,
}

// Eliminar un elemento
delete(ages, "Juan")
fmt.Println("Después de eliminar Juan:", ages)

// Eliminar clave que no existe (no hace nada)
delete(ages, "Pedro") // No causa error
```

### Limpiar Map Completo

```go
ages := map[string]int{
    "Felipe": 25,
    "María":  30,
}

// Go 1.21+
clear(ages)
fmt.Println("Después de clear:", ages) // map[]

// Alternativa para versiones anteriores
ages = make(map[string]int) // Crear nuevo map vacío
```

## Iteración sobre Maps

```go
teams := map[string]int{
    "Chelsea FC":  2,
    "Real Madrid": 15,
    "Barcelona":   5,
}

// Iterar sobre clave y valor
for team, titles := range teams {
    fmt.Printf("%s: %d títulos\n", team, titles)
}

// Solo las claves
for team := range teams {
    fmt.Println("Equipo:", team)
}

// Solo los valores (raro)
for _, titles := range teams {
    fmt.Println("Títulos:", titles)
}
```

**Nota importante**: El orden de iteración en maps es **aleatorio** en Go.

## Verificar Existencia

### Patrón "Comma OK"

```go
ages := map[string]int{
    "Felipe": 25,
    "María":  30,
}

// Verificar si una clave existe
if age, exists := ages["Felipe"]; exists {
    fmt.Printf("Felipe tiene %d años\n", age)
} else {
    fmt.Println("Felipe no encontrado")
}

// Función helper
func checkAge(ages map[string]int, name string) {
    if age, exists := ages[name]; exists {
        fmt.Printf("%s tiene %d años\n", name, age)
    } else {
        fmt.Printf("%s no encontrado\n", name)
    }
}
```

## Longitud de Maps

```go
ages := map[string]int{
    "Felipe": 25,
    "María":  30,
    "Juan":   22,
}

fmt.Println("Número de personas:", len(ages)) // 3
```

## Maps Anidados

```go
// Map de maps
users := map[string]map[string]interface{}{
    "felipe": {
        "age":    25,
        "city":   "Madrid",
        "active": true,
    },
    "maria": {
        "age":    30,
        "city":   "Barcelona",
        "active": false,
    },
}

// Acceder a valores anidados
fmt.Println("Ciudad de Felipe:", users["felipe"]["city"])

// Modificar valores anidados
users["felipe"]["age"] = 26
```

## Paquete `maps` (Go 1.21+)

```go
import (
    "fmt"
    "maps"
)

original := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}

// Clonar map
clone := maps.Clone(original)
clone["d"] = 4
fmt.Println("Original:", original) // map[a:1 b:2 c:3]
fmt.Println("Clone:", clone)       // map[a:1 b:2 c:3 d:4]

// Verificar igualdad
fmt.Println("Son iguales:", maps.Equal(original, clone)) // false

// Copiar de un map a otro
maps.Copy(clone, original) // Copia elementos de original a clone
```

## Comparación de Maps

Los maps **NO son comparables** directamente:

```go
map1 := map[string]int{"a": 1, "b": 2}
map2 := map[string]int{"a": 1, "b": 2}

// ❌ Esto NO compila
// fmt.Println(map1 == map2) // Error

// ✅ Solo comparable con nil
fmt.Println(map1 == nil) // false

// ✅ Usar maps.Equal (Go 1.21+)
fmt.Println(maps.Equal(map1, map2)) // true
```

### Función de Comparación Manual

```go
func equalMaps(a, b map[string]int) bool {
    if len(a) != len(b) {
        return false
    }
    
    for key, valueA := range a {
        if valueB, exists := b[key]; !exists || valueA != valueB {
            return false
        }
    }
    
    return true
}
```

## Maps como Sets

Usar maps para simular conjuntos (sets):

```go
// Set de strings
set := make(map[string]bool)

// Agregar elementos
set["apple"] = true
set["banana"] = true
set["orange"] = true

// Verificar si existe
if set["apple"] {
    fmt.Println("Apple está en el set")
}

// Iterar sobre el set
fmt.Println("Elementos en el set:")
for item := range set {
    fmt.Println("-", item)
}

// Alternativa más eficiente en memoria
type void struct{}
setOptimized := make(map[string]void)
setOptimized["apple"] = void{}
```

## Ejemplos Prácticos

### 1. Contador de Frecuencias

```go
func countWords(text string) map[string]int {
    words := strings.Fields(text)
    counts := make(map[string]int)
    
    for _, word := range words {
        counts[word]++
    }
    
    return counts
}

text := "go is great go is awesome go is fun"
wordCounts := countWords(text)
fmt.Println("Frecuencia de palabras:", wordCounts)
// map[awesome:1 fun:1 go:3 great:1 is:3]
```

### 2. Agrupar Elementos

```go
type Person struct {
    Name string
    Age  int
    City string
}

func groupByCity(people []Person) map[string][]Person {
    groups := make(map[string][]Person)
    
    for _, person := range people {
        groups[person.City] = append(groups[person.City], person)
    }
    
    return groups
}

people := []Person{
    {"Alice", 25, "Madrid"},
    {"Bob", 30, "Barcelona"},
    {"Carol", 35, "Madrid"},
}

grouped := groupByCity(people)
for city, residents := range grouped {
    fmt.Printf("Residentes en %s:\n", city)
    for _, person := range residents {
        fmt.Printf("  - %s (%d años)\n", person.Name, person.Age)
    }
}
```

### 3. Cache Simple

```go
type Cache struct {
    data map[string]interface{}
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]interface{}),
    }
}

func (c *Cache) Set(key string, value interface{}) {
    c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
    value, exists := c.data[key]
    return value, exists
}

func (c *Cache) Delete(key string) {
    delete(c.data, key)
}

// Uso
cache := NewCache()
cache.Set("user:123", "Alice")

if value, exists := cache.Get("user:123"); exists {
    fmt.Println("Usuario encontrado:", value)
}
```

### 4. Índice de Búsqueda

```go
type Product struct {
    ID    int
    Name  string
    Price float64
}

type ProductIndex struct {
    byID   map[int]*Product
    byName map[string]*Product
}

func NewProductIndex() *ProductIndex {
    return &ProductIndex{
        byID:   make(map[int]*Product),
        byName: make(map[string]*Product),
    }
}

func (pi *ProductIndex) Add(product *Product) {
    pi.byID[product.ID] = product
    pi.byName[product.Name] = product
}

func (pi *ProductIndex) FindByID(id int) *Product {
    return pi.byID[id]
}

func (pi *ProductIndex) FindByName(name string) *Product {
    return pi.byName[name]
}
```

## Consideraciones de Rendimiento

### 1. Claves Apropiadas

```go
// ✅ Buenas claves (rápidas de hash)
var intMap = make(map[int]string)
var stringMap = make(map[string]int)

// ❌ Claves potencialmente lentas
var sliceMap = make(map[[]int]string) // Error: slice no es comparable
```

### 2. Pre-asignación

```go
// Si conoces el tamaño aproximado
expectedSize := 1000
data := make(map[string]int, expectedSize) // Go 1.21+
```

### 3. Evitar Conversiones Costosas

```go
// ❌ Conversión en cada acceso
func lookup(data map[string]int, id int) int {
    return data[strconv.Itoa(id)] // Conversión costosa
}

// ✅ Usar el tipo correcto desde el principio
func lookupFast(data map[int]int, id int) int {
    return data[id]
}
```

## Limitaciones y Consideraciones

1. **No thread-safe**: Necesitas sincronización para acceso concurrente
2. **Orden aleatorio**: No mantienes orden de inserción
3. **Claves comparables**: Solo tipos comparables pueden ser claves
4. **Referencias**: Maps son tipos de referencia

```go
// Maps son referencias
original := map[string]int{"a": 1}
copy := original  // Misma referencia
copy["b"] = 2
fmt.Println(original) // map[a:1 b:2] - se modificó también
```

Los maps son una herramienta fundamental en Go para crear asociaciones eficientes entre claves y valores. Su flexibilidad y eficiencia los hacen ideales para muchos casos de uso, desde simples lookups hasta estructuras de datos complejas.
