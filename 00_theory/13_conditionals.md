# Condicionales en Go

Los condicionales en Go permiten ejecutar diferentes bloques de código basándose en condiciones específicas. Go ofrece `if`, `else if`, `else` y `switch` para el control de flujo condicional.

## Declaración `if`

### Sintaxis Básica

```go
if condición {
    // código a ejecutar si la condición es verdadera
}
```

**Ejemplo:**
```go
age := 18
if age >= 18 {
    fmt.Println("Eres mayor de edad")
}
```

### `if` con `else`

```go
age := 16
if age >= 18 {
    fmt.Println("Eres mayor de edad")
} else {
    fmt.Println("Eres menor de edad")
}
```

### `if` con `else if`

```go
score := 85

if score >= 90 {
    fmt.Println("Excelente")
} else if score >= 80 {
    fmt.Println("Muy bueno")
} else if score >= 70 {
    fmt.Println("Bueno")
} else if score >= 60 {
    fmt.Println("Aprobado")
} else {
    fmt.Println("Reprobado")
}
```

## `if` con Declaración Inicial

Go permite declarar variables dentro de la condición `if`:

```go
// La variable 'err' solo existe dentro del bloque if
if err := doSomething(); err != nil {
    fmt.Println("Error:", err)
    return
}
// 'err' no está disponible aquí
```

**Ejemplos prácticos:**
```go
// Verificar si existe una clave en un map
if value, exists := myMap["key"]; exists {
    fmt.Println("El valor es:", value)
} else {
    fmt.Println("La clave no existe")
}

// Conversión de tipos segura
if num, ok := value.(int); ok {
    fmt.Println("El número es:", num)
} else {
    fmt.Println("No es un entero")
}
```

## Declaración `switch`

`switch` es una alternativa más limpia a múltiples declaraciones `if-else if`.

### `switch` Básico

```go
day := "lunes"

switch day {
case "lunes":
    fmt.Println("Inicio de semana")
case "martes", "miércoles", "jueves":
    fmt.Println("Día de trabajo")
case "viernes":
    fmt.Println("¡Por fin viernes!")
case "sábado", "domingo":
    fmt.Println("Fin de semana")
default:
    fmt.Println("Día no válido")
}
```

### `switch` sin Expresión (como if-else if)

```go
temperature := 25

switch {
case temperature < 0:
    fmt.Println("Congelándose")
case temperature < 15:
    fmt.Println("Frío")
case temperature < 25:
    fmt.Println("Fresco")
case temperature < 35:
    fmt.Println("Cálido")
default:
    fmt.Println("Caliente")
}
```

### `switch` con Declaración Inicial

```go
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X")
case "linux":
    fmt.Println("Linux")
default:
    fmt.Printf("Otro OS: %s\n", os)
}
```

### `switch` de Tipos

Para verificar el tipo de una interface{}:

```go
func checkType(x interface{}) {
    switch v := x.(type) {
    case int:
        fmt.Printf("Es un entero: %d\n", v)
    case string:
        fmt.Printf("Es una cadena: %s\n", v)
    case bool:
        fmt.Printf("Es un booleano: %t\n", v)
    default:
        fmt.Printf("Tipo desconocido: %T\n", v)
    }
}
```

## Control de Flujo Avanzado

### `fallthrough` en `switch`

Por defecto, los casos en Go no "caen" al siguiente. Usa `fallthrough` explícitamente:

```go
grade := 'B'

switch grade {
case 'A':
    fmt.Println("Excelente")
    fallthrough
case 'B':
    fmt.Println("Muy bueno")
    fallthrough  
case 'C':
    fmt.Println("Aprobado")
default:
    fmt.Println("Estudia más")
}
// Con grade='B', imprime: "Muy bueno" y "Aprobado"
```

### Condicionales Complejas

```go
func evaluateUser(age int, hasLicense bool, hasExperience bool) string {
    // Condiciones complejas con operadores lógicos
    if age >= 18 && hasLicense {
        if age >= 25 || hasExperience {
            return "Puede conducir cualquier vehículo"
        } else {
            return "Puede conducir con restricciones"
        }
    } else if age >= 16 {
        return "Puede obtener licencia de aprendiz"
    } else {
        return "Demasiado joven para conducir"
    }
}
```

## Patrones Comunes

### 1. Validación Temprana (Early Return)

```go
func processUser(user *User) error {
    if user == nil {
        return errors.New("usuario no puede ser nil")
    }
    
    if user.Name == "" {
        return errors.New("nombre requerido")
    }
    
    if user.Age < 0 {
        return errors.New("edad debe ser positiva")
    }
    
    // Procesar usuario...
    return nil
}
```

