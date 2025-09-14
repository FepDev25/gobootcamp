package main

import (
	"fmt"
)

/*
EJERCICIOS NIVEL AVANZADO
=========================

Estos ejercicios cubren conceptos avanzados de Go:
- Funciones variádicas y funciones como valores
- Defer, panic y recover
- Init y exit
- Manejo avanzado de errores
- Patrones de diseño básicos

Instrucciones: Implementa cada función según la descripción en los comentarios.
Puedes probar tu código ejecutando: go run advanced_exercises.go
*/

func main() {
	fmt.Println("=== EJERCICIOS NIVEL AVANZADO ===")

	// Descomenta las líneas siguientes cuando hayas implementado las funciones
	// testExercise1()
	// testExercise2()
	// testExercise3()
	// testExercise4()
	// testExercise5()
	// testExercise6()
	// testExercise7()
	// testExercise8()
	// testExercise9()
	// testExercise10()

	fmt.Println("¡Todos los ejercicios completados!")
}

/*
EJERCICIO 1: Funciones Variádicas Avanzadas
==========================================
Implementa una calculadora que acepta múltiples operaciones.

Instrucciones:
1. Crea una función variádica que acepta un operador y múltiples números
2. Soporta operaciones: "sum", "mult", "max", "min", "avg"
3. Implementa otra función variádica que acepta múltiples operadores y números
4. Usa defer para logging de operaciones
5. Maneja errores apropiadamente (división por cero, operadores inválidos)

Ejemplo:
calcular("sum", 1, 2, 3, 4, 5) → 15, nil
calcular("avg", 10, 20, 30) → 20.0, nil
operacionesMultiples("sum", 1, 2, "mult", 3, 4, "max", 5, 6) → [3, 12, 6], nil
*/
func ejercicio1() {
	// TODO: Implementa aquí tu solución
	// Implementa las funciones: calcular(operador string, numeros ...float64) y operacionesMultiples(args ...interface{})

}

/*
EJERCICIO 2: Sistema de Logging con Defer
========================================
Implementa un sistema de logging robusto usando defer.

Instrucciones:
1. Crea funciones que usen defer para:
  - Medir tiempo de ejecución de funciones
  - Limpiar recursos (archivos, conexiones)
  - Logging de entrada y salida de funciones
  - Capturar y registrar panics

2. Implementa una función que simule operaciones costosas
3. Demuestra el uso correcto de múltiples defer en orden

Ejemplo:
[ENTRADA] procesarArchivo(archivo.txt)
[TIEMPO] Iniciando operación...
[RECURSO] Archivo abierto: archivo.txt
[PROCESANDO] Leyendo contenido...
[RECURSO] Archivo cerrado: archivo.txt
[TIEMPO] Operación completada en 150ms
[SALIDA] procesarArchivo completada exitosamente
*/
func ejercicio2() {
	// TODO: Implementa aquí tu solución
	// Crea funciones auxiliares y demuestra el uso de defer

}

/*
EJERCICIO 3: Manejo Robusto de Panic y Recover
=============================================
Implementa un sistema que maneja panics graciosamente.

Instrucciones:
1. Crea funciones que puedan hacer panic en diferentes escenarios
2. Implementa un middleware de recuperación que:
  - Capture cualquier panic
  - Registre el error con stack trace
  - Permita que el programa continúe
  - Retorne un error apropiado

3. Implementa un servidor HTTP simulado que use este sistema
4. Demuestra recovery en goroutines (bonus)

Ejemplo:
operacionRiesgosa1() → panic: "división por cero"
operacionRiesgosa2() → panic: "acceso a slice fuera de límites"
sistemaRecuperacion() → captura panics, registra errores, continúa ejecución
*/
func ejercicio3() {
	// TODO: Implementa aquí tu solución
	// Crea funciones que hagan panic y un sistema de recovery

}

