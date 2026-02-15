# Go Mini Market API

Una API REST robusta y segura para la gestión de un mini mercado, construida con Go (Golang) y PostgreSQL. Este proyecto demuestra prácticas de ingeniería de software sólidas, incluyendo manejo de concurrencia, transacciones, precisión monetaria, middleware y testing completo.

## Características Principales

* **Arquitectura Limpia:** Estructura de proyecto estándar (`cmd/`, `internal/`, `migrations/`).
* **CRUDs Completos:** Gestión de Categorías, Productos y Clientes.
* **Gestión de Ventas Transaccional:**
  * Integridad de datos garantizada con transacciones SQL (`BEGIN`, `COMMIT`, `ROLLBACK`).
  * Actualización automática de stock.
* **Manejo Seguro de Dinero:** Uso de la librería `shopspring/decimal` para evitar errores de redondeo de punto flotante.
* **Prevención de Race Conditions:** Uso de `SELECT ... FOR UPDATE` para bloquear filas de inventario durante la venta.
* **Middleware Completo:**
  * **Logger:** Registro de todas las requests con método, ruta, status y duración.
  * **Recovery:** Captura de panics para evitar caídas del servidor.
  * **CORS:** Soporte para solicitudes cross-origin configurable.
  * **Request ID:** Identificador único por request para tracing.
  * **Timeout:** Límite de tiempo para requests.
* **Testing Robusto:** Tests unitarios para handlers y middleware con mocking.
* **Configuración Segura:** Variables de entorno vía `.env`.
* **Base de Datos Eficiente:** Uso de `pgx/v5` con pool de conexiones.

## Tech Stack

* **Lenguaje:** Go 1.22+
* **Base de Datos:** PostgreSQL
* **Drivers & Librerías:**
  * [`pgx/v5`](github.com/jackc/pgx): Driver de PostgreSQL de alto rendimiento.
  * [`shopspring/decimal`](github.com/shopspring/decimal): Matemáticas decimales de precisión arbitraria.
  * [`godotenv`](github.com/joho/godotenv): Carga de variables de entorno.
  * [`google/uuid`](github.com/google/uuid): Generación de UUIDs para request IDs.
* **Router:** `net/http` (Go 1.22 standard library router).

## Estructura del Proyecto

```bash
.
├── cmd
│   ├── api          # Punto de entrada del servidor (main.go)
│   └── seed         # Script para poblar la base de datos con datos de prueba
├── internal
│   ├── database     # Lógica de acceso a datos (Repositorio/Service)
│   ├── handlers     # Controladores HTTP
│   ├── middleware   # Middleware (logger, recovery, cors, requestid, timeout)
│   ├── models       # Estructuras de dominio
│   └── server       # Configuración del servidor HTTP y rutas
├── migrations       # Scripts SQL para crear el esquema
├── docs
│   └── API_TESTING.md   # Guía rápida de comandos curl
└── go.mod
```

## Configuración y Ejecución

### 1. Prerrequisitos

* Go 1.22 o superior
* PostgreSQL corriendo

### 2. Variables de Entorno

Crea un archivo `.env` en la raíz del proyecto:

```env
PORT=8080
APP_ENV=development
DB_CONN_STR=postgres://usuario:password@localhost:5432/web_go?sslmode=disable
CORS_ORIGINS=*                    # * para desarrollo, o dominios específicos separados por coma
```

### 3. Base de Datos

Crea la base de datos `web_go` y ejecuta los scripts de migración en orden:

```bash
psql -U postgres -d web_go -f migrations/001_create_categories.sql
psql -U postgres -d web_go -f migrations/002_create_products.sql
psql -U postgres -d web_go -f migrations/003_create_customers.sql
psql -U postgres -d web_go -f migrations/004_create_sales.sql
```

### 4. Poblar Base de Datos (Seed)

Ejecuta el script de seed para generar categorías, productos, clientes y ventas de prueba:

```bash
go run cmd/seed/main.go
```

### 5. Iniciar el Servidor

```bash
go run cmd/api/main.go
```

El servidor iniciará en `http://localhost:8080`.

## Testing

### Ejecutar Todos los Tests

```bash
go test ./...
```

### Ejecutar Tests con Verbosidad

```bash
go test ./... -v
```

### Ejecutar Tests con Coverage

```bash
go test ./... -cover
```

### Tests Disponibles

* **`internal/handlers/`**: Tests unitarios para Productos y Ventas
  * Validación de JSON
  * Casos de éxito y error
  * Mock de base de datos
  
* **`internal/middleware/`**: Tests para todos los middlewares
  * Recovery de panics
  * CORS
  * Request ID
  * Timeout

## Middleware

Los middlewares se aplican en cadena en el siguiente orden:

1. **Timeout** - Cancela requests que exceden 30 segundos
2. **CORS** - Habilita CORS según configuración
3. **RequestID** - Agrega UUID único a cada request
4. **Recovery** - Captura panics y retorna 500
5. **Logger** - Loguea método, ruta, status y duración

### Headers de Respuesta

Cada respuesta incluye:

* `X-Request-ID`: UUID único del request para tracing
* `Access-Control-Allow-*`: Headers CORS configurables

## Pruebas de API

Consulta el archivo [`docs/API_TESTING.md`](./docs/API_TESTING.md) para ver una lista completa de comandos `curl` listos para copiar y pegar.

## Notas de Implementación

### Precisión Monetaria

No usamos `float64` para precios. Usamos `decimal.Decimal` para garantizar que `$10.10 + $20.20` sea exactamente `$30.30`.

### Concurrencia en Inventario

El sistema es seguro para entornos de alta concurrencia mediante `SELECT ... FOR UPDATE`.

### Arquitectura de Middleware

```go
handler = middleware.Logger(handler)
handler = middleware.Recovery(handler)
handler = middleware.RequestID(handler)
handler = middleware.Timeout(30 * time.Second)(handler)
handler = middleware.CORS(origins)(handler)
```

Esta cadena se aplica a todas las rutas de forma uniforme.