### 2. Manejo de Errores

```go
func readFile(filename string) ([]byte, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            return nil, fmt.Errorf("archivo no encontrado: %s", filename)
        }
        return nil, fmt.Errorf("error leyendo archivo: %w", err)
    }
    return data, nil
}
```

### 3. Configuración Condicional

```go
func getConfig() Config {
    config := DefaultConfig()
    
    if env := os.Getenv("APP_ENV"); env != "" {
        switch env {
        case "development":
            config.Debug = true
            config.LogLevel = "debug"
        case "production":
            config.Debug = false
            config.LogLevel = "error"
        case "testing":
            config.Debug = true
            config.LogLevel = "info"
        }
    }
    
    return config
}
```

## Ejemplos Prácticos Completos

### 1. Calculadora Simple

```go
func calculate(a, b float64, operation string) (float64, error) {
    switch operation {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        if b == 0 {
            return 0, errors.New("división por cero")
        }
        return a / b, nil
    default:
        return 0, errors.New("operación no válida")
    }
}
```

### 2. Clasificador de Edad

```go
func classifyAge(age int) string {
    switch {
    case age < 0:
        return "Edad inválida"
    case age < 13:
        return "Niño"
    case age < 18:
        return "Adolescente"
    case age < 65:
        return "Adulto"
    default:
        return "Adulto mayor"
    }
}
```

### 3. Validador de Contraseña

```go
func validatePassword(password string) []string {
    var errors []string
    
    if len(password) < 8 {
        errors = append(errors, "Debe tener al menos 8 caracteres")
    }
    
    hasUpper := false
    hasLower := false
    hasDigit := false
    hasSpecial := false
    
    for _, char := range password {
        switch {
        case char >= 'A' && char <= 'Z':
            hasUpper = true
        case char >= 'a' && char <= 'z':
            hasLower = true
        case char >= '0' && char <= '9':
            hasDigit = true
        case strings.ContainsRune("!@#$%^&*", char):
            hasSpecial = true
        }
    }
    
    if !hasUpper {
        errors = append(errors, "Debe tener al menos una mayúscula")
    }
    if !hasLower {
        errors = append(errors, "Debe tener al menos una minúscula")
    }
    if !hasDigit {
        errors = append(errors, "Debe tener al menos un dígito")
    }
    if !hasSpecial {
        errors = append(errors, "Debe tener al menos un carácter especial")
    }
    
    return errors
}
```

## Mejores Prácticas

### 1. Condiciones Claras y Legibles

```go
// ❌ Difícil de leer
if !(user.Age < 18) && user.HasLicense && (user.Experience > 2 || user.Age > 25) {
    // ...
}

// ✅ Más claro
isAdult := user.Age >= 18
hasExperience := user.Experience > 2 || user.Age > 25

if isAdult && user.HasLicense && hasExperience {
    // ...
}
```

### 2. Evita Anidación Excesiva

```go
// ❌ Muy anidado
func processRequest(req *Request) error {
    if req != nil {
        if req.User != nil {
            if req.User.IsActive {
                if req.User.HasPermission("read") {
                    // procesar...
                    return nil
                } else {
                    return errors.New("sin permisos")
                }
            } else {
                return errors.New("usuario inactivo")
            }
        } else {
            return errors.New("usuario requerido")
        }
    } else {
        return errors.New("request no puede ser nil")
    }
}

// ✅ Mejor con early returns
func processRequest(req *Request) error {
    if req == nil {
        return errors.New("request no puede ser nil")
    }
    
    if req.User == nil {
        return errors.New("usuario requerido")
    }
    
    if !req.User.IsActive {
        return errors.New("usuario inactivo")
    }
    
    if !req.User.HasPermission("read") {
        return errors.New("sin permisos")
    }
    
    // procesar...
    return nil
}
```

### 3. Usar `switch` para Múltiples Condiciones

```go
// ❌ Múltiples if-else
if status == "active" || status == "pending" || status == "processing" {
    // manejar estados activos
} else if status == "inactive" || status == "disabled" || status == "suspended" {
    // manejar estados inactivos
} else {
    // manejar estado desconocido
}

// ✅ Más limpio con switch
switch status {
case "active", "pending", "processing":
    // manejar estados activos
case "inactive", "disabled", "suspended":
    // manejar estados inactivos
default:
    // manejar estado desconocido
}
```

Los condicionales son fundamentales para controlar el flujo de ejecución en Go. Usar las estructuras correctas y seguir las mejores prácticas hace que tu código sea más legible, mantenible y menos propenso a errores.