/*
EJERCICIO 4: Sistema de Inicialización con Init
==============================================
Implementa un sistema de configuración usando init.

Instrucciones:
1. Usa múltiples funciones init para:
  - Cargar configuración desde variables de entorno
  - Validar prerrequisitos del sistema
  - Inicializar pools de conexiones
  - Configurar logging

2. Simula carga de configuración desde diferentes fuentes
3. Implementa un sistema de fallbacks para configuración
4. Demuestra orden de ejecución de init

Ejemplo:
init1: Cargando configuración básica...
init2: Validando prerrequisitos...
init3: Inicializando pools de conexión...
init4: Configurando sistema de logging...
main: Sistema iniciado correctamente
*/
func ejercicio4() {
	// TODO: Implementa aquí tu solución
	// Usa múltiples funciones init y variables globales

}

/*
EJERCICIO 5: Sistema de Exit Codes y Cleanup
===========================================
Implementa un programa que maneja exit codes apropiadamente.

Instrucciones:
1. Crea un sistema que:
  - Valide argumentos de línea de comandos
  - Ejecute diferentes operaciones basadas en argumentos
  - Use códigos de salida apropiados (0, 1, 2, etc.)
  - Implemente cleanup antes de salir

2. Maneja señales del sistema (simulado)
3. Implementa timeout para operaciones
4. Demuestra uso apropiado de os.Exit vs return

Ejemplo de uso:
./programa --operation=validate --file=data.txt → exit 0
./programa --operation=invalid → exit 2 (uso incorrecto)
./programa --operation=process --file=missing.txt → exit 1 (error general)
*/
func ejercicio5() {
	// TODO: Implementa aquí tu solución
	// Simula argumentos de línea de comandos y maneja exit codes

}

/*
EJERCICIO 6: Factory Pattern con Funciones
=========================================
Implementa el patrón Factory usando funciones.

Instrucciones:
1. Define interfaces para diferentes tipos de procesadores de datos
2. Implementa factory functions que retornen diferentes implementaciones
3. Usa funciones como valores para configurar comportamiento
4. Implementa un sistema de plugins usando funciones
5. Demuestra polimorfismo con funciones

Ejemplo:
procesadorJSON := crearProcesador("json", configuraciones...)
procesadorXML := crearProcesador("xml", configuraciones...)
procesadorCSV := crearProcesador("csv", configuraciones...)

datos := "datos de ejemplo"
procesadorJSON.procesar(datos) → output JSON
procesadorXML.procesar(datos) → output XML
*/
func ejercicio6() {
	// TODO: Implementa aquí tu solución
	// Define interfaces y implementa factory pattern

}

/*
EJERCICIO 7: Pipeline de Procesamiento con Funciones
===================================================
Implementa un pipeline de procesamiento usando funciones como valores.

Instrucciones:
1. Define un tipo de función para etapas del pipeline
2. Implementa diferentes etapas de procesamiento:
  - Validación de datos
  - Transformación
  - Filtrado
  - Agregación
  - Formateo de salida

3. Combina etapas en un pipeline configurable
4. Implementa manejo de errores en cada etapa
5. Permite pipeline asíncrono (bonus)

Ejemplo:

	pipeline := []EtapaProcesamiento{
	    validarEntrada,
	    transformarDatos,
	    filtrarValidos,
	    agregarEstadisticas,
	    formatearSalida,
	}

resultado := ejecutarPipeline(datos, pipeline...)
*/
func ejercicio7() {
	// TODO: Implementa aquí tu solución
	// Define tipos de función y implementa pipeline

}

