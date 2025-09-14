package main

import (
	"fmt"
	"os"
)

/*
PROYECTO FINAL: SISTEMA DE GESTIÓN DE BIBLIOTECA
===============================================

Este proyecto integra TODOS los conceptos aprendidos en el bootcamp de Go:
- Variables, tipos de datos y constantes
- Operadores y control de flujo
- Arrays, slices y maps
- Funciones (básicas, múltiples retornos, variádicas)
- Defer, panic, recover
- Init y exit
- Manejo de errores

DESCRIPCIÓN DEL PROYECTO:
========================
Desarrollarás un sistema completo de gestión de biblioteca que incluye:

1. 📚 GESTIÓN DE LIBROS
   - Agregar, buscar, actualizar y eliminar libros
   - Categorización por género, autor, año
   - Sistema de ISBN y validaciones

2. 👥 GESTIÓN DE USUARIOS
   - Registro de miembros de la biblioteca
   - Diferentes tipos de membresía (estudiante, profesor, público)
   - Validación de datos de usuario

3. 📖 SISTEMA DE PRÉSTAMOS
   - Préstamo y devolución de libros
   - Control de fechas de vencimiento
   - Multas por retraso
   - Historial de préstamos

4. 📊 REPORTES Y ESTADÍSTICAS
   - Libros más populares
   - Usuarios más activos
   - Estadísticas de la biblioteca
   - Reportes de multas

5. 💾 PERSISTENCIA DE DATOS
   - Guardar y cargar datos (simulado con archivos de texto)
   - Backup y restauración
   - Validación de integridad de datos

REQUISITOS TÉCNICOS:
==================
✅ Usar todos los tipos de datos aprendidos
✅ Implementar al menos 20 funciones diferentes
✅ Usar funciones variádicas donde sea apropiado
✅ Implementar manejo de errores robusto
✅ Usar defer para cleanup de recursos
✅ Implementar recovery para operaciones críticas
✅ Usar init para configuración inicial
✅ Manejar exit codes apropiadamente
✅ Validación exhaustiva de datos de entrada
✅ Comentarios detallados en todo el código

ESTRUCTURA SUGERIDA:
==================
1. Definir structs para Libro, Usuario, Prestamo
2. Implementar funciones de inicialización
3. Crear funciones CRUD para cada entidad
4. Implementar lógica de negocio (préstamos, multas)
5. Crear sistema de reportes
6. Implementar persistencia básica
7. Crear interfaz de línea de comandos
8. Manejar errores y casos extremos

¡Este es tu momento de brillar! Demuestra todo lo que has aprendido.
*/

// CONSTANTES DEL SISTEMA
const (
	// Tipos de membresía
	ESTUDIANTE = "estudiante"
	PROFESOR   = "profesor"
	PUBLICO    = "publico"

	// Límites del sistema
	MAX_LIBROS_PRESTAMO_ESTUDIANTE = 3
	MAX_LIBROS_PRESTAMO_PROFESOR   = 10
	MAX_LIBROS_PRESTAMO_PUBLICO    = 5

	// Días de préstamo por tipo de usuario
	DIAS_PRESTAMO_ESTUDIANTE = 14
	DIAS_PRESTAMO_PROFESOR   = 30
	DIAS_PRESTAMO_PUBLICO    = 7

	// Multas (por día de retraso)
	MULTA_ESTUDIANTE = 0.50
	MULTA_PROFESOR   = 1.00
	MULTA_PUBLICO    = 0.75

	// Códigos de salida
	EXIT_SUCCESS = 0
	EXIT_ERROR   = 1
	EXIT_USAGE   = 2
)

// ESTRUCTURAS DE DATOS

/*
TODO: Define la estructura Libro
Campos sugeridos:
- ISBN (string)
- Titulo (string)
- Autor (string)
- Genero (string)
- AnoPublicacion (int)
- Disponible (bool)
- FechaAdquisicion (string)
*/
type Libro struct {
	// TODO: Implementa aquí la estructura
}

/*
TODO: Define la estructura Usuario
Campos sugeridos:
- ID (int)
- Nombre (string)
- Email (string)
- TipoMembresia (string)
- FechaRegistro (string)
- LibrosPrestados ([]string) // ISBNs
- MultasPendientes (float64)
*/
type Usuario struct {
	// TODO: Implementa aquí la estructura
}

