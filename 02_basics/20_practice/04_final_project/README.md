# Proyecto Final: Sistema de Gestión de Biblioteca

## 📋 Descripción General

Este proyecto final integra **TODOS** los conceptos aprendidos en el bootcamp de Go. Desarrollarás un sistema completo de gestión de biblioteca que demuestre tu dominio de:

- ✅ Variables, tipos de datos y constantes
- ✅ Operadores y control de flujo
- ✅ Arrays, slices y maps
- ✅ Funciones (básicas, múltiples retornos, variádicas)
- ✅ Defer, panic, recover
- ✅ Init y exit
- ✅ Manejo de errores

## 🎯 Objetivos del Proyecto

### Funcionalidades Principales

1. **📚 Gestión de Libros**
   - Agregar, buscar, actualizar y eliminar libros
   - Categorización por género, autor, año
   - Sistema de ISBN y validaciones

2. **👥 Gestión de Usuarios**
   - Registro de miembros de la biblioteca
   - Diferentes tipos de membresía (estudiante, profesor, público)
   - Validación de datos de usuario

3. **📖 Sistema de Préstamos**
   - Préstamo y devolución de libros
   - Control de fechas de vencimiento
   - Multas por retraso
   - Historial de préstamos

4. **📊 Reportes y Estadísticas**
   - Libros más populares
   - Usuarios más activos
   - Estadísticas de la biblioteca
   - Reportes de multas

5. **💾 Persistencia de Datos**
   - Guardar y cargar datos (simulado)
   - Backup y restauración
   - Validación de integridad

## 🛠️ Requisitos Técnicos

### Obligatorios
- [ ] Usar todos los tipos de datos aprendidos
- [ ] Implementar al menos 20 funciones diferentes
- [ ] Usar funciones variádicas apropiadamente
- [ ] Implementar manejo de errores robusto
- [ ] Usar defer para cleanup de recursos
- [ ] Implementar recovery para operaciones críticas
- [ ] Usar init para configuración inicial
- [ ] Manejar exit codes apropiadamente
- [ ] Validación exhaustiva de datos de entrada
- [ ] Comentarios detallados en todo el código

### Opcionales (Bonus)
- [ ] Interfaz de línea de comandos completa
- [ ] Sistema de búsqueda avanzada
- [ ] Notificaciones de vencimiento
- [ ] Estadísticas temporales

## 📁 Estructura del Proyecto

```
04_final_project/
├── README.md                           # Este archivo
├── library_management_system.go        # Archivo principal con estructura
├── docs/                              # Documentación adicional
│   ├── user_manual.md                 # Manual de usuario
│   └── technical_specs.md             # Especificaciones técnicas
└── data/                              # Archivos de datos (crear durante desarrollo)
    ├── biblioteca.data                # Datos principales
    └── backups/                       # Copias de seguridad
```

## 🚀 Guía de Implementación

### Fase 1: Fundamentos (Nivel Básico)
1. **Definir Estructuras de Datos**
   - `Libro`, `Usuario`, `Prestamo`, `Biblioteca`
   - Implementar constructores básicos

2. **Funciones Básicas CRUD**
   - `agregarLibro()`, `registrarUsuario()`
   - `buscarLibro()`, `buscarUsuario()`
   - Validaciones básicas

3. **Sistema de Préstamos Simple**
   - `prestarLibro()`, `devolverLibro()`
   - Control básico de disponibilidad

### Fase 2: Funcionalidad Intermedia (Nivel Intermedio)
1. **Búsquedas Avanzadas**
   - Búsqueda por múltiples criterios
   - Filtros y ordenamiento

2. **Sistema de Multas**
   - Cálculo automático de multas
   - Gestión de fechas de vencimiento

3. **Reportes Básicos**
   - Estadísticas generales
   - Listados formateados

### Fase 3: Características Avanzadas (Nivel Avanzado)
1. **Persistencia de Datos**
   - Guardar/cargar desde archivos
   - Sistema de backup

2. **Manejo de Errores Robusto**
   - Recovery de panics
   - Logging detallado

3. **Interfaz de Usuario**
   - Línea de comandos
   - Menús interactivos

## 📖 Ejemplos de Uso