/*
EJERCICIO 8: Sistema de Cache con TTL
===================================
Implementa un sistema de cache con tiempo de vida usando defer y goroutines.

Instrucciones:
1. Implementa un cache que:
  - Almacene datos con tiempo de expiración (TTL)
  - Use defer para cleanup automático
  - Implemente diferentes políticas de expulsión
  - Maneje concurrencia básica (simulada)

2. Implementa funciones de utilidad para el cache
3. Demuestra uso con diferentes tipos de datos
4. Implementa estadísticas del cache

Ejemplo:
cache := nuevoCache(ttl: 5segundos, maxSize: 100)
cache.set("usuario:123", datos, 10segundos)
datos, encontrado := cache.get("usuario:123")
stats := cache.estadisticas() → hits: 45, misses: 12, size: 67
*/
func ejercicio8() {
	// TODO: Implementa aquí tu solución
	// Implementa estructuras de cache y métodos

}

/*
EJERCICIO 9: Validador de Datos Configurable
===========================================
Implementa un sistema de validación usando funciones variádicas.

Instrucciones:
1. Define un tipo de función para validadores
2. Implementa validadores básicos:
  - Longitud mínima/máxima
  - Formato email
  - Rango numérico
  - Patrones regex
  - Validaciones personalizadas

3. Combina validadores usando funciones variádicas
4. Implementa validación de structs completos
5. Retorna errores detallados

Ejemplo:
validarEmail := validador.combinar(

	noVacio(),
	longitudMaxima(100),
	formatoEmail(),
	dominiosPermitidos("gmail.com", "hotmail.com"),

)
errores := validarEmail("test@gmail.com") → nil
errores := validarEmail("invalid") → ["formato inválido", "dominio no permitido"]
*/
func ejercicio9() {
	// TODO: Implementa aquí tu solución
	// Define tipos de función y implementa validadores

}

/*
EJERCICIO 10: Framework de Testing Personalizado
===============================================
Implementa un mini framework de testing usando defer, panic y recover.

Instrucciones:
1. Implementa funciones para:
  - Definir casos de prueba
  - Ejecutar suites de tests
  - Capturar panics en tests
  - Generar reportes detallados
  - Benchmarking básico

2. Usa defer para cleanup después de cada test
3. Implementa assertions personalizadas
4. Permite setup y teardown por test
5. Genera estadísticas de ejecución

Ejemplo:
suite := nuevaSuiteTest("MiModulo")

	suite.agregar("TestSuma", func(t *Test) {
	    resultado := suma(2, 3)
	    t.assertEqual(resultado, 5)
	})

	suite.agregar("TestDivision", func(t *Test) {
	    defer t.cleanup(func() {  })
	    resultado := division(10, 0)
	    t.expectPanic("división por cero")
	})

reporte := suite.ejecutar() → Tests: 25, Pasados: 23, Fallidos: 2, Tiempo: 1.5s
*/
func ejercicio10() {
	// TODO: Implementa aquí tu solución
	// Implementa framework de testing básico

}

// Funciones de prueba - NO MODIFICAR
func testExercise1() {
	fmt.Println("\n--- EJERCICIO 1 ---")
	ejercicio1()
}

func testExercise2() {
	fmt.Println("\n--- EJERCICIO 2 ---")
	ejercicio2()
}

func testExercise3() {
	fmt.Println("\n--- EJERCICIO 3 ---")
	ejercicio3()
}

func testExercise4() {
	fmt.Println("\n--- EJERCICIO 4 ---")
	ejercicio4()
}

func testExercise5() {
	fmt.Println("\n--- EJERCICIO 5 ---")
	ejercicio5()
}

func testExercise6() {
	fmt.Println("\n--- EJERCICIO 6 ---")
	ejercicio6()
}

func testExercise7() {
	fmt.Println("\n--- EJERCICIO 7 ---")
	ejercicio7()
}

func testExercise8() {
	fmt.Println("\n--- EJERCICIO 8 ---")
	ejercicio8()
}

func testExercise9() {
	fmt.Println("\n--- EJERCICIO 9 ---")
	ejercicio9()
}

func testExercise10() {
	fmt.Println("\n--- EJERCICIO 10 ---")
	ejercicio10()
}