/*
TODO: Define la estructura Prestamo
Campos sugeridos:
- ID (int)
- UsuarioID (int)
- ISBN (string)
- FechaPrestamo (string)
- FechaVencimiento (string)
- FechaDevolucion (string) // vacío si no se ha devuelto
- MultaAplicada (float64)
*/
type Prestamo struct {
	// TODO: Implementa aquí la estructura
}

/*
TODO: Define la estructura Biblioteca
Campos sugeridos:
- Libros (map[string]Libro) // ISBN -> Libro
- Usuarios (map[int]Usuario) // ID -> Usuario
- Prestamos ([]Prestamo)
- ProximoUsuarioID (int)
- ProximoPrestamoID (int)
*/
type Biblioteca struct {
	// TODO: Implementa aquí la estructura
}

// VARIABLES GLOBALES
var (
	biblioteca *Biblioteca
	// TODO: Agrega otras variables globales necesarias
)

// FUNCIONES DE INICIALIZACIÓN

/*
TODO: Implementa función init
Debe:
1. Inicializar la biblioteca
2. Cargar datos desde archivos (si existen)
3. Configurar el sistema de logging
4. Validar prerrequisitos
*/
func init() {
	// TODO: Implementa aquí
}

/*
TODO: Implementa nuevaBiblioteca
Retorna una instancia inicializada de Biblioteca
*/
func nuevaBiblioteca() *Biblioteca {
	// TODO: Implementa aquí
	return nil
}

// FUNCIONES DE GESTIÓN DE LIBROS

/*
TODO: Implementa agregarLibro
Parámetros: isbn, titulo, autor, genero string, ano int
Retorna: error
Debe:
1. Validar que el ISBN no exista
2. Validar formato de ISBN
3. Validar que todos los campos estén completos
4. Agregar el libro a la biblioteca
*/
func agregarLibro(isbn, titulo, autor, genero string, ano int) error {
	// TODO: Implementa aquí
	return nil
}

/*
TODO: Implementa buscarLibros
Parámetros variádicos para diferentes criterios de búsqueda
Retorna: []Libro, error
Debe permitir buscar por: titulo, autor, genero, año
*/
func buscarLibros(criterios ...interface{}) ([]Libro, error) {
	// TODO: Implementa aquí
	return nil, nil
}

/*
TODO: Implementa actualizarLibro
Parámetros: isbn string, campos map[string]interface{}
Retorna: error
Debe permitir actualizar campos específicos del libro
*/
func actualizarLibro(isbn string, campos map[string]interface{}) error {
	// TODO: Implementa aquí
	return nil
}

/*
TODO: Implementa eliminarLibro
Parámetros: isbn string
Retorna: error
Debe validar que el libro no esté prestado antes de eliminar
*/
func eliminarLibro(isbn string) error {
	// TODO: Implementa aquí
	return nil
}

// FUNCIONES DE GESTIÓN DE USUARIOS

/*
TODO: Implementa registrarUsuario
Parámetros: nombre, email, tipoMembresia string
Retorna: int (ID del usuario), error
Debe:
1. Validar formato de email
2. Validar tipo de membresía
3. Generar ID único
4. Agregar usuario al sistema
*/
func registrarUsuario(nombre, email, tipoMembresia string) (int, error) {
	// TODO: Implementa aquí
	return 0, nil
}

/*
TODO: Implementa buscarUsuario
Parámetros: criterio, valor string (ejemplo: "id", "123" o "email", "user@example.com")
Retorna: Usuario, bool (encontrado), error
*/
func buscarUsuario(criterio, valor string) (Usuario, bool, error) {
	// TODO: Implementa aquí
	return Usuario{}, false, nil
}

/*
TODO: Implementa actualizarUsuario
Parámetros: id int, campos map[string]interface{}
Retorna: error
*/
func actualizarUsuario(id int, campos map[string]interface{}) error {
	// TODO: Implementa aquí
	return nil
}

// FUNCIONES DE GESTIÓN DE PRÉSTAMOS