### Agregar un Libro
```bash
go run library_management_system.go agregar-libro \
  --isbn="978-0134190440" \
  --titulo="The Go Programming Language" \
  --autor="Alan Donovan" \
  --genero="Programación" \
  --ano=2015
```

### Registrar Usuario
```bash
go run library_management_system.go registrar-usuario \
  --nombre="Juan Pérez" \
  --email="juan@example.com" \
  --tipo="estudiante"
```

### Realizar Préstamo
```bash
go run library_management_system.go prestar \
  --usuario-id=1 \
  --isbn="978-0134190440"
```

### Generar Reporte
```bash
go run library_management_system.go reporte \
  --tipo="general"
```

## 🧪 Testing y Validación

### Casos de Prueba Mínimos
- [ ] Agregar libro con datos válidos
- [ ] Rechazar libro con ISBN duplicado
- [ ] Registrar usuario con email válido
- [ ] Rechazar usuario con email inválido
- [ ] Prestar libro disponible
- [ ] Rechazar préstamo de libro no disponible
- [ ] Calcular multa correctamente
- [ ] Guardar y cargar datos sin pérdida

### Casos de Prueba Avanzados
- [ ] Manejo de archivos corruptos
- [ ] Recovery de panics en operaciones críticas
- [ ] Validación de límites de préstamo por tipo de usuario
- [ ] Generación de reportes con datos complejos

## 🏆 Criterios de Evaluación

### Excelente (90-100%)
- ✅ Todas las funcionalidades implementadas
- ✅ Código limpio y bien documentado
- ✅ Manejo de errores robusto
- ✅ Casos edge manejados apropiadamente
- ✅ Interfaz de usuario intuitiva

### Bueno (80-89%)
- ✅ Funcionalidades principales implementadas
- ✅ Código funcional con documentación básica
- ✅ Manejo de errores básico
- ✅ Algunos casos edge manejados

### Satisfactorio (70-79%)
- ✅ Funcionalidades básicas implementadas
- ✅ Código compila y ejecuta
- ✅ Manejo mínimo de errores

### Necesita Mejora (<70%)
- ❌ Funcionalidades incompletas
- ❌ Errores de compilación
- ❌ Falta de manejo de errores

## 💡 Consejos y Mejores Prácticas

### Durante el Desarrollo
1. **Desarrolla Incrementalmente**
   - Implementa una función a la vez
   - Prueba antes de continuar

2. **Maneja Errores Apropiadamente**
   - Nunca ignores errores
   - Proporciona mensajes descriptivos

3. **Usa Defer Correctamente**
   - Para cerrar archivos
   - Para cleanup de recursos

4. **Documenta Tu Código**
   - Explica la lógica compleja
   - Incluye ejemplos de uso

### Debugging
1. **Usa Logging**
   - Registra operaciones importantes
   - Incluye timestamps

2. **Valida Datos de Entrada**
   - Verifica todos los parámetros
   - Maneja casos edge

3. **Prueba Casos Extremos**
   - Datos vacíos
   - Valores límite
   - Errores de sistema

## 🎯 Entrega del Proyecto

### Archivos Requeridos
1. `library_management_system.go` - Código principal completamente implementado
2. `README.md` - Documentación de tu implementación
3. Archivos de prueba con datos de ejemplo
4. Screenshots o ejemplos de funcionamiento

### Formato de Entrega
- Código limpio y comentado
- Funciones implementadas y probadas
- Documentación clara de cómo ejecutar
- Ejemplos de uso incluidos

## 🌟 ¡A Programar!

Este proyecto es tu oportunidad de demostrar todo lo que has aprendido. No tengas miedo de experimentar y ser creativo. Recuerda:

- 🧠 **Piensa antes de codificar** - Planifica tu enfoque
- 🔍 **Lee cuidadosamente** - Entiende cada requerimiento
- 🧪 **Prueba frecuentemente** - Verifica que todo funciona
- 📝 **Documenta todo** - Tu código debe ser legible
- 🎯 **Enfócate en la calidad** - Es mejor menos funciones bien hechas

¡Buena suerte y que disfrutes construyendo tu sistema de biblioteca! 🚀📚