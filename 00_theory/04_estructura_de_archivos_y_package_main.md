# Estructura de Archivos y package main en Go

## ¿Por qué obtengo errores de "main redeclared" en Go?

Si estás viendo un error como:

```
main redeclared in this block
```

O te preguntas por qué no puedes ejecutar `go run` en un archivo después de cambiar el nombre del paquete, no estás solo. Muchos estudiantes encuentran esto mientras siguen el curso, así que vamos a explicarlo paso a paso y aclarar cómo Go maneja archivos, paquetes y la función main.

## La Configuración en Este Curso

En este curso, construimos el entendimiento archivo por archivo: primero hello_world.go, luego data_types.go, variables.go, y así sucesivamente.

Pero solo un archivo existe en la carpeta main a la vez.

Una vez que un concepto se cubre en una lección, ese archivo es:

- Renombrado (por ejemplo, a .txt), para que Go no trate de compilarlo.
- Movido a una subcarpeta como basics/, y su línea de paquete se actualiza para coincidir con la carpeta (package basics).

Esto se hace intencionalmente para mantener la experiencia de aprendizaje limpia y prevenir errores como múltiples declaraciones de main().

## ¿Por qué sucede este error?

Digamos que tienes dos archivos en la misma carpeta:

**random.go**
```go
package main

func main() {
    fmt.Println("Código aleatorio")
}
```

**variables.go**
```go
package main

func main() {
    fmt.Println("Código de variables")
}
```

Obtendrás un error:

```
main redeclared in this block
```

Esto es porque Go permite solo UNA función main() por paquete, y por defecto, todos los archivos .go en una carpeta pertenecen al mismo paquete (a menos que se especifique explícitamente lo contrario).

## ¿Qué es package main?

En Go, si estás escribiendo un programa ejecutable, debe:

- Usar `package main`
- Tener una función `main()` como punto de entrada

Si tratas de ejecutar un archivo sin `package main` o sin una función `main()`, obtendrás un error como:

```
go run: cannot run non-main package
```

## ¿Entonces qué es package basics, package utils, etc.?

Estos son paquetes de estilo biblioteca, usados para código reutilizable. No necesitan una función `main()`, y no están destinados a ser ejecutados directamente con `go run`.

Por eso, en algunas lecciones, cambio el archivo de:

```go
package main
```

a:

```go
package basics
```

Antes de archivarlo o moverlo: esto le dice a Go "Oye, esto ya no es un programa para ejecutar".

## Tus Opciones

Aquí te explico cómo puedes evitar o arreglar estos errores:

### Opción 1: Mantener Solo Un main() a la Vez

Si quieres mantener todo en una carpeta (por ejemplo, durante el aprendizaje temprano), solo asegúrate de que un archivo tenga una función `main()`.

Puedes renombrar temporalmente otros archivos .go a .txt o comentar funciones `main()` adicionales.

### Opción 2: Usar Subcarpetas para Diferentes Programas

Estructura tu código así:

```
/Go_course/
  /random/
    random.go   --> package main (tiene su propio main())
  /variables/
    variables.go --> package main (tiene su propio main())
```

Ahora puedes hacer `cd` en cada carpeta y ejecutar:

```bash
go run random.go
```

Cada archivo es parte de un programa Go diferente ahora. ¡Sin conflictos!

### Opción 3: Combinar Lógica en Un Archivo

En lugar de tener dos funciones `main()`, divide la funcionalidad en funciones nombradas:

```go
package main

func main() {
    saludar()
    mostrarVariables()
}

func saludar() {
    fmt.Println("Hola desde Random")
}

func mostrarVariables() {
    fmt.Println("Hola desde Variables")
}
```

Esto es bueno para practicar y ayuda a organizar mejor tu código.

## ¿Por qué renombro archivos o cambio el paquete?

Cuando me veas renombrando archivos o cambiando `package main` a `package basics`, es para evitar errores de compilación. De esta manera, puedo mantener el código de la lección anterior como referencia sin que Go piense que estoy tratando de ejecutar múltiples programas a la vez.

## Reglas Importantes para Recordar

### Permitido:
- Un `main()` en un paquete
- Archivos en diferentes carpetas, cada uno con `package main`
- `go run file.go` con `package main`
- Archivar archivos .go antiguos cambiando el paquete o la extensión

### No Permitido:
- Múltiples funciones `main()` en el mismo paquete
- Dos archivos main() en la misma carpeta
- `go run` en un archivo con un nombre de paquete diferente
- Dejar muchos archivos .go con `main()` en una carpeta

## Preguntas Frecuentes

**P: ¿Puedo tener múltiples archivos .go con main()?**
R: Sí, pero solo si están en carpetas o paquetes separados.

**P: ¿Por qué cambiar el nombre del paquete a basics?**
R: Para indicar que el archivo ahora es parte de una biblioteca, no destinado a ser ejecutado.

**P: ¿Por qué go run da un error después de cambiar el nombre del paquete?**
R: Porque solo `package main` puede ser ejecutado como un ejecutable.

**P: ¿Por qué VS Code muestra líneas onduladas cuando tengo múltiples archivos?**
R: Porque detecta múltiples funciones `main()` o declaraciones duplicadas.

## ¿Por qué solo cambio el nombre del paquete (no la función main())?

Podrías haber notado que dentro de las carpetas basics, intermediate y advanced, muchos archivos todavía tienen una función `main()`, y eso está completamente bien. No renombro la función `main()` en esos archivos porque no estamos tratando de ejecutarlos directamente mientras están en esas carpetas. Cada una de esas carpetas es como un "área de almacenamiento" para ejemplos de código que ya hemos cubierto.

En lugar de gastar tiempo eliminando o renombrando cada función `main()`, simplemente cambio el nombre del paquete para que coincida con la carpeta (package basics, package intermediate, etc.). De esa manera, los archivos todavía se guardan para referencia, pero Go ya no los tratará como programas ejecutables, los tratará como código de biblioteca regular que no está destinado a ser ejecutado directamente.

Más tarde, si quieres ejecutar cualquiera de esos archivos de ejemplo otra vez, solo mueve el archivo a la carpeta principal, renombra su paquete a main, y asegúrate de que no haya otro archivo con una función `main()` en esa carpeta. Entonces puedes ejecutarlo normalmente usando:

```bash
go run tu_archivo.go
```

Esta configuración nos ayuda a mantener las cosas limpias, enfocadas y organizadas por tema, sin quedarnos atascados en las restricciones de Go alrededor de las funciones `main()`.

A medida que progresemos en el curso y te sientas más cómodo, aprenderás cómo estructurar proyectos de manera más profesional, pero ahora mismo, el objetivo es entender los fundamentos de Go, no la arquitectura a gran escala.