/*
TODO: Implementa prestarLibro
Parámetros: usuarioID int, isbn string
Retorna: int (ID del préstamo), error
Debe:
1. Validar que el usuario existe
2. Validar que el libro existe y está disponible
3. Verificar límites de préstamo del usuario
4. Calcular fecha de vencimiento
5. Crear registro de préstamo
*/
func prestarLibro(usuarioID int, isbn string) (int, error) {
	// TODO: Implementa aquí
	return 0, nil
}

/*
TODO: Implementa devolverLibro
Parámetros: prestamoID int
Retorna: float64 (multa aplicada), error
Debe:
1. Buscar el préstamo
2. Calcular multa si hay retraso
3. Marcar libro como disponible
4. Actualizar registro de préstamo
*/
func devolverLibro(prestamoID int) (float64, error) {
	// TODO: Implementa aquí
	return 0, nil
}

/*
TODO: Implementa renovarPrestamo
Parámetros: prestamoID int
Retorna: string (nueva fecha de vencimiento), error
Debe validar que el préstamo puede renovarse
*/
func renovarPrestamo(prestamoID int) (string, error) {
	// TODO: Implementa aquí
	return "", nil
}

/*
TODO: Implementa obtenerPrestamosVencidos
Retorna: []Prestamo, error
Lista todos los préstamos que han vencido
*/
func obtenerPrestamosVencidos() ([]Prestamo, error) {
	// TODO: Implementa aquí
	return nil, nil
}

// FUNCIONES DE REPORTES Y ESTADÍSTICAS

/*
TODO: Implementa generarReporteGeneral
Retorna: map[string]interface{}, error
Debe incluir:
- Total de libros en la biblioteca
- Total de usuarios registrados
- Libros prestados actualmente
- Multas pendientes totales
- Estadísticas por tipo de usuario
*/
func generarReporteGeneral() (map[string]interface{}, error) {
	// TODO: Implementa aquí
	return nil, nil
}

/*
TODO: Implementa obtenerLibrosPopulares
Parámetros: limite int
Retorna: []map[string]interface{}, error
Lista los libros más prestados
*/
func obtenerLibrosPopulares(limite int) ([]map[string]interface{}, error) {
	// TODO: Implementa aquí
	return nil, nil
}

/*
TODO: Implementa obtenerUsuariosActivos
Parámetros: limite int
Retorna: []map[string]interface{}, error
Lista los usuarios con más préstamos
*/
func obtenerUsuariosActivos(limite int) ([]map[string]interface{}, error) {
	// TODO: Implementa aquí
	return nil, nil
}

// FUNCIONES DE PERSISTENCIA

/*
TODO: Implementa guardarDatos
Parámetros: archivo string
Retorna: error
Debe usar defer para cerrar archivos
Guarda el estado completo de la biblioteca
*/
func guardarDatos(archivo string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error crítico guardando datos: %v\n", r)
		}
	}()

	// TODO: Implementa aquí
	return nil
}

/*
TODO: Implementa cargarDatos
Parámetros: archivo string
Retorna: error
Debe usar defer para cerrar archivos
Carga el estado de la biblioteca desde archivo
*/
func cargarDatos(archivo string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error crítico cargando datos: %v\n", r)
		}
	}()

	// TODO: Implementa aquí
	return nil
}

/*
TODO: Implementa crearBackup
Retorna: string (nombre del archivo de backup), error
Crea una copia de seguridad con timestamp
*/
func crearBackup() (string, error) {
	// TODO: Implementa aquí
	return "", nil
}

// FUNCIONES DE VALIDACIÓN

/*
TODO: Implementa validarISBN
Parámetros: isbn string
Retorna: bool
Valida formato de ISBN-10 o ISBN-13
*/
func validarISBN(isbn string) bool {
	// TODO: Implementa aquí
	return false
}

/*
TODO: Implementa validarEmail
Parámetros: email string
Retorna: bool
Valida formato básico de email
*/
func validarEmail(email string) bool {
	// TODO: Implementa aquí
	return false
}

/*
TODO: Implementa validarTipoMembresia
Parámetros: tipo string
Retorna: bool
Valida que el tipo de membresía sea válido
*/
func validarTipoMembresia(tipo string) bool {
	// TODO: Implementa aquí
	return false
}

// FUNCIONES DE UTILIDAD

