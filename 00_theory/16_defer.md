# Defer

La palabra clave `defer` pospone la ejecución de una función hasta que la función que la contiene termine, ejecutándose justo antes del retorno. Es fundamental para gestión de recursos y limpieza.

## Uso Básico

```go
func ejemplo() {
    defer fmt.Println("Primero en declarar")
    fmt.Println("Segundo en ejecutar")
}
// Salida:
// Segundo en ejecutar
// Primero en declarar
```

## Patrón de Liberación de Recursos

El uso más común es cerrar recursos después de abrirlos:

```go
func leerArchivo(nombre string) error {
    archivo, err := os.Open(nombre)
    if err != nil {
        return err
    }
    defer archivo.Close()
    
    // Procesar archivo...
    // Close() se ejecutará automáticamente al retornar
    return nil
}
```

## Evaluación de Argumentos

Los argumentos de `defer` se evalúan inmediatamente, pero la función se ejecuta al final:

```go
func evaluacion() {
    i := 0
    defer fmt.Println(i)    // Imprime 0, valor actual
    i++
    fmt.Println(i)          // Imprime 1
}
```

## Múltiples Defers

Los defers se apilan (LIFO: Last In, First Out):

```go
func apilado() {
    defer fmt.Println("Primero")
    defer fmt.Println("Segundo")
    defer fmt.Println("Tercero")
    fmt.Println("Ejecutando")
}
// Salida:
// Ejecutando
// Tercero
// Segundo
// Primero
```

## Defers en Bucles

Dentro de bucles, los defers se acumulan y ejecutan al finalizar la función, no la iteración:

```go
func bucle() {
    for i := 0; i < 3; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("Fin del bucle")
}
// Salida:
// Fin del bucle
// 2
// 1
// 0
```

Para ejecutar al final de cada iteración, envuelva en función anónima:

```go
func bucleCorregido() {
    for i := 0; i < 3; i++ {
        func() {
            defer fmt.Println(i)
        }()
    }
}
```

## Defers y Valores de Retorno

Los defers pueden modificar valores de retorno nombrados:

```go
func contar() (resultado int) {
    defer func() {
        resultado++    // Incrementa antes de retornar
    }()
    return 0
}

fmt.Println(contar())    // 1
```

## Funciones Anónimas con Defer

Para ejecutar código complejo pospuesto:

```go
func transaccion() error {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado:", r)
        }
    }()
    
    // Operación que podría fallar
    return nil
}
```

## Rendimiento

Los defers tienen un costo mínimo. En Go 1.14+, el overhead es despreciable para la mayoría de casos de uso. No evite `defer` por razones de rendimiento salvo en código crítico medido.

## Cuándo Usar Defer

- Cerrar archivos, conexiones de red, bases de datos
- Liberar locks (mutex)
- Registrar tiempo de ejecución (logging)
- Recuperación de pánico (con función anónima)

## Cuándo Evitar Defer

- Cuando necesite control exacto del momento de ejecución
- En bucles con muchas iteraciones donde el acumulador de defers consume memoria

## Resumen

- `defer` ejecuta la función al finalizar la función contenedora
- Los argumentos se evalúan inmediatamente; la ejecución se pospone
- Múltiples defers se ejecutan en orden inverso (LIFO)
- Los defers en bucles se acumulan hasta el final de la función
- Pueden modificar valores de retorno nombrados
- Ideal para liberar recursos de forma limpia y segura