/*
TODO: Implementa calcularFechaVencimiento
Parámetros: tipoMembresia string
Retorna: string (fecha de vencimiento)
Calcula fecha de vencimiento basada en tipo de membresía
*/
func calcularFechaVencimiento(tipoMembresia string) string {
	// TODO: Implementa aquí
	return ""
}

/*
TODO: Implementa calcularMulta
Parámetros: fechaVencimiento, fechaDevolucion, tipoMembresia string
Retorna: float64
Calcula multa por días de retraso
*/
func calcularMulta(fechaVencimiento, fechaDevolucion, tipoMembresia string) float64 {
	// TODO: Implementa aquí
	return 0
}

/*
TODO: Implementa formatearReporte
Parámetros: datos map[string]interface{}
Imprime un reporte formateado
*/
func formatearReporte(datos map[string]interface{}) {
	// TODO: Implementa aquí
}

// INTERFAZ DE LÍNEA DE COMANDOS

/*
TODO: Implementa mostrarMenu
Muestra el menú principal del sistema
*/
func mostrarMenu() {
	// TODO: Implementa aquí
}

/*
TODO: Implementa procesarComando
Parámetros: comando string, args []string
Retorna: error
Procesa comandos de la interfaz de línea de comandos
*/
func procesarComando(comando string, args []string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error procesando comando '%s': %v\n", comando, r)
		}
	}()

	// TODO: Implementa aquí
	return nil
}

// FUNCIÓN PRINCIPAL

func main() {
	defer func() {
		// Guardar datos antes de salir
		if err := guardarDatos("biblioteca.data"); err != nil {
			fmt.Printf("Error guardando datos: %v\n", err)
		}
		fmt.Println("Sistema de biblioteca cerrado.")
	}()

	fmt.Println("=== SISTEMA DE GESTIÓN DE BIBLIOTECA ===")
	fmt.Println("Bienvenido al sistema de gestión de biblioteca")

	// Verificar argumentos de línea de comandos
	if len(os.Args) < 2 {
		mostrarMenu()
		os.Exit(EXIT_USAGE)
	}

	comando := os.Args[1]
	args := os.Args[2:]

	// Procesar comando
	if err := procesarComando(comando, args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(EXIT_ERROR)
	}

	// TODO: Implementa la lógica principal del programa
	// Sugerencias:
	// 1. Cargar datos existentes
	// 2. Procesar comandos de línea
	// 3. Ejecutar operaciones solicitadas
	// 4. Mostrar resultados
	// 5. Guardar cambios

	fmt.Println("Operación completada exitosamente")
	os.Exit(EXIT_SUCCESS)
}

/*
INSTRUCCIONES DE IMPLEMENTACIÓN:
===============================

1. 🚀 EMPEZAR:
   - Define primero todas las estructuras de datos
   - Implementa las funciones init y constructor
   - Crea datos de prueba para verificar funcionalidad

2. 📝 DESARROLLO:
   - Implementa una función a la vez
   - Prueba cada función antes de continuar
   - Maneja todos los casos de error

3. 🧪 TESTING:
   - Crea casos de prueba para cada función
   - Verifica comportamiento con datos inválidos
   - Asegúrate de que el sistema no haga panic

4. 📊 FUNCIONALIDAD MÍNIMA:
   - Al menos poder agregar libros y usuarios
   - Realizar préstamos y devoluciones básicas
   - Generar reporte simple
   - Guardar y cargar datos

5. 🎯 FUNCIONALIDAD AVANZADA:
   - Sistema completo de multas
   - Reportes detallados
   - Validaciones exhaustivas
   - Interfaz de línea de comandos completa

6. 🏆 BONUS:
   - Búsqueda avanzada con múltiples criterios
   - Sistema de reservas
   - Notificaciones de vencimiento
   - Estadísticas temporales

CRITERIOS DE EVALUACIÓN:
======================
✅ Código compila sin errores
✅ Todas las funciones están implementadas
✅ Manejo apropiado de errores
✅ Uso correcto de defer, panic, recover
✅ Validaciones de datos robustas
✅ Comentarios claros y útiles
✅ Funcionalidad básica opera correctamente
✅ Código bien estructurado y legible

¡Demuestra todo lo que has aprendido! 🚀
*/
